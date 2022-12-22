import ast
import asyncio
import json
import leveldb
import os
import re
import math

from gmpy import binary, mpz
from gmpy2 import mpz_from_old_binary
from zkrp_pyo3 import pedersen_aggregate, pedersen_commit, zkrp_verify, zkrp_prove, zkrp_prove_mul, zkrp_verify_mul, other_base_commit, product_com

INPUTMASK_SHARES_DIR = os.getenv(
    'INPUTMASK_SHARES', '/opt/hbswap/inputmask-shares',
)


def parse_contract(name):
    contract = json.load(open(f'ratel/genfiles/build/contracts/{name}.json'))
    return contract['abi'], contract['bytecode']


def sign_and_send(tx, web3, account):
    signedTx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)
    tx_hash = web3.eth.send_raw_transaction(signedTx.rawTransaction)
    return web3.eth.wait_for_transaction_receipt(tx_hash)


def getAccount(web3, keystoreDir):
    for filename in os.listdir(keystoreDir):
        with open(keystoreDir + filename) as keyfile:
            encryptedKey = keyfile.read()
            privateKey = web3.eth.account.decrypt(encryptedKey, '')
            return web3.eth.account.privateKeyToAccount(privateKey)


class MultiAcquire(asyncio.Task):
    _check_lock = asyncio.Lock()  # to suspend for creating task that acquires objects
    _release_event = asyncio.Event()  # to suspend for any object was released

    def __init__(self, locks):
        super().__init__(self._task_coro())
        self._locks = locks
        # Here we use decorator to subscribe all release() calls,
        # _release_event would be set in this case:
        for l in self._locks:
            l.release = self._notify(l.release)


    async def _task_coro(self):
        while True:
            # Create task to acquire all locks and break on success:
            async with type(self)._check_lock:
                if not any(l.locked() for l in self._locks):  # task would be created only if all objects can be acquired
                    task = asyncio.gather(*[l.acquire() for l in self._locks])  # create task to acquire all objects
                    await asyncio.sleep(0)  # start task without waiting for it
                    break
            # Wait for any release() to try again:
            await type(self)._release_event.wait()
        # Wait for task:
        return await task


    def _notify(self, func):
        def wrapper(*args, **kwargs):
            type(self)._release_event.set()
            type(self)._release_event.clear()
            return func(*args, **kwargs)
        return wrapper


def mpcPort(seq, concurrency):
    if seq > 0:
        return mpc_port + seq % concurrency * 100
    else:
        return mpc_port + seq * 100


def key_inputmask_index(idx):
    return f'inputmask_index_{idx}'.encode()


def key_inputmask_version(idx):
    return f'inputmask_version_{idx}'.encode()

def key_zkrp_blinding_index(idx, num = 1):
    return f'zkrp_blinding_index_{idx}_{num}'.encode()

def key_zkrp_blinding_commitment_index(idx):
    return f'zkrp_blinding_commitment_index_{idx}'.encode()

def key_zkrp_agg_commitment_index(idx):
    return f'zkrp_agg_commitment_index_{idx}'.encode()


def key_serverval_index(idx):
    idx = idx.lower()
    return f'{idx}'.encode()


def location_sharefile(server_id, base_port):
    return f'Persistence/Transactions-P{server_id}-{base_port}.data'


def location_db(server_id):
    db_path = os.getenv('DB_PATH', '/opt/hbswap/db')
    return f'{db_path}/server-{server_id}'


def location_inputmask(server_id, players):
    inputmask_shares_dir = os.getenv(
        'INPUTMASK_SHARES', '/opt/hbswap/inputmask-shares',
    )
    return f'{inputmask_shares_dir}/{players}-MSp-{prime_bit_length}/Randoms-MSp-P{server_id}'


def openDB(location):
    return leveldb.LevelDB(location)


def hex_to_int(x):
    x = mpz_from_old_binary(x)
    return int((x * inverse_R) % prime)


def int_to_hex(x):
    x = mpz(x)
    x = (x * R) % prime
    x = binary(int(x))
    x += b'\x00' * (32 - len(x))
    return x


def get_inverse(a):
    ret = 1
    b = prime - 2
    while b:
        if b % 2 == 1:
            ret = (ret * a) % prime
        b //= 2
        a = (a * a) % prime
    return ret


def recover_input(db, masked_value, idx): # return: int
    try:
        input_mask_share = db.Get(key_inputmask_index(idx))
    except KeyError:
        input_mask_share = bytes(0)
    input_mask_share = int.from_bytes(input_mask_share, 'big')
    return (masked_value - input_mask_share) % prime


def players(contract):
    players = contract.functions.N().call()
    return players


def threshold(contract):
    threshold = contract.functions.T().call()
    return threshold


def list_to_str(list):
    st = ''
    for v in list:
        st += f"{',' if len(st) > 0 else ''}{v}"
    return st


async def execute_cmd(cmd, info=''):
    retry = mpc_failed_retry
    while True:
        proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
        stdout, stderr = await proc.communicate()

        print(f'[{cmd!r} exited with {proc.returncode}]')
        if info:
            print(f'[info]\n{info}')
        if stdout:
            print(f'[stdout]\n{stdout.decode()}')

        returncode = proc.returncode
        if returncode != 0:
            print(f'[stderr]\n{stderr.decode()}')

        retry -= 1

        if retry <= 0 or returncode == 0:
            if returncode != 0:
                print('**** ERROR')

            return returncode


def mark_finish(server, seq):
    port = mpcPort(seq, server.concurrency)
    server.portLock[port].release()

    key = 'execHistory'
    execHistory = read_db(server, key)
    execHistory = bytes_to_dict(execHistory)

    execHistory[seq] = True

    execHistory = dict_to_bytes(execHistory)
    write_db(server, key, execHistory)


def read_db(server, key, finalize_on_chain=False):
    key = key.lower()
    try:
        value = server.db.Get(key.encode())
    except KeyError:
        value = bytes(0)

    if key in server.dbLock.keys():
        server.dbLockCnt[key] -= 1
        if server.dbLockCnt[key] == 0:
            if not finalize_on_chain:
                server.dbLock[key].release()

    return value


def write_db(server, key, value, finalize_on_chain=False):
    key = key.lower()
    server.db.Put(key.encode(), value)

    if key in server.dbLock.keys():
        server.dbLockCnt[key] -= 1
        if server.dbLockCnt[key] == 0:
            if not finalize_on_chain:
                server.dbLock[key].release()


def bytes_to_int(value):
    return int.from_bytes(value, 'big')


def bytes_to_list(value):
    try:
        value = value.decode(encoding='utf-8')
        value = list(ast.literal_eval(value))
    except:
        value = []
    return value


def bytes_to_dict(value):
    try:
        value = value.decode(encoding='utf-8')
        value = dict(ast.literal_eval(value))
    except:
        value = {}
    return value


def int_to_bytes(value):
    return value.to_bytes((value.bit_length() + 7) // 8, 'big')


def list_to_bytes(value):
    return bytes(str(value), encoding='utf-8')


def dict_to_bytes(value):
    return bytes(str(value), encoding='utf-8')


async def verify_proof(server, x, zkpstmt, type_Mul = 0, y = 1, r = 0):
    [idxValueBlinding, maskedValueBlinding, proof, commitment] = zkpstmt
    # TODO:
    # proof, commitment, blinding_ = zkrp_prove(2022, 32)
    if proof is None or commitment is None or not zkrp_verify(proof, commitment):
        print("[Error]: Committed secret value does not pass range proof verification!")
        return False

    blinding = recover_input(server.db, maskedValueBlinding, idxValueBlinding)
    x,y,r = int(x), int(y), int(r)

    if type_Mul == 0:
        pfval = x % prime
        if pfval < 0:
            pfval = (pfval % prime + prime) % prime
        # print('pfval2:',pfval)
        
        # TODO: where is the blinding mask created? we also need to share it.
        value1_bytes = list(pfval.to_bytes(32, byteorder='little'))
        blinding_bytes = list(blinding.to_bytes(32, byteorder='little'))

        share_commitment = pedersen_commit(value1_bytes, blinding_bytes)

        # print('share_commitment:',share_commitment)
        # TODO: create the function to commit to the unmasked secret shares.
        # TODO: we also need to change the current zkrp interface to allow specifying r and choose range to prove.

        server.zkrpShares[f'{idxValueBlinding}'] = share_commitment
        results = await server.get_zkrp_shares(players(server.contract), f'{idxValueBlinding}')
        # print(")))))))", results)
        agg_commitment = pedersen_aggregate(results, [x + 1 for x in list(range(server.players))])

        # print("agg_commit:", agg_commitment, commitment)
        return agg_commitment == commitment
    else: ### x * y >= r
        r = -r
        r = (r % prime + prime) % prime
        if type_Mul >= 3:
            x = -x
            x = (x % prime + prime) % prime
        ############# (1) compute g^[x] #############
        zer = 0
        x_bytes = list(x.to_bytes(32, byteorder='little'))
        zer_bytes = list(zer.to_bytes(32, byteorder='little'))
        g_x_share = pedersen_commit(x_bytes, zer_bytes) 
        server.zkrpShares[f'{idxValueBlinding}_{0}'] = g_x_share
        results_g_x = await server.get_zkrp_shares(players(server.contract), f'{idxValueBlinding}_{0}')
        # print('results_g_x',results_g_x)
        g_x_bytes = pedersen_aggregate(results_g_x, [x + 1 for x in list(range(server.players))])

        ############# (2) compute (g^x)^[y] * h^[rz] #############
        rz_bytes = list(blinding.to_bytes(32, byteorder='little'))
        y_bytes = list(y.to_bytes(32, byteorder='little'))
        g_xy_h_rz_bytes = other_base_commit(g_x_bytes, y_bytes, rz_bytes)
        # print('g_xy_h_rz_bytes',g_xy_h_rz_bytes)
        server.zkrpShares[f'{idxValueBlinding}_{1}'] = g_xy_h_rz_bytes
        results_g_xy_h_rz = await server.get_zkrp_shares(players(server.contract), f'{idxValueBlinding}_{1}')
        # print('results_g_xy_h_rz',results_g_xy_h_rz)
        agg_gxyhrz_commitment = pedersen_aggregate(results_g_xy_h_rz, [x + 1 for x in list(range(server.players))])
        # print('agg_commitment',agg_gxyhrz_commitment)

        if r != 0:
            ### fixme!
            r_bytes = list(r.to_bytes(32, byteorder='little'))
            g_r = pedersen_commit(r_bytes, zer_bytes) 
            print('g_r', g_r)
            agg_commitment = product_com(g_r,agg_gxyhrz_commitment)
        else:
            agg_commitment = agg_gxyhrz_commitment
        # print('agg_commitment',agg_commitment)
        # print('commitment',commitment)
        return agg_commitment == commitment



def get_zkrp(secret_value, exp_str, r, isSfix = False):
    value = secret_value

    # if exp_str == '>=':
    #     value = int((value - r) * fac)
    # elif exp_str == '>': #secret_value > r <==> secret_value - r -1 >= 0
    #     value = int((value - r) * fac) - 1
    # elif exp_str == '<=': # secret_value <= r <==> r - secret_value >= 0 
    #     value = int((r - value) * fac)
    # elif exp_str == '<': #secret_value < r <==> r - secret_value - 1 >= 0
    #     value = int((r - value) * fac) - 1

    if isSfix:
        value = int(value * fp)
        r = int(r * fp)
    
    if exp_str == '>=':
        value = value - r
    elif exp_str == '>': #secret_value > r <==> secret_value - r -1 >= 0
        value = value - r - 1
    elif exp_str == '<=': # secret_value <= r <==> r - secret_value >= 0 
        value = r - value
    elif exp_str == '<': #secret_value < r <==> r - secret_value - 1 >= 0
        value = r - value - 1


    value = (value % prime + prime) % prime
    # if value < 0 :
    #     value = (value % prime + prime) % prime

    print('value:',value)

    #To prove value >= 0
    bits = 32
    proof, commitment, blinding_bytes = zkrp_prove(value, bits)
    blinding = int.from_bytes(blinding_bytes, byteorder='little')
    return proof, commitment, blinding

leaderHostname = 'mpcnode0'

prog = './malicious-shamir-party.x'
offline_prog = './mal-shamir-offline.x'

### blsPrime
# prime = 52435875175126190479447740508185965837690552500527637822603658699938581184513
# R = 10920338887063814464675503992315976177888879664585288394250266608035967270910
# prime_bit_length = 255

### Ristretto group order
prime = 7237005577332262213973186563042994240857116359379907606001950938285454250989
R = 7237005577332262213973186563042994240413239274941949949428319933631315875101
prime_bit_length = 253
inverse_R = get_inverse(R)

inv_10 = 723700557733226221397318656304299424085711635937990760600195093828545425099

fp = 2 ** 16
decimal = 10 ** 15
sz = 32

http_host = "0.0.0.0"
http_port = 4000

mpc_port = 5000

spareShares = 1000
shareBatchSize = 10000

confirmation = 2

trade_key_num = 7

repeat_experiment = 1

mpc_failed_retry = 3

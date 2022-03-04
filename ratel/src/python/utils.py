import asyncio
import json
import leveldb
import os

from gmpy import binary, mpz
from gmpy2 import mpz_from_old_binary

INPUTMASK_SHARES_DIR = os.getenv(
    'INPUTMASK_SHARES', '/opt/hbswap/inputmask-shares',
)

def parse_contract(name):
    contract = json.load(open(f'ratel/genfiles/build/contracts/{name}.json'))
    return contract['abi'], contract['bytecode']

def sign_and_send(tx, web3, account):
    signedTx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)
    web3.eth.send_raw_transaction(signedTx.rawTransaction)
    web3.eth.wait_for_transaction_receipt(signedTx.hash)
    return signedTx.hash

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
    return mpc_port + seq % concurrency * 100

def key_inputmask(idx):
    return f'inputmask_{idx}'.encode()

def location_sharefile(server_id, base_port):
    return f'Persistence/Transactions-P{server_id}-{base_port}.data'

def location_db(server_id):
    db_path = os.getenv('DB_PATH', '/opt/hbswap/db')
    return f'{db_path}/server-{server_id}'

def location_inputmask(server_id):
    return f'{INPUTMASK_SHARES_DIR}/4-MSp-255/Randoms-MSp-P{server_id}'

def openDB(location):
    return leveldb.LevelDB(location)

def hex_to_int(x):
    return int((mpz_from_old_binary(x) * inverse_R) % blsPrime)

def int_to_hex(x):
    x = mpz(x)
    x = (x * R) % blsPrime
    x = binary(int(x))
    x += b'\x00' * (32 - len(x))
    return x

def get_inverse(a):
    ret = 1
    b = blsPrime - 2
    while b:
        if b % 2 == 1:
            ret = (ret * a) % blsPrime
        b //= 2
        a = (a * a) % blsPrime
    return ret

def recover_input(db, masked_value, idx): # return: int
    try:
        input_mask_share = db.Get(key_inputmask(idx))
    except KeyError:
        input_mask_share = bytes(0)
    input_mask_share = int.from_bytes(input_mask_share, 'big')
    return (masked_value - input_mask_share) % blsPrime

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

prog = './malicious-shamir-party.x'
blsPrime = 52435875175126190479447740508185965837690552500527637822603658699938581184513
leaderHostname = 'mpcnode0'
R = 10920338887063814464675503992315976177888879664585288394250266608035967270910
inverse_R = get_inverse(R)
fp = 2 ** 16
decimal = 10 ** 15
sz = 32

http_host = "0.0.0.0"
http_port = 4000

mpc_port = 5000
# concurrency = 2

spareShares = 100
shareBatchSize = 1000

confirmation = 2

replay = 1

trade_key_num = 7
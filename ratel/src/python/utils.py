import leveldb
import os

from gmpy import binary, mpz
from gmpy2 import mpz_from_old_binary

def mpcPort(seq):
    return mpc_port + seq % concurrency

def key_inputmask(idx):
    return f'inputmask_{idx}'.encode()

def location_sharefile(server_id):
    return f'Persistence/Transactions-P{server_id}.data'

def location_db(server_id):
    db_path = os.getenv('DB_PATH', '/opt/hbswap/db')
    return f'{db_path}/server-{server_id}'

def location_inputmask(server_id):
    inputmask_shares_dir = os.getenv(
        'INPUTMASK_SHARES', '/opt/hbswap/inputmask-shares',
    )
    return f'{inputmask_shares_dir}/4-MSp-255/Randoms-MSp-P{server_id}'

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
    print('players', players)
    return players

def threshold(contract):
    threshold = contract.functions.T().call()
    print('threshold', threshold)
    return threshold

def reconstruct(shares, n):
    inputmask = 0
    for i in range(1, n + 1):
        tot = 1
        for j in range(1, n + 1):
            if i == j:
                continue
            tot = tot * j * get_inverse(j - i) % blsPrime
        inputmask = (inputmask + shares[i - 1] * tot) % blsPrime
    print(inputmask)
    return inputmask

prog = './malicious-shamir-party.x'
blsPrime = 52435875175126190479447740508185965837690552500527637822603658699938581184513
leaderHostname = 'mpcnode0'
R = 10920338887063814464675503992315976177888879664585288394250266608035967270910
inverse_R = get_inverse(R)
fp = 2 ** 16
decimal = 10 ** 15

sz = 32

spareShares = 100
batchShares = 1000

http_host = "0.0.0.0"
http_port = 4000

mpc_port = 5000
concurrency = 1

confirmation = 2
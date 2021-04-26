import os
import time

import leveldb

from gmpy import binary, mpz
from gmpy2 import mpz_from_old_binary

##############################################################################
# keys                                                                       #
##############################################################################


def key_balance(token, user):
    return f'balance_{token}_{user}'.encode()


def key_inputmask(idx):
    return f'inputmask_{idx}'.encode()

def key_cnt_failed_trade():
    return f'cnt_failed_trade'.encode()

def key_cnt_succeed_trade():
    return f'cnt_succeed_trade'.encode()

def key_pool(token_A, token_B, token):
    return f'pool_{token_A}_{token_B}_{token}'.encode()


def key_individual_price(trade_seq):
    return f'trade_price_{trade_seq}'.encode()


# def key_trade_time(trade_seq):
#     return f'trade_time_{trade_seq}'.encode()


def key_total_price(token_A, token_B):
    return f'trade_price_{token_A}_{token_B}'.encode()


def key_total_cnt(token_A, token_B):
    return f'total_cnt_{token_A}-{token_B}'.encode()


def key_total_supply(token_A, token_B):
    return f'total_supply_{token_A}_{token_B}'.encode()


##############################################################################
# files                                                                      #
##############################################################################


def location_db(server_id):
    db_path = os.getenv('DB_PATH', '/opt/hbswap/db')
    return f'{db_path}/server-{server_id}'


def location_inputmask(server_id):
    inputmask_shares_dir = os.getenv(
        'INPUTMASK_SHARES', '/opt/hbswap/inputmask-shares',
    )
    return f'{inputmask_shares_dir}/4-MSp-255/Randoms-MSp-P{server_id}'


def location_private_output(server_id):
    prep_dir = os.getenv('PREP_DIR', '/opt/hbswap/preprocessing-data')
    return f'{prep_dir}/Private-Output-{server_id}'


def location_sharefile(server_id):
    return f'Persistence/Transactions-P{server_id}.data'


##############################################################################
# functions                                                                  #
##############################################################################


def openDB(location):
    while True:
        try:
            return leveldb.LevelDB(location)
        except leveldb.LevelDBError:
            print('db not ready')
            time.sleep(10)


def get_value(db, key): # return: hex
    try:
        return bytes(db.Get(key))
    except KeyError:
        return int_to_hex(0)


def get_inverse(a):
    ret = 1
    b = p - 2
    while b:
        if b % 2 == 1:
            ret = (ret * a) % p
        b //= 2
        a = (a * a) % p
    return ret


# def float_to_fix(x):
#     return int(round(float(x) * fp))


def fix_to_float(x):
    return 1. * x / fp


def hex_to_int(x):
    return int((mpz_from_old_binary(x) * inverse_R) % p)


def int_to_hex(x):
    x = mpz(x)
    x = (x * R) % p
    x = binary(int(x))
    x += b'\x00' * (32 - len(x))
    return x


def check_consistency(shares):
    value = reconstruct(shares, t + 1)
    for i in range(t + 2, n + 1):
        if reconstruct(shares, i) != value:
            print('inconsistent')


def reconstruct(shares, n):
    inputmask = 0
    for i in range(1, n + 1):
        tot = 1
        for j in range(1, n + 1):
            if i == j:
                continue
            tot = tot * j * get_inverse(j - i) % p
        inputmask = (inputmask + shares[i - 1] * tot) % p
    print(inputmask)
    return inputmask

def recover_input(db, masked_value, idx): # return: hex
    input_mask_share = hex_to_int(get_value(db, key_inputmask(idx)))
    return int_to_hex((masked_value - input_mask_share) % p)

n = 4
t = 1
p = 52435875175126190479447740508185965837690552500527637822603658699938581184513
R = 10920338887063814464675503992315976177888879664585288394250266608035967270910
inverse_R = get_inverse(R)
fp = 1 << 16
sz = 32

display_precision = 4

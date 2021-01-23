import os
import time

import leveldb

from gmpy import binary, mpz
from gmpy2 import mpz_from_old_binary

##############################################################################
# keys                                                                       #
##############################################################################


def key_balance(token, user):
    return f"balance_{token}_{user}".encode()


def key_inputmask(idx):
    return f"inputmask_{idx}".encode()


def key_pool(token_A, token_B, token):
    return f"pool_{token_A}_{token_B}_{token}".encode()


def key_price(trade_seq):
    return f"price_{trade_seq}".encode()


def key_total_price(token_A, token_B):
    return f"trade_price_{token_A}_{token_B}".encode()


def key_total_supply(token_A, token_B):
    return f"total_supply_{token_A}_{token_B}".encode()


def key_trade_cnt(token_A, token_B):
    return f"trade_cnt_{token_A}-{token_B}".encode()


##############################################################################
# files                                                                      #
##############################################################################


def location_db(server_id):
    db_path = os.getenv("DB_PATH", "/opt/hbswap/db")
    return f"{db_path}/server-{server_id}"


def location_inputmask(server_id):
    inputmask_shares_dir = os.getenv(
        "INPUTMASK_SHARES", "/opt/hbswap/inputmask-shares",
    )
    return f"{inputmask_shares_dir}/4-MSp-255/Randoms-MSp-P{server_id}"


def location_private_output(server_id):
    prep_dir = os.getenv("PREP_DIR", "/opt/hbswap/preprocessing-data")
    return f"{prep_dir}/Private-Output-{server_id}"


def location_sharefile(server_id):
    return f"Persistence/Transactions-P{server_id}.data"


##############################################################################
# functions                                                                  #
##############################################################################


def openDB(location):
    while True:
        try:
            return leveldb.LevelDB(location)
        except leveldb.LevelDBError:
            time.sleep(3)


def get_value(db, key):
    try:
        return bytes(db.Get(key))
    except KeyError:
        return to_hex(0)


def get_inverse(a):
    ret = 1
    b = p - 2
    while b:
        if b % 2 == 1:
            ret = (ret * a) % p
        b //= 2
        a = (a * a) % p
    return ret


def from_float(x):
    return int(round(float(x) * (2 ** fp)))


def from_hex(x):
    return int((mpz_from_old_binary(x) * inverse_R) % p)


def to_hex(x):
    x = mpz(x)
    x = (x * R) % p
    x = binary(int(x))
    x += b"\x00" * (32 - len(x))
    return x


def check_consistency(shares):
    value = reconstruct(shares, t + 1)
    for i in range(t + 2, n + 1):
        if reconstruct(shares, i) != value:
            print("inconsistent")


def reconstruct(shares, n):
    inputmask = 0
    for i in range(1, n + 1):
        tot = 1
        for j in range(1, n + 1):
            if i == j:
                continue
            tot = tot * j * get_inverse(j - i) % p
        inputmask = (inputmask + shares[i - 1] * tot) % p
    return inputmask


n = 4
t = 1
p = 52435875175126190479447740508185965837690552500527637822603658699938581184513
R = 10920338887063814464675503992315976177888879664585288394250266608035967270910
inverse_R = get_inverse(R)
fp = 16
sz = 32

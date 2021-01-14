import os
import sys
import time

import leveldb

from utils import from_hex, sz, to_hex

input_parameter_num = 8

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    trade_seq = sys.argv[4]
    db_path = os.getenv("DB_PATH", "/opt/hbswap/db")

    file = f"Persistence/Transactions-P{server_id}.data"
    pool_A, pool_B = 0, 0
    change_A, change_B = 0, 0
    with open(file, "rb") as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)
        change_A = f.read(sz)
        change_B = f.read(sz)
        trade_price = f.read(sz)

    # file = f"Scripts/hbswap/data/Pool-{token_A}-{token_B}-P{server_id}.data"
    # with open(file, 'wb') as f:
    #     f.write(pool_A + pool_B)

    print(from_hex(change_A), from_hex(change_B))

    while True:
        try:
            db = leveldb.LevelDB(f"{db_path}/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    db.Put(f"pool-{token_A}-{token_B}:{token_A}".encode(), pool_A)
    db.Put(f"pool-{token_A}-{token_B}:{token_B}".encode(), pool_B)

    key = f"price_{trade_seq}".encode()
    print(key)
    db.Put(key, trade_price)

    key_price = f"trade_price_{token_A}-{token_B}".encode()
    try:
        total_price = from_hex(bytes(db.Get(key_price)))
    except KeyError:
        total_price = 0
    total_price += from_hex(trade_price)
    db.Put(key_price, to_hex(str(total_price)))

    key_cnt = f"trade_cnt_{token_A}-{token_B}".encode()
    try:
        total_cnt = from_hex(bytes(db.Get(key_cnt)))
    except KeyError:
        total_cnt = 0
    total_cnt += 1
    db.Put(key_cnt, to_hex(str(total_cnt)))

import os
import sys
import time

import leveldb

from utils import from_hex, to_hex, from_float, zero

if __name__ == "__main__":
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = sys.argv[4]
    flag = bool(int(sys.argv[5]))

    db_path = os.getenv("DB_PATH", "/opt/hbswap/db")

    while True:
        try:
            db = leveldb.LevelDB(f"{db_path}/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    key = f"balance{token}{user}".encode()
    try:
        balance = from_hex(bytes(db.Get(key)))
    except KeyError:
        balance = 0

    if balance == zero:
        balance = 0
    print("old balance", balance)
    balance += from_float(amt) if flag else int(amt)
    print("updated balance", balance)
    db.Put(key, to_hex(str(balance)))

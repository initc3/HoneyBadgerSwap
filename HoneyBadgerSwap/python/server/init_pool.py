import os
import sys
import time

import leveldb

from utils import to_hex, from_float

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    amt_A = sys.argv[4]
    amt_B = sys.argv[5]
    amt = sys.argv[6]

    db_path = os.getenv("DB_PATH", "/opt/hbswap/db")

    while True:
        try:
            db = leveldb.LevelDB(f"{db_path}/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    key_A = f"pool-{token_A}-{token_B}:{token_A}".encode()
    db.Put(key_A, to_hex(from_float(amt_A)))

    key_B = f"pool-{token_A}-{token_B}:{token_B}".encode()
    db.Put(key_B, to_hex(from_float(amt_B)))

    key = f"total_supply_{token_A}-{token_B}".encode()
    db.Put(key, to_hex(from_float(amt)))

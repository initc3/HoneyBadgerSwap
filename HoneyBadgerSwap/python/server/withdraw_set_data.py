import os
import sys
import time

import leveldb

from utils import to_hex, fp

if __name__ == "__main__":
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = to_hex(str(round(float(sys.argv[4]) * (2 ** fp))))

    db_path = os.getenv("DB_PATH", "/opt/hbswap/db")

    while True:
        try:
            db = leveldb.LevelDB(f"{db_path}/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    try:
        balance = bytes(db.Get(f"balance{token}{user}".encode()))
    except KeyError:
        balance = to_hex(str(1))

    file = f"Persistence/Transactions-P{server_id}.data"
    with open(file, "wb") as f:
        f.write(balance + amt)

import leveldb
import sys
import time

from utils import to_hex, fp

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    amt_A = sys.argv[4]
    amt_B = sys.argv[5]

    while True:
        try:
            db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    key_A = f'pool-{token_A}-{token_B}:{token_A}'.encode()
    pool_A = int(round(float(amt_A) * (2 ** fp)))
    db.Put(key_A, to_hex(str(pool_A)))

    key_B = f'pool-{token_A}-{token_B}:{token_B}'.encode()
    pool_B = int(round(float(amt_B) * (2 ** fp)))
    db.Put(key_B, to_hex(str(pool_B)))
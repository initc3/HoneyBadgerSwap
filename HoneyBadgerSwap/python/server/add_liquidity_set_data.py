import leveldb
import sys
import time

from utils import to_hex, from_float, from_hex

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    amt_A = to_hex(from_float(sys.argv[5]))
    amt_B = to_hex(from_float(sys.argv[6]))

    while True:
        try:
            db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    pool_A = bytes(db.Get(f'pool-{token_A}-{token_B}:{token_A}'.encode()))
    pool_B = bytes(db.Get(f'pool-{token_A}-{token_B}:{token_B}'.encode()))

    key = f'total_supply_{token_A}-{token_B}'.encode()
    total_supply = bytes(db.Get(key))
    print(from_hex(total_supply))

    file = f"Persistence/Transactions-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(pool_A + pool_B + amt_A + amt_B + total_supply)
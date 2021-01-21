import leveldb
import sys
import time

from utils import from_hex, sz, from_float, to_hex

input_parameter_num = 4

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    amt = sys.argv[4]

    file = f"Persistence/Transactions-P{server_id}.data"
    pool_A, pool_B = 0, 0
    with open(file, 'rb') as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)

    while True:
        try:
            db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    db.Put(f'pool-{token_A}-{token_B}:{token_A}'.encode(), pool_A)
    db.Put(f'pool-{token_A}-{token_B}:{token_B}'.encode(), pool_B)

    key = f'total_supply_{token_A}-{token_B}'.encode()
    total_supply = from_hex(bytes(db.Get(key)))
    total_supply -= from_float(amt)
    db.Put(key, to_hex(str(total_supply)))

    file = f"Player-Data/Private-Output-{server_id}"
    with open(file, 'rb') as f:
        amt_A = f.read(sz)
        amt_B = f.read(sz)
    print(from_hex(amt_A), from_hex(amt_B))

    import os
    print('filesize:', os.path.getsize(file))
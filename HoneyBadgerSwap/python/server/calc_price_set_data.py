import leveldb
import sys
import time

from utils import from_hex, to_hex

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]

    while True:
        try:
            db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
            break
        except leveldb.LevelDBError:
            time.sleep(3)

    key_price = f'trade_price_{token_A}-{token_B}'.encode()
    try:
        total_price = from_hex(bytes(db.Get(key_price)))
    except KeyError:
        total_price = 0

    key_cnt = f'trade_cnt_{token_A}-{token_B}'.encode()
    try:
        total_cnt = from_hex(bytes(db.Get(key_cnt)))
    except KeyError:
        total_cnt = 0

    file = f"Persistence/Transactions-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(to_hex(str(total_price)) + to_hex(str(total_cnt)))

    db.Put(key_price, to_hex(0))
    db.Put(key_cnt, to_hex(0))
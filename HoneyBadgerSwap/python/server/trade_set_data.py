import leveldb
import sys

from utils import to_hex, sz

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    idx_A = sys.argv[5]
    idx_B = sys.argv[6]
    masked_amt_A = to_hex(sys.argv[7])
    masked_amt_B = to_hex(sys.argv[8])

    file = f"Scripts/hbswap/data/Pool-P{server_id}.data"
    pool_A, pool_B = 0, 0
    with open(file, 'rb') as f:
        pool_A = f.read(sz)
        pool_B = f.read(sz)

    db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
    mask_share_A = bytes(db.Get(idx_A.encode()))
    mask_share_B = bytes(db.Get(idx_B.encode()))

    try:
        balance_A = bytes(db.Get(f'balance{token_A}{user}'.encode()))
    except KeyError:
        balance_A = to_hex(str(1))
    try:
        balance_B = bytes(db.Get(f'balance{token_B}{user}'.encode()))
    except KeyError:
        balance_B = to_hex(str(1))

    file = f"Persistence/Transactions-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(pool_A + pool_B + balance_A + balance_B + mask_share_A + mask_share_B + masked_amt_A + masked_amt_B)
import leveldb
import sys

from utils import to_hex, sz

if __name__=='__main__':
    server_id = sys.argv[1]
    idx_ETH = sys.argv[2]
    idx_TOK = sys.argv[3]
    masked_amt_ETH = to_hex(sys.argv[4])
    masked_amt_TOK = to_hex(sys.argv[5])

    file = f"Scripts/hbswap/data/Pool-P{server_id}.data"
    pool_ETH, pool_TOK = 0, 0
    with open(file, 'rb') as f:
        pool_ETH = f.read(sz)
        pool_TOK = f.read(sz)

    db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
    mask_share_ETH = bytes(db.Get(idx_ETH.encode()))
    mask_share_TOK = bytes(db.Get(idx_TOK.encode()))

    file = f"Persistence/Transactions-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(pool_ETH + pool_TOK + mask_share_ETH + mask_share_TOK + masked_amt_ETH + masked_amt_TOK)

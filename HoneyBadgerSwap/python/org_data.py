import sys

from utils import sz

if __name__=='__main__':
    server_id = sys.argv[1]

    file = f"Persistence/Transactions-P{server_id}.data"
    pool_ETH, pool_TOK = 0, 0
    with open(file, 'rb') as f:
        f.seek(6 * sz)
        pool_ETH = f.read(sz)
        pool_TOK = f.read(sz)
    file = f"Scripts/hbswap/data/Pool-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(pool_ETH + pool_TOK)
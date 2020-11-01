from utils import sz

if __name__=='__main__':
    n = 3

    for i in range(n):
        file = f"Persistence/Transactions-P{i}.data"
        pool_ETH, pool_TOK = 0, 0
        with open(file, 'rb') as f:
            f.seek(6 * sz)
            pool_ETH = f.read(sz)
            pool_TOK = f.read(sz)

        with open(file, 'wb') as f:
            f.write(pool_ETH + pool_TOK)
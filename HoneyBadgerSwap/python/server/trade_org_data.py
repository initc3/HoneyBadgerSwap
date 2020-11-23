import sys

from utils import from_hex, sz

input_parameter_num = 8

if __name__=='__main__':
    server_id = sys.argv[1]

    file = f"Persistence/Transactions-P{server_id}.data"
    pool_A, pool_B = 0, 0
    change_A, change_B = 0, 0
    with open(file, 'rb') as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)
        change_A = f.read(sz)
        change_B = f.read(sz)
    file = f"Scripts/hbswap/data/Pool-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(pool_A + pool_B)

    print(from_hex(change_A), from_hex(change_B))
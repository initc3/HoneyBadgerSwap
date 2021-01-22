import leveldb
import sys
import time

sys.path.insert(1, 'Scripts/hbswap/python')
from utils import from_hex, to_hex, sz

if __name__=='__main__':
    server_id = sys.argv[1]

    print(from_hex(to_hex(1123411)))

    # file = f"Persistence/Transactions-P{server_id}.data"
    #     # with open(file, 'wb') as f:
    #     #     f.write(to_hex(0))

    # file = f"Player-Data/Private-Output-{server_id}"
    # with open(file, 'rb') as f:
    #     zero = f.read(sz)
    #     print(zero)
    #     print(from_hex(zero))
    #
    # import os
    # print('filesize:', os.path.getsize(file))
import leveldb
import sys
import time

from utils import to_hex

if __name__=='__main__':
    server_id = sys.argv[1]

    file = f"Persistence/Transactions-P{server_id}.data"
    with open(file, 'wb') as f:
        f.write(to_hex(0))
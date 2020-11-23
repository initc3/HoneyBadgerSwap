import leveldb
import sys

from utils import from_hex, to_hex, fp

if __name__=='__main__':
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = int(sys.argv[4])
    flag = bool(int(sys.argv[5]))

    db = leveldb.LevelDB(f"Scripts/hbswap/db/server{server_id}")
    key = f'balance{token}{user}'.encode()
    try:
        balance = bytes(db.Get(key))
        balance = from_hex(balance)
    except KeyError:
        balance = 0

    print("old balance", balance)
    balance += int(round(float(amt) * (2 ** fp))) if flag else amt
    print("updated balance", balance)
    db.Put(key, to_hex(str(balance)))
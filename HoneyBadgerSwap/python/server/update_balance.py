import sys

sys.path.insert(1, "Scripts/hbswap/python")
from utils import (
    key_balance,
    location_db,
    openDB,
    get_value,
    from_float,
    from_hex,
    to_hex,
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = sys.argv[4]
    flag = bool(int(sys.argv[5]))

    db = openDB(location_db(server_id))

    key = key_balance(token, user)
    balance = from_hex(get_value(db, key))

    print("old balance:", balance)
    balance += from_float(amt) if flag else int(amt)
    print("updated balance:", balance)
    db.Put(key, to_hex(balance))

import sys

from ..utils import (
    key_balance,
    location_db,
    openDB,
    get_value,
    fix_to_float,
    hex_to_int,
    int_to_hex,
    p
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = int(sys.argv[4])

    db = openDB(location_db(server_id))
    key = key_balance(token, user)
    balance = (hex_to_int(get_value(db, key)) + amt) % p
    print(f'balance {fix_to_float(balance)}')
    db.Put(key, int_to_hex(balance))

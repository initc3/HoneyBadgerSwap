import sys

from ..utils import (
    key_balance,
    location_db,
    location_private_output,
    openDB,
    get_value,
    hex_to_int,
    int_to_hex,
    sz,
    p
)

if __name__=='__main__':
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = int(sys.argv[4]) # fix

    file = location_private_output(server_id)
    with open(file, 'rb') as f:
        enough = hex_to_int(f.read(sz))
    print(enough)

    if enough == 1:
        db = openDB(location_db(server_id))

        key = key_balance(token, user)
        balance = hex_to_int(get_value(db, key)) # fix
        balance = (balance - amt) % p
        db.Put(key, int_to_hex(balance))
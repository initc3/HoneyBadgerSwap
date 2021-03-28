import math
import sys

from ..utils import (
    key_balance,
    key_pool,
    key_total_supply,
    location_db,
    location_private_output,
    openDB,
    get_value,
    hex_to_int,
    int_to_hex,
    sz,
    p,
    display_precision,
)

input_parameter_num = 6

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    user = sys.argv[4]
    amt_A = int(sys.argv[5]) # fix
    amt_B = int(sys.argv[6]) # fix

    file = location_private_output(server_id)
    with open(file, 'rb') as f:
        valid_order = hex_to_int(f.read(sz))
    print(valid_order)

    if valid_order == 1:
        db = openDB(location_db(server_id))

        db.Put(key_pool(token_A, token_B, token_A), int_to_hex(amt_A))
        db.Put(key_pool(token_A, token_B, token_B), int_to_hex(amt_B))

        balance_A = hex_to_int(get_value(db, key_balance(token_A, user))) # fix
        balance_A = (balance_A - amt_A) % p
        db.Put(key_balance(token_A, user), int_to_hex(balance_A))

        balance_B = hex_to_int(get_value(db, key_balance(token_B, user))) # fix
        balance_B = (balance_B - amt_B) % p
        db.Put(key_balance(token_B, user), int_to_hex(balance_B))

        amt_LT = int(math.sqrt(amt_A * amt_B)) % p

        db.Put(key_balance(f'{token_A}+{token_B}', user), int_to_hex(amt_LT))

        db.Put(key_total_supply(token_A, token_B), int_to_hex(amt_LT))

        init_price = 1. * amt_B / amt_A
        print('{:.{}f}'.format(init_price, display_precision))


import sys

from ..utils import (
    key_balance,
    key_pool,
    key_total_supply,
    location_db,
    location_sharefile,
    location_private_output,
    openDB,
    hex_to_int,
    sz,
)

input_parameter_num = 7

if __name__ == "__main__":
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]

    file = location_sharefile(server_id)
    with open(file, "rb") as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)
        balance_A = f.read(sz)
        balance_B = f.read(sz)
        balance_LT = f.read(sz)
        total_supply_LT = f.read(sz)

    db = openDB(location_db(server_id))

    db.Put(key_pool(token_A, token_B, token_A), pool_A)
    db.Put(key_pool(token_A, token_B, token_B), pool_B)

    db.Put(key_balance(token_A, user), balance_A)
    db.Put(key_balance(token_B, user), balance_B)
    db.Put(key_balance(f'{token_A}+{token_B}', user), balance_LT)

    db.Put(key_total_supply(token_A, token_B), total_supply_LT)

    file = location_private_output(server_id)
    with open(file, 'rb') as f:
        zero_total_LT = hex_to_int(f.read(sz))
    print(zero_total_LT)

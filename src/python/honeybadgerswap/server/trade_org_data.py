import sys

from ..utils import (
    key_balance,
    key_pool,
    location_db,
    location_private_output,
    location_sharefile,
    openDB,
    hex_to_int,
    sz
)

input_parameter_num = 6

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]

    file = location_sharefile(server_id)
    with open(file, 'rb') as f:
        f.seek(input_parameter_num * sz)
        change_A = hex_to_int(f.read(sz)) # int
        change_B = hex_to_int(f.read(sz)) # int
        pool_A = f.read(sz) # hex
        pool_B = f.read(sz) # hex
        balance_A = hex_to_int(f.read(sz)) # int
        balance_B = hex_to_int(f.read(sz)) # int

    db = openDB(location_db(server_id))

    db.Put(key_pool(token_A, token_B, token_A), pool_A)
    db.Put(key_pool(token_A, token_B, token_B), pool_B)

    file = location_private_output(server_id)
    with open(file, 'rb') as f:
        order_succeed = hex_to_int(f.read(sz))
    print(order_succeed)

    print(balance_A)
    print(balance_B)
    print(change_B)
    print(change_A)

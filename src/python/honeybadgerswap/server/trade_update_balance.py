import sys

from ..utils import (
    key_balance,
    location_db,
    openDB,
    int_to_hex
)

input_parameter_num = 6

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    balance_A = int_to_hex(int(sys.argv[5])) # hex
    balance_B = int_to_hex(int(sys.argv[6])) # hex

    db = openDB(location_db(server_id))

    db.Put(key_balance(token_A, user), balance_A)
    db.Put(key_balance(token_B, user), balance_B)
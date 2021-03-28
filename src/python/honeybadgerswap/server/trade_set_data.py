import sys

from ..utils import (
    key_balance,
    key_pool,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    recover_input,
)

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    idx_A = sys.argv[5]
    masked_amt_A = int(sys.argv[6]) # fix
    idx_B = sys.argv[7]
    masked_amt_B = int(sys.argv[8]) # fix

    db = openDB(location_db(server_id))

    amt_A = recover_input(db, masked_amt_A, idx_A) # hex
    amt_B = recover_input(db, masked_amt_B, idx_B) # hex

    balance_A = get_value(db, key_balance(token_A, user)) # hex
    balance_B = get_value(db, key_balance(token_B, user)) # hex

    pool_A = get_value(db, key_pool(token_A, token_B, token_A)) # hex
    pool_B = get_value(db, key_pool(token_A, token_B, token_B)) # hex

    file = location_sharefile(server_id)
    with open(file, 'wb') as f:
        f.write(
            amt_A +
            amt_B +
            balance_B +
            pool_A +
            pool_B +
            balance_A
        )

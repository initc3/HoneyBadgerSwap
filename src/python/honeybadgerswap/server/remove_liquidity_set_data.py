import sys

from ..utils import (
    key_balance,
    key_pool,
    key_total_supply,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    recover_input
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    idx_LT = sys.argv[5]
    masked_amt_LT = int(sys.argv[6]) # fix

    db = openDB(location_db(server_id))

    balance_LT = get_value(db, key_balance(f'{token_A}+{token_B}', user)) # hex
    amt_LT = recover_input(db, masked_amt_LT, idx_LT) # hex

    pool_A = get_value(db, key_pool(token_A, token_B, token_A)) # hex
    pool_B = get_value(db, key_pool(token_A, token_B, token_B)) # hex

    total_supply_LT = get_value(db, key_total_supply(token_A, token_B)) # hex

    balance_A = get_value(db, key_balance(token_A, user)) # hex
    balance_B = get_value(db, key_balance(token_B, user)) # hex

    file = location_sharefile(server_id)
    with open(file, "wb") as f:
        f.write(
            balance_LT +
            amt_LT +
            pool_A +
            pool_B +
            total_supply_LT +
            balance_A +
            balance_B
        )

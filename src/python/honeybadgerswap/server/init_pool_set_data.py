import sys

from ..utils import (
    key_balance,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    int_to_hex,
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    amt_A = int(sys.argv[5]) # fix
    amt_B = int(sys.argv[6]) # fix

    db = openDB(location_db(server_id))

    balance_A = get_value(db, key_balance(token_A, user)) # hex
    balance_B = get_value(db, key_balance(token_B, user)) # hex
    total_supply_LT = get_value(db, key_balance(f'{token_A}+{token_B}', user)) # hex

    file = location_sharefile(server_id)
    with open(file, "wb") as f:
        f.write(
            balance_A +
            int_to_hex(amt_A) +
            balance_B +
            int_to_hex(amt_B) +
            total_supply_LT
        )

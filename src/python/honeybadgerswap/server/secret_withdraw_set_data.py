import sys

from ..utils import (
    key_balance,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    int_to_hex
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    user = sys.argv[2]
    token = sys.argv[3]
    amt = int(sys.argv[4]) # fix

    db = openDB(location_db(server_id))

    balance = get_value(db, key_balance(token, user)) # hex

    file = location_sharefile(server_id)
    with open(file, "wb") as f:
        f.write(
            balance +
            int_to_hex(amt)
        )

import sys

from ..utils import (
    key_balance,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    from_float,
    to_hex,
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    token = sys.argv[2]
    user = sys.argv[3]
    amt = to_hex(from_float(sys.argv[4]))

    db = openDB(location_db(server_id))
    balance = get_value(db, key_balance(token, user))

    file = location_sharefile(server_id)
    with open(file, "wb") as f:
        f.write(balance + amt)

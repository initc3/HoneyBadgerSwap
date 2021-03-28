import sys

from ..utils import (
    key_total_price,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    int_to_hex
)

if __name__=='__main__':
    server_id = sys.argv[1]
    change_B = int_to_hex(int(sys.argv[2])) # hex
    change_A = int_to_hex(int(sys.argv[3])) # hex
    token_A = sys.argv[4]
    token_B = sys.argv[5]

    db = openDB(location_db(server_id))

    total_price = get_value(db, key_total_price(token_A, token_B)) # hex

    file = location_sharefile(server_id)
    with open(file, 'wb') as f:
        f.write(
            change_B +
            change_A +
            total_price
        )
import sys

from ..utils import(
    key_total_price,
    key_total_cnt,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    fix_to_float,
    hex_to_int
)

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]

    db = openDB(location_db(server_id))

    total_price = get_value(db, key_total_price(token_A, token_B)) # hex

    total_cnt = get_value(db, key_total_cnt(token_A, token_B)) # hex
    cnt = fix_to_float(hex_to_int(total_cnt))
    print('cnt', cnt)

    file = location_sharefile(server_id)
    with open(file, 'wb') as f:
        f.write(
            total_price +
            total_cnt
        )
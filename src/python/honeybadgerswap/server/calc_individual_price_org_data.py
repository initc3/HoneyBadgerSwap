import sys

from ..utils import (
    key_individual_price,
    key_total_price,
    key_total_cnt,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    hex_to_int,
    int_to_hex,
    sz,
    fp
)

input_parameter_num = 3

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    trade_seq = sys.argv[4]

    file = location_sharefile(server_id)
    with open(file, "rb") as f:
        f.seek(input_parameter_num * sz)
        price = f.read(sz) # hex
        total_price = f.read(sz) # hex

    db = openDB(location_db(server_id))

    db.Put(key_total_price(token_A, token_B), total_price)

    k_total_cnt = key_total_cnt(token_A, token_B)
    total_cnt = get_value(db, k_total_cnt) # hex
    total_cnt = hex_to_int(total_cnt) + fp # int
    db.Put(k_total_cnt, int_to_hex(total_cnt))

    db.Put(key_individual_price(trade_seq), price)
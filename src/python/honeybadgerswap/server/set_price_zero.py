import sys

from ..utils import (
    key_individual_price,
    location_db,
    openDB,
    int_to_hex,
    get_value, key_cnt_failed_trade, hex_to_int)

input_parameter_num = 3

if __name__ == "__main__":
    server_id = sys.argv[1]
    trade_seq = sys.argv[2]

    db = openDB(location_db(server_id))

    db.Put(key_individual_price(trade_seq), int_to_hex(0))

    cnt_failed_trade = hex_to_int(get_value(db, key_cnt_failed_trade())) + 1
    print('cnt_failed_trade', cnt_failed_trade)
    db.Put(key_cnt_failed_trade(), int_to_hex(cnt_failed_trade))
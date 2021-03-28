import sys

from ..utils import (
    key_individual_price,
    location_db,
    openDB,
    int_to_hex,
)

input_parameter_num = 3

if __name__ == "__main__":
    server_id = sys.argv[1]
    trade_seq = sys.argv[2]

    db = openDB(location_db(server_id))

    db.Put(key_individual_price(trade_seq), int_to_hex(0))
import sys

sys.path.insert(1, "Scripts/hbswap/python")
from utils import (
    key_pool,
    key_price,
    key_total_price,
    key_trade_cnt,
    location_db,
    location_sharefile,
    openDB,
    get_value,
    from_hex,
    to_hex,
    sz,
)

input_parameter_num = 8

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    trade_seq = sys.argv[4]

    file = location_sharefile(server_id)
    with open(file, "rb") as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)
        change_A = from_hex(f.read(sz))
        change_B = from_hex(f.read(sz))
        trade_price = from_hex(f.read(sz))

    print(change_A, change_B)

    db = openDB(location_db(server_id))

    db.Put(key_pool(token_A, token_B, token_A), pool_A)
    db.Put(key_pool(token_A, token_B, token_B), pool_B)

    k_price = key_price(trade_seq)
    db.Put(k_price, k_price)

    k_total_price = key_total_price(token_A, token_B)
    total_price = from_hex(get_value(db, k_total_price))
    print("total_price before", total_price)
    total_price += trade_price
    print("total_price after", total_price)
    db.Put(k_total_price, to_hex(total_price))

    k_trade_cnt = key_trade_cnt(token_A, token_B)
    trade_cnt = from_hex(get_value(db, k_trade_cnt))
    print("total_cnt before", trade_cnt)
    trade_cnt += 1
    print("total_cnt after", trade_cnt)
    db.Put(k_trade_cnt, to_hex(trade_cnt))

import sys
import time

from ..utils import key_pool, key_trade_price, key_trade_time, key_total_price, key_total_cnt, location_db, location_sharefile, openDB, get_value, from_hex, to_hex, sz

input_parameter_num = 8

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    trade_seq = sys.argv[4]

    file = location_sharefile(server_id)
    with open(file, 'rb') as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)
        change_A = from_hex(f.read(sz))
        change_B = from_hex(f.read(sz))
        trade_price = f.read(sz)
        trade_cnt = from_hex(f.read(sz))

    print(change_A, change_B)

    db = openDB(location_db(server_id))

    db.Put(key_pool(token_A, token_B, token_A), pool_A)
    db.Put(key_pool(token_A, token_B, token_B), pool_B)

    k_trade_price = key_trade_price(trade_seq)
    db.Put(k_trade_price, trade_price)
    k_trade_time = key_trade_time(trade_seq)
    print(k_trade_time)
    db.Put(k_trade_time, to_hex(int(time.time())))

    k_total_price = key_total_price(token_A, token_B)
    total_price = from_hex(get_value(db, k_total_price))
    print('total_price before', total_price)
    total_price += from_hex(trade_price)
    print('total_price after', total_price)
    db.Put(k_total_price, to_hex(total_price))

    k_total_cnt = key_total_cnt(token_A, token_B)
    total_cnt = from_hex(get_value(db, k_total_cnt))
    print('total_cnt before', total_cnt)
    total_cnt += trade_cnt
    print('total_cnt after', total_cnt)
    db.Put(k_total_cnt, to_hex(total_cnt))

    print("pool_A", from_hex(pool_A))
    print("pool_B", from_hex(pool_B))
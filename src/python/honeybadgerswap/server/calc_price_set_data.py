import sys

from ..utils import key_total_price, key_total_cnt, location_db, location_sharefile, openDB, get_value

if __name__=='__main__':
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]

    db = openDB(location_db(server_id))

    k_total_price = key_total_price(token_A, token_B)
    total_price = get_value(db, k_total_price)

    k_total_cnt = key_total_cnt(token_A, token_B)
    total_cnt = get_value(db, k_total_cnt)

    file = location_sharefile(server_id)
    with open(file, 'wb') as f:
        f.write(total_price + total_cnt)

    db.Delete(k_total_price)
    db.Delete(k_total_cnt)
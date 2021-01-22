import sys

sys.path.insert(1, 'Scripts/hbswap/python')
from utils import key_pool, key_total_supply, location_db, location_sharefile, openDB, from_float, from_hex, to_hex

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    amt_A = to_hex(from_float(sys.argv[5]))
    amt_B = to_hex(from_float(sys.argv[6]))

    db = openDB(location_db(server_id))

    pool_A = bytes(db.Get(key_pool(token_A, token_B, token_A)))
    pool_B = bytes(db.Get(key_pool(token_A, token_B, token_B)))

    key = key_total_supply(token_A, token_B)
    total_supply = bytes(db.Get(key))
    print(from_hex(total_supply))

    file = location_sharefile(server_id)
    with open(file, 'wb') as f:
        f.write(pool_A + pool_B + amt_A + amt_B + total_supply)
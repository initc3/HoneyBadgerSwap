import sys

sys.path.insert(1, "Scripts/hbswap/python")
from utils import key_pool, key_total_supply, location_db, openDB, from_float, to_hex

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    amt_A = to_hex(from_float(sys.argv[4]))
    amt_B = to_hex(from_float(sys.argv[5]))
    amt_liquidity = to_hex(from_float(sys.argv[6]))

    db = openDB(location_db(server_id))

    key_A = key_pool(token_A, token_B, token_A)
    db.Put(key_A, amt_A)

    key_B = key_pool(token_A, token_B, token_B)
    db.Put(key_B, amt_B)

    key_liquidity = key_total_supply(token_A, token_B)
    db.Put(key_liquidity, amt_liquidity)

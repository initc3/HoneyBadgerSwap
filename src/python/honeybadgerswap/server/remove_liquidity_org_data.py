import sys

from ..utils import (
    key_pool,
    key_total_supply,
    location_db,
    location_sharefile,
    openDB,
    from_float,
    from_hex,
    to_hex,
    sz,
)

input_parameter_num = 4

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]
    amt_liquidity = from_float(sys.argv[4])

    file = location_sharefile(server_id)
    with open(file, "rb") as f:
        f.seek(input_parameter_num * sz)
        pool_A = f.read(sz)
        pool_B = f.read(sz)
        amt_A = from_hex(f.read(sz))
        amt_B = from_hex(f.read(sz))

    db = openDB(location_db(server_id))

    db.Put(key_pool(token_A, token_B, token_A), pool_A)
    db.Put(key_pool(token_A, token_B, token_B), pool_B)

    key = key_total_supply(token_A, token_B)
    total_supply = from_hex(bytes(db.Get(key)))
    total_supply -= amt_liquidity
    db.Put(key, to_hex(total_supply))

    print(amt_A, amt_B)

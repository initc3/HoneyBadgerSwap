import sys

from ..utils import(
    key_total_price,
    key_total_cnt,
    location_db,
    location_private_output,
    openDB,
    fix_to_float,
    hex_to_int,
    sz,
    display_precision
)

if __name__ == "__main__":
    server_id = sys.argv[1]
    token_A = sys.argv[2]
    token_B = sys.argv[3]

    db = openDB(location_db(server_id))

    db.Delete(key_total_price(token_A, token_B))
    db.Delete(key_total_cnt(token_A, token_B))

    file = location_private_output(server_id)
    with open(file, "rb") as f:
        batch_price = fix_to_float(hex_to_int(f.read(sz)))
    print('{:.{}f}'.format(batch_price, display_precision))

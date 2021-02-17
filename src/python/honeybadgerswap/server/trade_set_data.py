import sys

from ..utils import key_balance, key_inputmask, key_pool, location_db, location_sharefile, openDB, get_value, to_hex, from_hex

if __name__=='__main__':
    server_id = sys.argv[1]
    user = sys.argv[2]
    token_A = sys.argv[3]
    token_B = sys.argv[4]
    idx_A = sys.argv[5]
    idx_B = sys.argv[6]
    masked_amt_A = to_hex(sys.argv[7])
    masked_amt_B = to_hex(sys.argv[8])

    db = openDB(location_db(server_id))

    pool_A = bytes(db.Get(key_pool(token_A, token_B, token_A)))
    pool_B = bytes(db.Get(key_pool(token_A, token_B, token_B)))

    mask_share_A = bytes(db.Get(key_inputmask(idx_A)))
    mask_share_B = bytes(db.Get(key_inputmask(idx_B)))

    balance_A = get_value(db, key_balance(token_A, user))
    balance_B = get_value(db, key_balance(token_B, user))

    print("pool_A", from_hex(pool_A))
    print("pool_B", from_hex(pool_B))
    print("balance_A", from_hex(balance_A))
    print("balance_B", from_hex(balance_B))
    print("mask_share_A", from_hex(mask_share_A))
    print("mask_share_B", from_hex(mask_share_B))
    print("masked_amt_A", from_hex(masked_amt_A))
    print("masked_amt_B", from_hex(masked_amt_B))

    file = location_sharefile(server_id)
    with open(file, 'wb') as f:
        f.write(pool_A + pool_B + balance_A + balance_B + mask_share_A + mask_share_B + masked_amt_A + masked_amt_B)

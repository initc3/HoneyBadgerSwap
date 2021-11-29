import sys

from ..utils import key_inputmask, location_db, location_inputmask, openDB, to_hex, p

if __name__ == "__main__":
    server_id = sys.argv[1]
    init_idx = int(sys.argv[2])

    db = openDB(location_db(server_id))

    file = location_inputmask(server_id)
    with open(file, "r") as f:
        idx = init_idx
        for line in f.readlines():
            key = key_inputmask(idx)
            share = to_hex(int(line) % p)
            db.Put(key, share)
            idx += 1

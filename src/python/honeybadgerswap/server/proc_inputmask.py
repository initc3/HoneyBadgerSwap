import logging
import sys

from ..utils import key_inputmask, location_db, location_inputmask, openDB, to_hex, p

logger = logging.getLogger(name=__name__)

if __name__ == "__main__":
    server_id = sys.argv[1]
    init_idx = int(sys.argv[2])

    logger.info("open db to write inputmasks ...")
    db = openDB(location_db(server_id))
    logger.info("db has been opened")

    file = location_inputmask(server_id)
    with open(file, "r") as f:
        idx = init_idx
        for line in f.readlines():
            key = key_inputmask(idx)
            share = to_hex(int(line) % p)
            logger.info(
                f"will write to db: inputmasks {idx}, key: {key}, share: {share}"
            )
            db.Put(key, share)
            idx += 1
            logger.info(f"inputmasks {idx}, key: {key}, share: {share} written to db")

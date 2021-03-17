import time

from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware

from .config import settings
from ..utils import (
    key_balance,
    key_inputmask,
    key_trade_price,
    key_trade_time,
    location_db,
    openDB,
    get_value,
    from_hex,
)


async def db_get(key):
    db = openDB(location_db(settings.NODE_ID))
    return from_hex(get_value(db, key))


async def db_get_non_balance(key):
    while True:
        v = await db_get(key)
        if v == 0:
            raise HTTPException(status_code=404, detail=f"Key {key} not found")
        else:
            return v


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins="*",
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/")
async def root():
    return {"message": "Hello HoneyBadgerSwap!"}


@app.get("/info")
async def info():
    return {"info": "hbswap http server"}


@app.get("/inputmasks/{mask_idxes}")
async def get_inputmasks(mask_idxes: str):
    print(f"s{settings.NODE_ID} processing request GET /inputmasks/{mask_idxes}")
    _mask_idxes = mask_idxes.split(",")
    res = ""
    for mask_idx in _mask_idxes:
        res += (
            f"{',' if len(res) > 0 else ''}"
            f"{await db_get_non_balance(key_inputmask(mask_idx))}"
        )
    data = {"inputmask_shares": res}
    print(f"s{settings.NODE_ID} response to GET /inputmasks/{mask_idxes}: {res}")
    return data


@app.get("/price/{trade_seq}")
async def get_price(trade_seq: str):
    print(f"s{settings.NODE_ID} processing request GET /price/{trade_seq}")
    cur_time = int(time.time())
    prev_time = await db_get_non_balance(key_trade_time(trade_seq))
    passed_time = cur_time - prev_time
    time.sleep(max(0, 10 - passed_time))

    res = await db_get_non_balance(key_trade_price(trade_seq))
    data = {"price": f"{res}"}
    print(f"s{settings.NODE_ID} response to GET /price/{trade_seq}: {res}")
    return data


@app.get("/balance/{token_user}")
async def get_balance(token_user: str):
    print(f"s{settings.NODE_ID} processing request GET /balance/{token_user}")
    token, user = token_user.split(",")
    res = await db_get(key_balance(token, user))
    data = {"balance": f"{res}"}
    print(f"s{settings.NODE_ID} response to GET /balance/{token_user}: {res}")
    return data


@app.get("/log/{n}")
async def get_logs(n: int):
    print(f"s{settings.NODE_ID} processing request GET /log/{n}")
    log_file = open(f"/usr/src/hbswap/log/mpc_server_{settings.NODE_ID}.log", "r")
    lines = log_file.readlines()
    last_lines = lines[-n:]
    res = ""
    for line in last_lines:
        res += line
    data = {"log": f"{res}"}
    print(f"s{settings.NODE_ID} response to GET /log/{n}: {res}")
    return data


def start_server():
    import argparse
    import uvicorn

    parser = argparse.ArgumentParser("Start HoneyBadgerSwap HTTP server.")
    parser.add_argument(
        "--log-level",
        type=str,
        choices=("critical", "error", "warning", "info", "debug", "trace"),
        default="info",
        help="set log level",
    )
    args = parser.parse_args()

    uvicorn.run(
        "honeybadgerswap.server.main:app",
        host=settings.Servers[settings.NODE_ID]["HttpHost"],
        port=settings.Servers[settings.NODE_ID]["HttpPort"],
        log_level=args.log_level,
    )


if __name__ == "__main__":
    start_server()

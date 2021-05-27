import sys

from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
#
from .config import settings
# from ..utils import (
#     key_balance,
#     key_inputmask,
#     key_individual_price,
#     location_db,
#     openDB,
#     get_value,
#     hex_to_int
# )
from ratel.src.python.utils import openDB, location_db, key_inputmask

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
    print(f"s{serverID} processing request GET /inputmasks/{mask_idxes}")
    mask_idx_list = mask_idxes.split(",")
    res = ""
    print(mask_idx_list)
    db = openDB(location_db(serverID))
    for mask_idx in mask_idx_list:
        try:
            share = int.from_bytes(bytes(db.Get(key_inputmask(mask_idx))), 'big')
        except KeyError:
            print('key error: ', mask_idx)
            res = ''
            break
        res += (
            f"{',' if len(res) > 0 else ''}{share}"
        )
    data = {"inputmask_shares": res}
    print(f"s{serverID} response to GET /inputmasks/{mask_idxes}: {res}")
    return data


# @app.get("/price/{trade_seq}")
# async def get_price(trade_seq: str):
#     print(f"s{settings.NODE_ID} processing request GET /price/{trade_seq}")
#     db = openDB(location_db(settings.NODE_ID))
#     res = ''
#     try:
#         res = hex_to_int(bytes(db.Get(key_individual_price(trade_seq))))
#     except KeyError:
#         pass
#     data = {"price": f"{res}"}
#     print(f"s{settings.NODE_ID} response to GET /price/{trade_seq}: {res}")
#     return data
#
#
# @app.get("/balance/{token_user}")
# async def get_balance(token_user: str):
#     print(f"s{settings.NODE_ID} processing request GET /balance/{token_user}")
#     token, user = token_user.split(",")
#     db = openDB(location_db(settings.NODE_ID))
#     res = hex_to_int(get_value(db, key_balance(token, user)))
#     data = {"balance": f"{res}"}
#     print(f"s{settings.NODE_ID} response to GET /balance/{token_user}: {res}")
#     return data
#
#
# @app.get("/log/{n}")
# async def get_logs(n: int):
#     print(f"s{settings.NODE_ID} processing request GET /log/{n}")
#     log_file = open(f"/usr/src/hbswap/log/mpc_server_{settings.NODE_ID}.log", "r")
#     lines = log_file.readlines()
#     last_lines = lines[-n:]
#     res = ""
#     for line in last_lines:
#         res += line
#     data = {"log": f"{res}"}
#     print(f"s{settings.NODE_ID} response to GET /log/{n}: {res}")
#     return data


def start_server():
    # import argparse
    import uvicorn

    # parser = argparse.ArgumentParser("Start HoneyBadgerSwap HTTP server.")
    # parser.add_argument(
    #     "--log-level",
    #     type=str,
    #     choices=("critical", "error", "warning", "info", "debug", "trace"),
    #     default="info",
    #     help="set log level",
    # )
    # args = parser.parse_args()

    uvicorn.run(
        "ratel.src.python.httpserver:app",
        host="0.0.0.0",
        port=4000 + serverID,
    )


if __name__ == "__main__":
    serverID = int(sys.argv[1])

    start_server()

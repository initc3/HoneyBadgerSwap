import aiohttp_cors
import asyncio
import re
import time

from aiohttp import web

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


class Server:
    def __init__(self, n, t, server_id, host, http_port):
        self.n = n
        self.t = t
        self.server_id = server_id

        self.host = host
        self.http_port = http_port

        print(f"http server {server_id} is running...")

    async def db_get(self, key):
        db = openDB(location_db(self.server_id))
        return from_hex(get_value(db, key))

    async def db_get_non_balance(self, key):
        while True:
            v = await self.db_get(key)
            if v == 0:
                print(f"{key} not ready. Try again...")
                time.sleep(10)
            else:
                return v

    async def http_server(self):
        async def handler_info(request):
            data = {
                "info": "hbswap http server",
            }
            return web.json_response(data)

        async def handler_inputmask(request):
            print(f"s{self.server_id} request: {request}")
            mask_idxes = re.split(",", request.match_info.get("mask_idxes"))
            res = ""
            for mask_idx in mask_idxes:
                res += (
                    f"{',' if len(res) > 0 else ''}"
                    f"{await self.db_get_non_balance(key_inputmask(mask_idx))}"
                )
            data = {
                "inputmask_shares": res,
            }
            print(f"s{self.server_id} response: {res}")
            return web.json_response(data)

        async def handler_price(request):
            print(f"s{self.server_id} request: {request}")
            trade_seq = request.match_info.get("trade_seq")

            cur_time = int(time.time())
            prev_time = await self.db_get_non_balance(key_trade_time(trade_seq))
            passed_time = cur_time - prev_time
            time.sleep(max(0, 10 - passed_time))

            res = await self.db_get_non_balance(key_trade_price(trade_seq))
            data = {
                "price": f"{res}",
            }
            print(f"s{self.server_id} response: {res}")
            return web.json_response(data)

        async def handler_balance(request):
            print(f"s{self.server_id} request: {request}")
            token_user = re.split(",", request.match_info.get("token_user"))
            token = token_user[0]
            user = token_user[1]
            res = await self.db_get(key_balance(token, user))
            data = {
                "balance": f"{res}",
            }
            print(f"s{self.server_id} response: {res}")
            return web.json_response(data)

        async def handler_log(request):
            print(f"s{self.server_id} request: {request}")
            n = int(request.match_info.get("n"))
            log_file = open(f"/usr/src/hbswap/log/mpc_server_{self.server_id}.log", "r")
            lines = log_file.readlines()
            last_lines = lines[-n:]
            res = ""
            for line in last_lines:
                res += line
            data = {
                "log": f"{res}",
            }
            print(f"s{self.server_id} response: {res}")
            return web.json_response(data)

        app = web.Application()

        cors = aiohttp_cors.setup(
            app,
            defaults={
                "*": aiohttp_cors.ResourceOptions(
                    allow_credentials=True, expose_headers="*", allow_headers="*",
                )
            },
        )

        resource = cors.add(app.router.add_resource("/info"))
        cors.add(resource.add_route("GET", handler_info))
        resource = cors.add(app.router.add_resource("/inputmasks/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_inputmask))
        resource = cors.add(app.router.add_resource("/price/{trade_seq}"))
        cors.add(resource.add_route("GET", handler_price))
        resource = cors.add(app.router.add_resource("/balance/{token_user}"))
        cors.add(resource.add_route("GET", handler_balance))
        resource = cors.add(app.router.add_resource("/log/{n}"))
        cors.add(resource.add_route("GET", handler_log))

        runner = web.AppRunner(app)
        await runner.setup()
        site = web.TCPSite(runner, host=self.host, port=self.http_port)
        await site.start()
        await asyncio.sleep(100 * 3600)

import asyncio
import re
import time

from aiohttp import web

from ..utils import location_db, openDB, from_hex


class Server:
    def __init__(self, n, t, server_id, host, http_port):
        self.n = n
        self.t = t
        self.server_id = server_id

        self.host = host
        self.http_port = http_port

        print(f"http server {server_id} is running...")

    def dbGet(self, key):
        while True:
            try:
                db = openDB(location_db(self.server_id))
                return from_hex(bytes(db.Get(key)))
            except:
                print(f"Inputmask share {key} not ready. Try again...")
                time.sleep(5)

    async def http_server(self):
        routes = web.RouteTableDef()

        @routes.get("/inputmasks/{mask_idxes}")
        # async def inputmask_handler(request):
        async def _handler(request):
            print(f"request: {request}")
            mask_idxes = re.split(",", request.match_info.get("mask_idxes"))
            res = ""
            for mask_idx in mask_idxes:
                res += f"{',' if len(res) > 0 else ''}{self.dbGet(f'inputmask_{mask_idx}'.encode())}"
            data = {
                "inputmask_shares": res,
            }
            print(f"request: {request}")
            print(f"response: {res}")
            return web.json_response(data)

        @routes.get("/balance/{token_user}")
        # async def balance_handler(request):
        async def _handler(request):
            print(f"request: {request}")
            token_user = request.match_info.get("token_user")
            res = self.dbGet(f"balance{token_user}".encode())
            data = {
                "balance": f"{res}",
            }
            print(f"request: {request}")
            print(f"response: {res}")
            return web.json_response(data)

        @routes.get("/price/{trade_seq}")
        # async def price_handler(request):
        async def _handler(request):
            print(f"request: {request}")
            trade_seq = request.match_info.get("trade_seq")
            res = self.dbGet(f"price_{trade_seq}".encode())
            data = {
                "price": f"{res}",
            }
            print(f"request: {request}")
            print(f"response: {res}")
            return web.json_response(data)

        app = web.Application()
        app.add_routes(routes)
        runner = web.AppRunner(app)
        await runner.setup()
        site = web.TCPSite(runner, host=self.host, port=self.http_port)
        await site.start()
        await asyncio.sleep(100 * 3600)

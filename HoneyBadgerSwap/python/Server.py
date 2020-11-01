import asyncio
import leveldb
import re
import time

from aiohttp import web
from utils import inverse_R, p

from gmpy2 import mpz_from_old_binary

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
                db = leveldb.LevelDB(f"Scripts/hbswap/db/server{self.server_id}")
                value = bytes(db.Get(key))
                value = (mpz_from_old_binary(value) * inverse_R) % p
                return value
            except:
                print(f"Inputmask share {key} not ready. Try again...")
                time.sleep(1)

    async def http_server(self):
        routes = web.RouteTableDef()

        @routes.get("/inputmasks/{mask_idxes}")
        async def _handler(request):
            print(f"request: {request}")
            mask_idxes = re.split(',', request.match_info.get("mask_idxes"))
            res = ''
            for mask_idx in mask_idxes:
                res += f"{',' if len(res) > 0 else ''}{self.dbGet(mask_idx.encode())}"
            data = {
                "inputmask_shares": res,
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
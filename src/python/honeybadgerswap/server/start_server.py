import asyncio
import os
import sys
import toml

from .Server import Server


async def main(node_id, config_file):
    config = toml.load(config_file)

    n = config["n"]
    t = config["t"]

    server_config = config["servers"][node_id]
    server = Server(
        n, t, server_id, server_config["http_host"], server_config["http_port"]
    )

    tasks = []
    tasks.append(asyncio.ensure_future(server.http_server()))

    for task in tasks:
        await task


# TODO Use argparse to accept config file to use.
if __name__ == "__main__":
    server_id = int(sys.argv[1])
    config_file = os.getenv("HBSWAP_SERVER_CONFIG", "/opt/hbswap/conf/server.toml")
    asyncio.run(main(server_id, config_file))

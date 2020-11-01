import asyncio
import sys
import toml

from Server import Server

async def main(node_id, config_file):
    config = toml.load(config_file)

    n = config['n']
    t = config['t']

    server_config = config["servers"][node_id]
    server = Server(n, t, server_id, server_config["host"], server_config["http_port"])

    tasks = []
    tasks.append(asyncio.ensure_future(server.http_server()))

    for task in tasks:
        await task

if __name__ == "__main__":
    server_id = int(sys.argv[1])
    config_file = "Scripts/hbswap/conf/config.toml"
    asyncio.run(main(server_id, config_file))
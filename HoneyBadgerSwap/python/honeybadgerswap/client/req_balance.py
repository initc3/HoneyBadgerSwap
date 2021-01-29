import asyncio
import os
import sys

from .Client import Client

from ..utils import fp

if __name__ == "__main__":
    token = sys.argv[1]
    user = sys.argv[2]

    config = os.getenv("HBSWAP_CLIENT_CONFIG", "/opt/hbswap/conf/client.toml")
    client = Client.from_toml_config(config)

    balance = asyncio.run(client.get_balance(token, user))
    balance = 1.0 * balance / (2 ** fp)
    print(balance)

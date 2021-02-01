import asyncio
import os
import sys

from Client import Client

from ..utils import fp

if __name__ == "__main__":
    seq = sys.argv[1]

    config = os.getenv("HBSWAP_CLIENT_CONFIG", "/opt/hbswap/conf/client.toml")
    client = Client.from_toml_config(config)

    price = asyncio.run(client.get_price(seq))
    price = 1.0 * price / (2 ** fp)
    print(price)

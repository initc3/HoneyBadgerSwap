import asyncio
import sys

from Client import Client
from utils import fp

if __name__ == "__main__":
    seq = sys.argv[1]

    client = Client.from_toml_config("Scripts/hbswap/conf/config.toml")

    price = asyncio.run(client.get_price(seq))
    price = 1. * price / (2 ** fp)
    print(price)
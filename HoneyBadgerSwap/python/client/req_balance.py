import asyncio
import sys

from Client import Client
from utils import fp

if __name__ == "__main__":
    token = sys.argv[1]
    user = sys.argv[2]

    client = Client.from_toml_config("Scripts/hbswap/conf/config.toml")

    balance = asyncio.run(client.get_balance(token, user))
    balance = 1. * balance / (2 ** fp)
    print(balance)
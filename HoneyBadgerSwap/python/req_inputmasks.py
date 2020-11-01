import asyncio
import sys

from Client import Client
from utils import f

if __name__ == "__main__":
    client = Client.from_toml_config("Scripts/hbswap/conf/config.toml")
    idx_ETH, idx_TOK = sys.argv[1], sys.argv[2]
    inputmasks = asyncio.run(client.get_inputmasks(f'{idx_ETH},{idx_TOK}'))
    amt_ETH, amt_TOK = int(round(float(sys.argv[3]) * (2 ** f))), int(round(float(sys.argv[4]) * (2 ** f)))
    masked_amt_ETH = amt_ETH + inputmasks[0]
    masked_amt_TOK = amt_TOK + inputmasks[1]
    print(masked_amt_ETH, masked_amt_TOK)

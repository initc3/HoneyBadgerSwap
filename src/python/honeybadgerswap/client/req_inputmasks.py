import asyncio
import os
import sys

from .Client import Client

from ..utils import float_to_fix

if __name__ == "__main__":
    config = os.getenv("HBSWAP_CLIENT_CONFIG", "/opt/hbswap/conf/client.toml")
    client = Client.from_toml_config(config)

    inputmask_indexes = ""
    for inputmask_index in sys.argv[1::2]:
        inputmask_indexes += f"{inputmask_index},"
    inputmask_indexes = inputmask_indexes[:-1]

    inputmasks = asyncio.run(client.get_inputmasks(inputmask_indexes))

    values = []
    for value in sys.argv[2::2]:
        values.append(float_to_fix(value))

    masked_values = []
    for (inputmask, value) in zip(inputmasks, values):
        masked_values.append(value + inputmask)

    out = ""
    for masked_value in masked_values:
        out += f"{masked_value} "
    print(out)

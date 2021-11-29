import sys

from ..utils import location_private_output, from_hex, sz

if __name__ == "__main__":
    server_id = sys.argv[1]

    file = location_private_output(server_id)

    with open(file, "rb") as f:
        enough_balance = from_hex(f.read(sz))

    print(enough_balance)

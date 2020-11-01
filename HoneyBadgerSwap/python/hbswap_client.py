# python3 Scripts/hbswap_client.py

from gmpy2 import mpz_from_old_binary
from gmpy import binary
import sys

from utils import get_inverse, n, p, R, f, sz, inverse_R

def format(x):
    if x >= p / 2:
        x -= p
    return x

def to_hex(x):
    x = (x * R) % p
    return binary(int(x))

def read(f, start_posn):
    f.seek(start_posn)
    _x = f.read(sz)
    x = (mpz_from_old_binary(_x) * inverse_R) % p
    return format(x)

def get_inputmask(idx):
    x, y = [], []

    for i in range(n):
        file = f"Persistence/Transactions-P{i}.data"
        with open(file, 'rb') as f:
            x.append(i + 1)
            y.append(read(f, idx * sz))

    inputmask = 0
    for i in range(n):
        tot = 1
        for j in range(n):
            if i == j:
                continue
            tot = tot * x[j] * get_inverse(x[j] - x[i]) % p
        inputmask = (inputmask + y[i] * tot) % p

    print(inputmask)
    return inputmask

def append_share(x):
    for i in range(n):
        file = f"Persistence/Transactions-P{i}.data"
        with open(file, 'ab') as f:
            f.write(to_hex(x))

if __name__=='__main__':
    amt_ETH = int(round(float(sys.argv[1]) * (2 ** f)))
    amt_TOK = int(round(float(sys.argv[2]) * (2 ** f)))

    masked_amt_ETH = amt_ETH + get_inputmask(2)
    masked_amt_TOK = amt_TOK + get_inputmask(3)

    append_share(masked_amt_ETH)
    append_share(masked_amt_TOK)
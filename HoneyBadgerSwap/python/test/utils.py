from gmpy import binary, mpz
from gmpy2 import mpz_from_old_binary

def get_inverse(a):
    ret = 1
    b = p - 2
    while b:
        if b % 2 == 1:
            ret = (ret * a) % p
        b //= 2
        a = (a * a) % p
    return ret

def to_hex(x):
    x = mpz(x)
    x = (x * R) % p
    ret = binary(int(x))
    return ret + b"0" * (32 - len(ret))

def from_hex(x):
    return (mpz_from_old_binary(x) * inverse_R) % p

def reconstruct(shares, n):
    inputmask = 0
    for i in range(1, n + 1):
        tot = 1
        for j in range(1, n + 1):
            if i == j:
                continue
            tot = tot * j * get_inverse(j - i) % p
        inputmask = (inputmask + shares[i - 1] * tot) % p
    return inputmask

def check_consistency(shares):
    value = reconstruct(shares, t + 1)
    print(value)
    for i in range(t + 2, n + 1):
        if reconstruct(shares, i) != value:
            print('inconsistent')

t = 1
n = 4
# p = 57896044618658097711785492504343953926634992332820282019728792003956566065153
p = 52435875175126190479447740508185965837690552500527637822603658699938581184513
# R = 57896044618658097711785492504343953926634992332820282019728792003956563574783
R = 10920338887063814464675503992315976177888879664585288394250266608035967270910
fp = 16
sz = 32

inverse_R = get_inverse(R)
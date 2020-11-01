from gmpy import binary, mpz

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
    return binary(int(x))

n = 3
p = 57896044618658097711785492504343953926634992332820282019728792003956566065153
R = 57896044618658097711785492504343953926634992332820282019728792003956563574783
f = 16
sz = 32

inverse_R = get_inverse(R)
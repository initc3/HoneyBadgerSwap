from ..utils import reconstruct, n

# def reconstruct(shares, n):
#     inputmask = 0
#     for i in range(1, n + 1):
#         tot = 1
#         for j in range(1, n + 1):
#             if i == j:
#                 continue
#             tot = tot * j * get_inverse(j - i) % p
#         inputmask = (inputmask + shares[i - 1] * tot) % p
#     return inputmask
#
# def check(shares):
#     value = reconstruct(shares, t + 1)
#     for i in range(t + 2, n + 1):
#         if reconstruct(shares, i) != value:
#             print('ohoh')

# mask = []
# for server_id in range(4):
#     file = f'Player-Data/4-MSp-255/Randoms-MSp-P{server_id}'
#     tmp = []
#     with open(file, 'r') as f:
#         for line in f.readlines():
#             data = int(line) % p
#             tmp.append(data)
#     mask.append(tmp)
#
# for k in range(0, len(mask[0])):
#     shares = []
#     for i in range(n):
#         shares.append(mask[i][k])
#
shares = [
    17783221102966205198584993843005729747643599811630564308773575716113800647019,
    35566442205932410397169987686011459495287199623261128617547151432227600310998,
    913788133772425116307241020831223405240246934364055103717068448402818790464,
    18697009236738630314892234863836953152883846745994619412490644164516618454443,
]
print(reconstruct(shares, n))

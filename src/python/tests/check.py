from honeybadgerswap.utils import check_consistency

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
    -6937119568177438155937978633434279269014265249325731686894819606448044537196,
    -13874239136354876311875957266868558538028530498651463373789639212896089729752,
    -20811358704532314467813935900302837807042795747977195060684458819344134922308,
    24687396902416437855695825974448848761633491503224711075024380274146401069649,
]
print(check_consistency(shares))

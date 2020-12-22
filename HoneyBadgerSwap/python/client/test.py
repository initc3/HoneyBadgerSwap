from utils import get_inverse, n, p, to_hex

n = 4
x, y = [1,2,3,4], \
       [18457106183554705880258410793484639730509589330547761375057804664270278923797,
        13349002587220538934323108005261136798800285149272711771854820196514835822696,
        8240898990886371988387805217037633867090980967997662168651835728759392721595,
        3132795394552205042452502428814130935381676786722612565448851261003949620494]

# for i in range(n):
#     file = f"Persistence/Transactions-P{i}.data"
#     with open(file, 'rb') as f:
#         x.append(i + 1)
#         y.append(read(f, idx * sz))

inputmask = 0
for i in range(n):
    tot = 1
    for j in range(n):
        if i == j:
            continue
        tot = tot * x[j] * get_inverse(x[j] - x[i]) % p
    inputmask = (inputmask + y[i] * tot) % p

print(inputmask)
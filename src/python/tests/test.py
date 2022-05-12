from ..utils import get_inverse, n, t, p

x = []
y = []

for i in range(n):
    file = f"PreProcessing-Data/4-MSp-255/Randoms-MSp-P{i}"
    with open(file, "r") as f:
        x.append(i + 1)
        for line in f.readlines():
            y.append(int(line))
            break

for nn in range(t + 1, n + 1):
    inputmask = 0
    for i in range(nn):
        tot = 1
        for j in range(nn):
            if i == j:
                continue
            tot = tot * x[j] * get_inverse(x[j] - x[i]) % p
        inputmask = (inputmask + y[i] * tot) % p
    print(inputmask)

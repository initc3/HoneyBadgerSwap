import re
import sys
import time

import matplotlib.pyplot as plt
import numpy as np


# Set the default text font size
plt.rc('font', size=15)
# Set the axes title font size
plt.rc('axes', titlesize=15)
# Set the axes labels font size
plt.rc('axes', labelsize=15)
# Set the font size for x tick labels
plt.rc('xtick', labelsize=15)
# Set the font size for y tick labels
plt.rc('ytick', labelsize=15)
# Set the legend font size
plt.rc('legend', fontsize=12)
# Set the font size of the figure title
plt.rc('figure', titlesize=20)


def sample():
    return np.random.choice(list(pdf.keys()), p=list(pdf.values())) / 10


def simulate():
    delay_dict = {}
    interval = 5
    mpc_time = 0
    with open(f'ratel/benchmark/src/swap/pool_data/{pool_name}.csv', 'r') as f:
        lines = f.readlines()
        for line in lines[1:]:
            element = re.split(',|\t|\n', line)
            timestamp = float(element[0])

            if timestamp < start_time:
                continue
            if timestamp > end_time:
                break

            if mpc_time > timestamp:
                delay = mpc_time - timestamp
                mpc_time += sample()
            else:
                delay = 0
                mpc_time = timestamp + sample()
            delay //= interval
            if delay not in delay_dict.keys():
                delay_dict[delay] = 0
            delay_dict[delay] += 1

    s = sum(delay_dict.values())
    # print('sum', s)

    m = max(delay_dict.keys())
    # print(f'm {m}')

    delay_dict = {k * interval: v / s for k, v in delay_dict.items()}

    tot = 0
    pos = 0
    for k in sorted(delay_dict.keys()):
        tot += delay_dict[k]
        if tot > 0.85:
            pos = k
            break
    # print(f'pos {pos}')

    tot = 0
    for k in delay_dict.keys():
        if k <= 25:
            tot += delay_dict[k]
    # print(f'tot {tot}')

    return m, pos, tot


if __name__ == '__main__':
    start_time = int(sys.argv[1])
    end_time = int(sys.argv[2])
    pool_name = sys.argv[3]

    with open(f'ratel/benchmark/src/swap/pdf.txt', 'r') as f:
        line = f.readlines()[0]
        pdf = eval(line)

    # rep = 100
    # m, pos, tot = 0, 0, 0
    # for i in range(rep):
    #     _m, _pos, _tot = simulate()
    #     m += _m
    #     pos += _pos
    #     tot += _tot
    # print(f'm {m / rep}')
    # print(f'pos {pos / rep}')
    # print(f'tot {tot / rep}')

    delay_dict = {}
    interval = 5
    mpc_time = 0
    with open(f'ratel/benchmark/src/swap/pool_data/{pool_name}.csv', 'r') as f:
        lines = f.readlines()
        for line in lines[1:]:
            element = re.split(',|\t|\n', line)
            timestamp = float(element[0])

            if timestamp < start_time:
                continue
            if timestamp > end_time:
                break

            if mpc_time > timestamp:
                delay = mpc_time - timestamp
                mpc_time += sample()
            else:
                delay = 0
                mpc_time = timestamp + sample()
            delay //= interval
            if delay not in delay_dict.keys():
                delay_dict[delay] = 0
            delay_dict[delay] += 1

    s = sum(delay_dict.values())
    print('sum', s)
    m = max(delay_dict.keys())
    print('max', m)
    delay_dict = {k * interval + interval / 2 : v / s for k, v in delay_dict.items()}

    fig = plt.figure()
    ax1 = fig.add_subplot(111)
    ax1.set_xticks(np.arange(0, 65, 20))
    ax1.set_xlim(0, 65)
    ax1.set_ylim(0, 0.6)
    ax1.bar(delay_dict.keys(), delay_dict.values(), width=interval, color='darkseagreen')
    ax1.set_xlabel('Wait time(s)')
    ax1.set_ylabel('Probability')
    plt.show()
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


if __name__ == '__main__':
    start_time = int(sys.argv[1])
    end_time = int(sys.argv[2])
    pool_name = sys.argv[3]

    with open(f'ratel/benchmark/src/swap/pdf.txt', 'r') as f:
        line = f.readlines()[0]
        pdf = eval(line)

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
    delay_dict = {k * interval : v / s for k, v in delay_dict.items()}

    plt.figure(figsize=(13, 4))
    plt.xlim(0, 200)
    plt.ylim(0, 0.6)
    plt.bar(delay_dict.keys(), delay_dict.values(), width=interval, color='salmon')
    plt.xlabel('Wait time(sec)')
    plt.ylabel('Probability')
    plt.show()
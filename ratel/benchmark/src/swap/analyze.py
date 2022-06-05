import re
import sys

import matplotlib.pyplot as plt
import numpy as np
from ratel.benchmark.src.trade_latency import idx_op, idx_time, op_lock_acquired, idx_seq


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

block_time = 15


def scan(data_dir, server_id):

    init_time = 0
    with open(f'{data_dir}/gas.csv', 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split('\t|\n', line)
            seq = element[0]
            time = float(element[1])
            if init_time == 0:
                init_time = time
            wait_time[seq] = -time
            tx_time[seq] = float(element[2])

    with open(f'{data_dir}/latency_{server_id}.csv', 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split('\t|\n', line)

            seq = element[idx_seq]
            op = element[idx_op]
            time = float(element[idx_time])

            if op == op_lock_acquired:
                wait_time[seq] += time


if __name__ == '__main__':
    players = int(sys.argv[1])
    data_dir = sys.argv[2]

    wait_time = {}
    tx_time = {}
    for server_id in range(players):
        scan(data_dir, server_id)

    wait_time = {k: v / players for k, v in wait_time.items()}

    x, y = [], []
    for k in wait_time.keys():
        x.append(tx_time[k])
        y.append(wait_time[k])

    print(x, y)

    fig = plt.figure()
    ax1 = fig.add_subplot(111)
    ax2 = ax1.twinx()

    ax1.scatter(x, y, s=2, color='fuchsia', label=f'wait time')
    ax1.set_xlim(0, 3600)
    ax1.set_xlabel('Tx submission time(s)')
    ax1.set_ylabel('Wait time(s)')

    blocks = {}
    for t in tx_time.values():
        time = t // block_time *\
               block_time
        if time not in blocks.keys():
            blocks[time] = 0
        blocks[time] += 1
    lists = sorted(blocks.items())
    x, y = zip(*lists)
    ax2.bar(x, y, width=block_time, color='cornflowerblue', label=f'tx density')
    ax2.set_ylabel('Tx density(/15s)')
    ax1.legend(loc='upper left')
    ax2.legend(loc='upper right')
    ax2.set_ylim(0, 20)
    ax2.set_yticks(np.arange(0, 20, step=2))

    ax1.set_zorder(ax2.get_zorder() + 1)
    ax1.set_frame_on(False)

    plt.show()

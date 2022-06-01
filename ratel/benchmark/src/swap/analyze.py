import re
import sys

import matplotlib.pyplot as plt
from ratel.benchmark.src.trade_latency import idx_op, idx_time, op_lock_acquired, idx_seq


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
    ax1.set_xlim(0, 4000)
    ax1.set_xlabel('Tx submission time(sec)')
    ax1.set_ylabel('Wait time(sec)')

    blocks = {}
    for t in tx_time.values():
        time = t // block_time * block_time
        if time not in blocks.keys():
            blocks[time] = 0
        blocks[time] += 1
    lists = sorted(blocks.items())
    x, y = zip(*lists)
    ax2.bar(x, y, width=block_time, color='cornflowerblue', label=f'tx density')
    ax2.set_ylabel('Tx number')
    ax1.legend(loc='upper left')
    ax2.legend(loc='upper right')

    ax1.set_zorder(ax2.get_zorder() + 1)
    ax1.set_frame_on(False)

    plt.show()

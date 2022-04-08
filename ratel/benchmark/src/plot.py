import matplotlib.pyplot as plt
import numpy as np
import re
import sys

from matplotlib.ticker import MultipleLocator
from ratel.benchmark.src.calc import idx_op, idx_time, op_start_mpc_chain, op_start_mpc, op_end_mpc, \
    op_end_mpc_chain, op_lock_acquired


interval = 60
width = 0.1


def add(map, key, num=1):
    if key not in map.keys():
        map[key] = num
    else:
        map[key] += num


def deal(x):
    x = np.array(x)
    x -= min(x)
    x //= interval

    y = {}
    for v in x:
        add(y, v)

    return y


def scan():
    send_request = []
    start_mpc_chain = []
    lock_acquired = []
    start_mpc = []
    end_mpc = []
    end_mpc_chain = []

    serverID = 0

    file = f'{dir}/latency_{serverID}.csv'
    with open(file, 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split('\t|\n', line)

            op = element[idx_op]
            time = float(element[idx_time])

            if op == op_start_mpc_chain:
                start_mpc_chain.append(time)
            elif op == op_lock_acquired:
                lock_acquired.append(time)
            elif op == op_start_mpc:
                start_mpc.append(time)
            elif op == op_end_mpc:
                end_mpc.append(time)
            elif op == op_end_mpc_chain:
                end_mpc_chain.append(time)

    file = f'{dir}/gas.csv'
    with open(file, 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split('\t|\n', line)
            time = float(element[10])
            send_request.append(time)

    send_request = deal(send_request)
    start_mpc_chain = deal(start_mpc_chain)
    lock_acquired = deal(lock_acquired)
    start_mpc = deal(start_mpc)
    end_mpc = deal(end_mpc)
    end_mpc_chain = deal(end_mpc_chain)

    return send_request, start_mpc_chain, lock_acquired, start_mpc, end_mpc, end_mpc_chain


def plot(plt, map, offset, label):
    lists = sorted(map.items())
    x, y = zip(*lists)
    x, y = np.array(x), np.array(y)
    plt.bar(x + width * offset, y, width=width, label=label)


def draw(send_request, start_mpc_chain, lock_acquired, start_mpc, end_mpc, end_mpc_chain):
    plt.figure(figsize=(13, 4))

    cnt = 0
    plot(plt, send_request, cnt, 'send_request')

    cnt += 1
    plot(plt, start_mpc_chain, cnt, 'start_mpc_chain')

    cnt += 1
    plot(plt, lock_acquired, cnt, 'lock_acquired')

    cnt += 1
    plot(plt, start_mpc, cnt, 'start_mpc')

    cnt += 1
    plot(plt, end_mpc, cnt, 'end_mpc')

    cnt += 1
    plot(plt, end_mpc_chain, cnt, 'end_mpc_chain')

    plt.xlabel(f"time(/{interval}secs)")
    plt.ylabel("count")
    ax = plt.gca()
    ax.xaxis.set_major_locator(MultipleLocator(1))
    ax.yaxis.set_major_locator(MultipleLocator(10))
    plt.legend()
    plt.savefig(f"{dir}/fig.pdf")


if __name__ == '__main__':
    dir = sys.argv[1]

    send_request, start_mpc_chain, lock_acquired, start_mpc, end_mpc, end_mpc_chain = scan()
    draw(send_request, start_mpc_chain, lock_acquired, start_mpc, end_mpc, end_mpc_chain)

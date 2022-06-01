import re
import sys

import matplotlib.pyplot as plt
from ratel.benchmark.src.trade_latency import idx_op, idx_time, idx_seq, op_start_mpc, op_end_mpc


def scan(data_dir, server_id):
    with open(f'{data_dir}/latency_{server_id}.csv', 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split('\t|\n', line)

            seq = element[idx_seq]
            op = element[idx_op]
            time = float(element[idx_time])

            if op == op_start_mpc:
                mpc_time[seq] = -time
            elif op == op_end_mpc:
                mpc_time[seq] += time


if __name__ == '__main__':
    players = int(sys.argv[1])
    data_dir = sys.argv[2]

    mpc_time = {}
    for server_id in range(players):
        scan(data_dir, server_id)

    pdf = {}
    for v in mpc_time.values():
        v = int(v // 0.1)
        if v not in pdf.keys():
            pdf[v] = 0
        pdf[v] += 1
    s = sum(pdf.values())
    pdf = {k: v / s for k, v in pdf.items()}
    print(pdf)
    with open(f'ratel/benchmark/src/swap/pdf.txt', 'w') as f:
        f.write(str(pdf))

    plt.figure(figsize=(13, 4))
    plt.bar(pdf.keys(), pdf.values())
    plt.show()
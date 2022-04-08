import re
import sys

idx_op = 4
idx_time = 6

op_start_mpc_chain = '1' ### lock acquired
op_lock_acquired = '2'
op_start_mpc = '3'
op_end_mpc = '4'
op_end_mpc_chain = '6'

def scan(serverID):
    mpc = 0
    cnt_mpc = 0
    mpc_chain = 0
    cnt_mpc_chain = 0

    file = f'{dir}/latency_{serverID}.csv'
    with open(file, 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split('\t|\n', line)

            op = element[idx_op]
            time = float(element[idx_time])

            if op == op_lock_acquired:
                mpc_chain -= time
                cnt_mpc_chain += 1
            elif op == op_start_mpc:
                mpc -= time
                cnt_mpc += 1
            elif op == op_end_mpc:
                mpc += time
            elif op == op_end_mpc_chain:
                mpc_chain += time

    avg_mpc = mpc / cnt_mpc
    avg_mpc_chain = mpc_chain / cnt_mpc_chain
    return avg_mpc, avg_mpc_chain


if __name__ == '__main__':
    players = int(sys.argv[1])
    dir = sys.argv[2]

    for serverID in range(players):
        print(scan(serverID))
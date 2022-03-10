import asyncio
import os
import sys

import matplotlib.pyplot as plt
import shutil
import time

from ratel.src.python.utils import mpc_port, prog, offline_prog, blsPrime

players = int(sys.argv[1])
threshold = 1
max_concurrency = 20
output_file = 'ratel/benchmark/data/mp-spdz.txt'

def set_up_share_files(concurrency):
    for server_id in range(players):
        for i in range(concurrency):
            port = mpc_port + i * 100
            shutil.copyfile(f'ratel/benchmark/data/sharefiles/Transactions-P{server_id}-{mpc_port}.data', f'Persistence/Transactions-P{server_id}-{port}.data')


async def run_online_ONLY(server_id, port):
    start_time = time.perf_counter()
    cmd = f'{prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {blsPrime} hbswapTrade1'
    proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
    stdout, stderr = await proc.communicate()
    print(f'[{cmd!r} exited with {proc.returncode}]')
    if stdout:
        print(f'[stdout]\n{stdout.decode()}')
    if stderr:
        print(f'[stderr]\n{stderr.decode()}')
    end_time = time.perf_counter()
    duration = end_time - start_time
    return server_id, port, start_time, end_time, duration


async def run_online(server_id, port):
    start_time = time.perf_counter()
    cmd = f'{prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {blsPrime} -F hbswapTrade1'
    proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
    stdout, stderr = await proc.communicate()
    print(f'[{cmd!r} exited with {proc.returncode}]')
    if stdout:
        print(f'[stdout]\n{stdout.decode()}')
    if stderr:
        print(f'[stderr]\n{stderr.decode()}')
    end_time = time.perf_counter()
    duration = end_time - start_time
    return server_id, port, start_time, end_time, duration


async def run_offline(server_id, port):
    # cmd = f'{offline_prog} -p {server_id} -P {blsPrime} hbswapTrade1'
    cmd = f'{offline_prog} -N {players} -T {threshold} -p {server_id} -pn {port} hbswapTrade1'
    proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
    stdout, stderr = await proc.communicate()
    print(f'[{cmd!r} exited with {proc.returncode}]')
    if stdout:
        print(f'[stdout]\n{stdout.decode()}')
    if stderr:
        print(f'[stderr]\n{stderr.decode()}')


async def run_offline_phase():
    tasks = []
    for server_id in range(players):
        tasks.append(run_offline(server_id, mpc_port))
    await asyncio.gather(*tasks)


async def run_test(concurrency):
    set_up_share_files(concurrency)
    tasks = []
    for server_id in range(players):
        for i in range(concurrency):
            port = mpc_port + i * 100
            tasks.append(run_online_ONLY(server_id, port))
    results = await asyncio.gather(*tasks)
    for result in results:
        print(result)
    return results


def write_to(st):
    with open(output_file, 'a') as f:
        f.write(st)


async def main():
    if os.path.exists(output_file):
            os.remove(output_file)
    else:
        print("Can not delete the file as it doesn't exists")

    set_up_share_files(max_concurrency)

    x, y = [], []
    for concurrency in range(1, 1 + max_concurrency):
        results = await run_test(concurrency)
        sum = 0
        for server_id, port, start_time, end_time, duration in results:
            sum += duration
            write_to(f'{server_id}\t{port}\t{start_time}\t{end_time}\t{duration}\n')
        sum /= players * concurrency
        write_to(f'\n'
                 f'----------\n'
                 f'{concurrency}\t{sum}\n'
                 f'----------\n\n')

        x.append(concurrency)
        y.append(sum)

    write_to(f'{x}\n{y}\n')

    plt.figure(figsize=(13, 4))
    plt.plot(x, y)
    plt.savefig(f'ratel/benchmark/data/mp-spdz.pdf')


if __name__ == '__main__':
    asyncio.run(run_offline_phase())
    # asyncio.run(run_test(1))
    # asyncio.run(main())
    # loop = asyncio.get_event_loop()
    # loop.run_until_complete(main(loop))
    # loop.close()



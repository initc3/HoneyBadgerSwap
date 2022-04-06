import asyncio
import matplotlib.pyplot as plt
import shutil
import sys
import time

from ratel.src.python.utils import mpc_port, prog, offline_prog, blsPrime, repeat_experiment


def set_up_share_files(concurrency):
    for i in range(concurrency):
        port = mpc_port + i * 100
        for server_id in range(players):
            shutil.copyfile(f'ratel/benchmark/data/sharefiles/Transactions-P{server_id}-{mpc_port}.data', f'Persistence/Transactions-P{server_id}-{port}.data')


async def execute(cmd):
    proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)
    stdout, stderr = await proc.communicate()
    print(f'[{cmd!r} exited with {proc.returncode}]')
    if stdout:
        print(f'[stdout]\n{stdout.decode()}')
    if stderr:
        print(f'[stderr]\n{stderr.decode()}')


async def run_online_ONLY(server_id, port):
    cmd = f'{prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {blsPrime} hbswapTrade1'
    await execute(cmd)


async def run_online(server_id, port):
    dst_dir = f'Player-data-port-{port}-copy'
    cmd = f'rm -rf {dst_dir}'
    await execute(cmd)

    src_dir = f'Player-data-port-{port}'
    cmd = f'cp -rf {src_dir} {dst_dir}'
    await execute(cmd)

    dir = dst_dir
    cmd = f'{prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {blsPrime} -F --prep-dir {dir} hbswapTrade1'
    await execute(cmd)


async def run_offline(server_id, port):
    dir = f'Player-data-port-{port}'
    cmd = f'{offline_prog} -N {players} -T {threshold} -p {server_id} -pn {port} -P {blsPrime} --prep-dir {dir} hbswapTrade1'
    await execute(cmd)


async def test(func, server_id, port):
    start_time = time.perf_counter()

    await eval(func)(server_id, port)

    end_time = time.perf_counter()
    duration = end_time - start_time

    return duration


async def run_test(func, concurrency):
    tasks = []
    for i in range(concurrency):
        port = mpc_port + i * 100
        for server_id in range(players):
            tasks.append(test(func, server_id, port))
    results = await asyncio.gather(*tasks)
    print(f'!!!! {func} {results}')
    return sum(results) / (players * concurrency)


async def rep(func, concurrency):
    sum = 0
    for i in range(repeat_experiment):
        sum += await run_test(func, concurrency)
    avg = sum / repeat_experiment
    return avg


async def main():
    x, y_offline, y_online, y_online_ONLY = [], [], [], []
    for concurrency in range(1, 1 + max_concurrency):
        x.append(concurrency)
        y_offline.append(await rep('run_offline', concurrency))
        y_online.append(await rep('run_online', concurrency))
        y_online_ONLY.append(await rep('run_online_ONLY', concurrency))

    with open('ratel/benchmark/data/mp-spdz.txt', 'w') as f:
        f.write(str(x) + '\n')
        f.write(str(y_offline) + '\n')
        f.write(str(y_online) + '\n')
        f.write(str(y_online_ONLY) + '\n')

    plt.figure(figsize=(13, 4))
    plt.scatter(x, y_offline)
    plt.scatter(x, y_online)
    plt.scatter(x, y_online_ONLY)
    plt.savefig(f'ratel/benchmark/data/mp-spdz.pdf')


if __name__ == '__main__':
    players = int(sys.argv[1])
    threshold = int(sys.argv[2])
    max_concurrency = int(sys.argv[3])

    set_up_share_files(max_concurrency)
    asyncio.run(main())




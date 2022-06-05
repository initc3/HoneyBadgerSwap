import asyncio
import os
import sys

from ratel.benchmark.src.test_mpc import run_test

if __name__ == '__main__':
    players = int(sys.argv[1])
    threshold = int(sys.argv[2])
    concurrency = int(sys.argv[3])

    directory = os.fsencode(f'ratel/genfiles/mpc')

    for file in os.listdir(directory):
        filename = os.fsdecode(file)
        if filename.endswith(".mpc"):
            asyncio.run(run_test('run_offline', players, threshold, concurrency, filename[:-4]))

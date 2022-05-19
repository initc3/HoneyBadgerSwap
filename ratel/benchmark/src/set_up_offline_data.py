import asyncio
import sys

from ratel.benchmark.src.test_mpc import run_test

if __name__ == '__main__':
    players = int(sys.argv[1])
    threshold = int(sys.argv[2])
    concurrency = int(sys.argv[3])
    asyncio.run(run_test('run_offline', players, threshold, concurrency, 'hbswapTrade1'))
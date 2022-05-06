import asyncio
import sys
import time

from web3 import Web3

from ratel.src.python.Server import Server
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract, repeat_experiment


async def rep(shareBatchSize):
    sum = 0
    for i in range(repeat_experiment):
        start_time = time.perf_counter()
        await server.genInputMask(shareBatchSize)
        end_time = time.perf_counter()
        duration = end_time - start_time

        with open(f'ratel/benchmark/data/inputmask_generation_latency_{server.serverID}.csv', 'a') as f:
            f.write(f'shareBatchSize\t{shareBatchSize}\t'
                    f'duration\t{duration}\n')
        print('**** duration', duration)

        sum += duration

    avg = sum / repeat_experiment
    return avg


async def test(shareBatchSize):
    print('**** shareBatchSize', shareBatchSize)
    avg_duration = await rep(shareBatchSize)
    with open(f'ratel/benchmark/data/inputmask_generation_latency_{server.serverID}.csv', 'a') as f:
        f.write(f'shareBatchSize\t{shareBatchSize}\t'
                f'avg_duration\t{avg_duration}\n')


if __name__ == '__main__':
    serverID = int(sys.argv[1])
    players = int(sys.argv[2])
    threshold = int(sys.argv[3])

    web3 = Web3(Web3.WebsocketProvider(url))
    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)
    concurrency = 1

    server = Server(
        serverID,
        web3,
        appContract,
        players,
        threshold,
        concurrency,
        None,
        0
    )

    batch = 10000
    for shareBatchSize in range(batch, 5 * batch, batch):
        asyncio.run(test(shareBatchSize))

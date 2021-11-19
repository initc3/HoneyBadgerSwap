import sys
import time

from web3 import Web3

from ratel.src.python.Server import Server
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract

def test(shareBatchSize):
    print('****', shareBatchSize)
    start_time = time.perf_counter()
    server.genInputMask(shareBatchSize)
    end_time = time.perf_counter()

    with open(f'ratel/benchmark/data/inputmask_generation_latency_{server.serverID}.csv', 'a') as f:
        f.write(f'shareBatchSize\t{shareBatchSize}\t'
                f'start_time\t{start_time}\t'
                f'end_time\t{end_time}\t'
                f'time_dif\t{end_time - start_time}\n')

if __name__ == '__main__':
    serverID = int(sys.argv[1])

    web3 = Web3(Web3.WebsocketProvider(url))
    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)
    concurrency = 1

    server = Server(
        serverID,
        web3,
        appContract,
        None,
        None,
        concurrency,
    )

    batch = 10000
    for shareBatchSize in range(batch, 10 * batch, batch):
        test(shareBatchSize)

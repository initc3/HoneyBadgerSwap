import asyncio
import sys

from web3 import Web3

from ratel.genfiles.python import rockPaperScissors
from ratel.genfiles.python.rockPaperScissorsRecover import recover
from ratel.src.python.Server import Server
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract

contract_name = 'rockPaperScissors'

if __name__ == '__main__':
    serverID = int(sys.argv[1])
    init_players = int(sys.argv[2])
    init_threshold = int(sys.argv[3])
    concurrency = int(sys.argv[4])
    test = bool(sys.argv[5])

    web3 = Web3(Web3.WebsocketProvider(url))

    ### App contract
    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)
    ###

    server = Server(
        serverID,
        web3,
        appContract,
        init_players,
        init_threshold,
        concurrency,
        recover,
        test,
    )

    server.loop.run_until_complete(server.init(rockPaperScissors.monitor(server)))
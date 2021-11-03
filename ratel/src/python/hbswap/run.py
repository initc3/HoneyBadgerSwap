import asyncio
import ratel.genfiles.python.hbswap as hbswap
import sys

### recover function
from ratel.genfiles.python.hbswapRecover import recover
###
from ratel.src.python.Server import Server
from ratel.src.python.deploy import parse_contract, url, addrs
from web3 import Web3

if __name__ == '__main__':
    serverID = int(sys.argv[1])
    init_players = int(sys.argv[2])
    init_threshold = int(sys.argv[3])

    web3 = Web3(Web3.WebsocketProvider(url))

    ### App contract
    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=addrs[0], abi=abi)
    ###

    server = Server(
        serverID,
        web3,
        appContract,
        init_players,
        init_threshold,
    )

    loop = asyncio.get_event_loop()
    loop.run_until_complete(server.init(recover, hbswap.monitor(server, loop)))
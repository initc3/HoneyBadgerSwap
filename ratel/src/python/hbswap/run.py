import asyncio
import ratel.genfiles.python.hbswap as hbswap
import sys

### recover function
from ratel.genfiles.python.hbswapRecover import recover
###
from ratel.src.python.Server import Server, getAccount
from ratel.src.python.deploy import parse_contract, appAddress, url
from ratel.src.python.utils import openDB, location_db, http_port, http_host, confirmation
from web3 import Web3

if __name__ == '__main__':
    serverID = int(sys.argv[1])
    init_players = int(sys.argv[2])
    init_threshold = int(sys.argv[3])

    db = openDB(location_db(serverID))
    web3 = Web3(Web3.WebsocketProvider(url))
    account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')

    ### App contract
    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)
    ###

    server = Server(
        serverID,
        db,
        http_host,
        http_port + serverID,
        appContract,
        web3,
        account,
        confirmation,
        init_players,
        init_threshold,
    )

    loop = asyncio.get_event_loop()
    loop.run_until_complete(server.init(recover, hbswap.monitor(server, loop)))
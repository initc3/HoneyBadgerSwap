import asyncio
import ratel.genfiles.python.hbswap as hbswap
import sys

from ratel.genfiles.python.hbswapRecover import recover
from ratel.src.python.Server import Server, getAccount
from ratel.src.python.deploy import parse_contract, appAddress, confirmation, url
from ratel.src.python.utils import openDB, location_db, http_port, http_host
from web3 import Web3

if __name__ == '__main__':
    serverID = int(sys.argv[1])
    db = openDB(location_db(serverID))
    web3 = Web3(Web3.WebsocketProvider(url))
    account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')

    ### ACCESS contract
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
        confirmation
    )

    ### EDIT application specific monitoring tasks below
    app_tasks = [
        hbswap.monitorSecretDeposit(web3, db, serverID, appContract, confirmation, account),
        hbswap.monitorSecretWithdraw(web3, db, serverID, appContract, confirmation, account),
        hbswap.monitorInitPool(web3, db, serverID, appContract, confirmation, account),
        hbswap.monitorAddLiquidity(web3, db, serverID, appContract, confirmation, account),
        hbswap.monitorRemoveLiquidity(web3, db, serverID, appContract, confirmation, account),
        hbswap.monitorTrade(web3, db, serverID, appContract, confirmation, account)
    ]
    ###

    ### CHANGE: recover function
    asyncio.run(server.init(recover, app_tasks))
    ###
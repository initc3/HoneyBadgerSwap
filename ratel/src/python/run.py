import asyncio
import sys

from web3 import Web3

from ratel.src.python.Server import Server
from ratel.src.python.deploy import parse_contract, getAccount, preprocessing, appAddress, confirmation, url
from ratel.src.python.utils import openDB, location_db, http_port, http_host
from ratel.genfiles.python.test import monitorSecretDeposit, monitorInitPool, monitorAddLiquidity, monitorTrade, \
    monitorRemoveLiquidity, monitorSecretWithdraw

if __name__ == '__main__':
    serverID = sys.argv[1]

    web3 = Web3(Web3.WebsocketProvider(url))

    account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')

    db = openDB(location_db(serverID))

    abi, bytecode = parse_contract('Test')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    server = Server(
        serverID, db, http_host, http_port + int(serverID)
    )

    loop = asyncio.get_event_loop()
    tasks = [monitorSecretDeposit(web3, db, serverID, appContract, confirmation, account),
             monitorSecretWithdraw(web3, db, serverID, appContract, confirmation, account),
             monitorInitPool(web3, db, serverID, appContract, confirmation, account),
             monitorAddLiquidity(web3, db, serverID, appContract, confirmation, account),
             monitorRemoveLiquidity(web3, db, serverID, appContract, confirmation, account),
             monitorTrade(web3, db, serverID, appContract, confirmation, account),
             preprocessing(db, appContract, serverID),
             server.http_server()]
    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()

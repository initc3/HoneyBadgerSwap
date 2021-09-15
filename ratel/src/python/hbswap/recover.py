import asyncio
import sys

from web3 import Web3

from ratel.genfiles.python.hbswapRecover import recover
from ratel.src.python.Server import Server, getAccount
from ratel.src.python.deploy import parse_contract, appAddress, confirmation, url
from ratel.src.python.utils import openDB, location_db, http_port, http_host
from ratel.genfiles.python.hbswap import monitorSecretDeposit, monitorInitPool, monitorAddLiquidity, monitorTrade, \
    monitorRemoveLiquidity, monitorSecretWithdraw

async def init():
    isServer = appContract.functions.isServer(account.address).call()
    print('!!!! isServer', isServer)
    if not isServer:
        server.registerServer()
        await server.recoverHistory(recover)

    tasks = [monitorSecretDeposit(web3, db, serverID, appContract, confirmation, account),
             monitorSecretWithdraw(web3, db, serverID, appContract, confirmation, account),
             monitorInitPool(web3, db, serverID, appContract, confirmation, account),
             monitorAddLiquidity(web3, db, serverID, appContract, confirmation, account),
             monitorRemoveLiquidity(web3, db, serverID, appContract, confirmation, account),
             monitorTrade(web3, db, serverID, appContract, confirmation, account),
             server.preprocessing(),
             server.monitorNewServer(),
             server.http_server()]
    await asyncio.gather(*tasks)

if __name__ == '__main__':
    serverID = int(sys.argv[1])

    db = openDB(location_db(serverID))

    web3 = Web3(Web3.WebsocketProvider(url))

    account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    server = Server(
        serverID, db, http_host, http_port + serverID, appContract, web3, account, confirmation
    )

    loop = asyncio.get_event_loop()
    tasks = [init(),
             server.monitorGenInputMask(),
            ]
    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()
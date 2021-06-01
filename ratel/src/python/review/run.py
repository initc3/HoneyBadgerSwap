import asyncio
import sys

from web3 import Web3

from ratel.genfiles.python.review import monitorPeerReview, monitorCalcResult
from ratel.src.python.Server import Server
from ratel.src.python.deploy import parse_contract, getAccount, preprocessing, appAddress, confirmation, url
from ratel.src.python.utils import openDB, location_db, http_port, http_host

contract_name = 'review'

if __name__ == '__main__':
    serverID = sys.argv[1]

    web3 = Web3(Web3.WebsocketProvider(url))

    account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')

    db = openDB(location_db(serverID))

    abi, bytecode = parse_contract(contract_name)
    contract = web3.eth.contract(address=appAddress, abi=abi)

    server = Server(
        serverID, db, http_host, http_port + int(serverID)
    )

    loop = asyncio.get_event_loop()
    tasks = [preprocessing(db, contract, serverID),
             server.http_server(),
             monitorPeerReview(web3, db, serverID, contract, confirmation, account),
             monitorCalcResult(web3, db, serverID, contract, confirmation, account),
            ]
    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()

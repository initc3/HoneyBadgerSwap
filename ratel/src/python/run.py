import json
import sys

from web3 import Web3

from ratel.src.python.utils import openDB, location_db
from ratel.genfiles.python.test import monitorSecretDeposit

if __name__ == "__main__":
    server_id = sys.argv[1]

    url = 'ws://0.0.0.0:8546'
    web3 = Web3(Web3.WebsocketProvider(url))

    db = openDB(location_db(server_id))

    address = '0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2'
    f = open('ratel/genfiles/build/contracts/Test.json')
    data = json.load(f)
    abi = data['abi']
    myContract = web3.eth.contract(address=address, abi=abi)
    confirmation = 1

    monitorSecretDeposit(web3, db, server_id, myContract, confirmation)


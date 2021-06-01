import asyncio
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks
from ratel.src.python.deploy import url, parse_contract, appAddress, tokenAddress, ETH, reserveInput
from ratel.src.python.utils import fp, blsPrime

contract_name = 'rockPaperScissors'

def createGame(appContract, value):
    idx = reserveInput(web3, appContract, 1)[0]
    mask = asyncio.run(get_inputmasks(f'{idx}'))[0]
    maskedValue = value + mask
    tx_hash = appContract.functions.createGame(idx, maskedValue).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.CreateGame().processReceipt(receipt)
    print(log)
    gameId = log[0]['args']['gameId']
    while True:
        time.sleep(1)
        status = appContract.functions.status(gameId).call()
        if status == 1:
            return gameId

def joinGame(appContract, gameId, value):
    idx = reserveInput(web3, appContract, 1)[0]
    print(idx)
    mask = asyncio.run(get_inputmasks(f'{idx}'))[0]
    maskedValue = value + mask
    tx_hash = appContract.functions.joinGame(gameId, idx, maskedValue).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    while True:
        time.sleep(1)
        status = appContract.functions.status(gameId).call()
        if status == 2:
            return

def startRecon(appContract, gameId):
    tx_hash = appContract.functions.startRecon(gameId).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    gameId = createGame(appContract, 1)
    joinGame(appContract, gameId, 1)
    startRecon(appContract, gameId)
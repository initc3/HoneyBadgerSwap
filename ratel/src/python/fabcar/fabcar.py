import asyncio
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks
from ratel.src.python.deploy import url, parse_contract, appAddress, reserveInput, getAccount

contract_name = 'fabcar'

def createTruck(appContract):
    tx_hash = appContract.functions.createTruck().transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.NewTruck().processReceipt(receipt)
    print(log)
    truckId = log[0]['args']['truckId']
    return truckId

def recordShipment(appContract, truckId, timeLoad, timeUnload, account):
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = timeLoad + maskA, timeUnload + maskB
    tx_hash = appContract.functions.recordShipment(truckId, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

def queryPositions(appContract, truckId, tL, tR, account):
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = tL + maskA, tR + maskB
    tx_hash = appContract.functions.queryPositions(truckId, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

def queryNumber(appContract, truckId, tL, tR, account):
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = tL + maskA, tR + maskB
    tx_hash = appContract.functions.queryNumber(truckId, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

def queryFirst(appContract, truckId, tL, tR, account):
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = tL + maskA, tR + maskB
    tx_hash = appContract.functions.queryFirst(truckId, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/server_0/')

    truckId = createTruck(appContract)

    recordShipment(appContract, truckId, 1, 3, account)
    recordShipment(appContract, truckId, 2, 4, account)
    recordShipment(appContract, truckId, 3, 5, account)

    time.sleep(10)

    # truckId = 1

    queryPositions(appContract, truckId, 4, 4, account)
    queryNumber(appContract, truckId, 4, 4, account)
    queryFirst(appContract, truckId, 4, 4, account)


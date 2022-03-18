import asyncio
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract, getAccount, players, blsPrime, sign_and_send

contract_name = 'colAuction'


def initAuction(appContract,account):
    web3.eth.defaultAccount = account.address
    tx = appContract.functions.initAuction().buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.InitAuction().processReceipt(receipt)
    colAuctionId = log[0]['args']['colAuctionId']
    while True:
        time.sleep(0.1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 1:
            return colAuctionId

# means I'll buy up to Amt if the prices reaches $X or below
def inputAuction(appContract,colAuctionId,X,Amt,account):
    idx = reserveInput(web3, appContract, 1, account)[0]
    mask = asyncio.run(get_inputmasks(players(appContract), f'{idx}'))[0]
    maskedX = (X + mask) % blsPrime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.inputAuction(colAuctionId, idx, maskedX, Amt).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    web3.eth.wait_for_transaction_receipt(tx_hash)

    while True:
        time.sleep(0.1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 2:
            return

def dutchAuctionSettle(appContract, colAuctionId, AmtToSell, StartPrice, LowestPrice, account):
    idxA, idxB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxA},{idxB}'))
    maskedA, maskedB = (AmtToSell + maskA)%blsPrime, (StartPrice + maskB)%blsPrime
    
    web3.eth.defaultAccount = account.address
    tx = appContract.functions.dutchAuctionSettle(colAuctionId,idxA,maskedA,idxB,maskedB,LowestPrice).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    web3.eth.wait_for_transaction_receipt(tx_hash)

    while True:
        res = appContract.functions.colres(colAuctionId).call()
        if res == 3:
            print(res)
            break
        time.sleep(0.1)


if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)



    client_1 = getAccount(web3,f'/opt/poa/keystore/client_1/')
    client_2 = getAccount(web3,f'/opt/poa/keystore/client_2/')
    client_3 = getAccount(web3,f'/opt/poa/keystore/client_3/')
    client_4 = getAccount(web3,f'/opt/poa/keystore/client_4/')
    client_5 = getAccount(web3,f'/opt/poa/keystore/client_5/')
    

    colAuctionId1 = initAuction(appContract,client_1)
    print(colAuctionId1)

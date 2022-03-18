import asyncio
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract, getAccount, players, blsPrime, sign_and_send

contract_name = 'colAuction'


# def toyGame(appContract,val1,account):
#     idx = reserveInput(web3, appContract, 1, account)[0]
#     mask = asyncio.run(get_inputmasks(players(appContract), f'{idx}'))[0]
#     maskedVal1 = (val1 + mask) % blsPrime

#     web3.eth.defaultAccount = account.address
#     tx = appContract.functions.toyGame(idx,maskedVal1).buildTransaction({
#         'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
#     })
#     tx_hash = sign_and_send(tx, web3, account)
#     receipt = web3.eth.get_transaction_receipt(tx_hash)
#     log = appContract.events.ToyGame().processReceipt(receipt)
#     toyId = log[0]['args']['toyId']
#     while True:
#         time.sleep(1)
#         status = appContract.functions.status(toyId).call()
#         if status == 1:
#             return toyId


def kick(appContract,tab,lot,account):
    idx = reserveInput(web3, appContract, 1, account)[0]
    mask = asyncio.run(get_inputmasks(players(appContract), f'{idx}'))[0]
    maskedlot = (lot + mask) % blsPrime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.kick(tab, idx, maskedlot).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.Kick().processReceipt(receipt)
    colAuctionId = log[0]['args']['colAuctionId']
    while True:
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 1:
            return colAuctionId

def tend(appContract, colAuctionId, lot, bid, account):
    idxlot, idxbid = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxlot},{idxbid}'))
    maskedlot = (lot + maskA) % blsPrime
    maskedbid = (bid + maskB) % blsPrime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.tend(colAuctionId, idxlot, maskedlot, idxbid, maskedbid).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    web3.eth.wait_for_transaction_receipt(tx_hash)

    while True:
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 2:
            return


if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)



    client_1 = getAccount(web3,f'/opt/poa/keystore/client_1/')
    client_2 = getAccount(web3,f'/opt/poa/keystore/client_2/')
    client_3 = getAccount(web3,f'/opt/poa/keystore/client_3/')
    
    # colId = toyGame(appContract,10,client_1)
    # print(colId)

    # colId = toyGame(appContract,1,client_1)
    # print(colId)


    # numAuct = 5
    # AuctAddrs = []
    # AuctAcc = []
    # for aucMemID in range(numAuct):
    #     account = getAccount(web3, f'/opt/poa/keystore/client_{aucMemID+1}/')
    #     AuctAcc.append(account)
    #     AuctAddrs.append(account.address)

    # usr: address to receive residual collateral after the auction
    # gal: address to receive raised DAI
    # bid: amount of DAI a bidder would like to pay
    # function kick(uint tab, uint lot, address usr, address gal, uint bid) public {
    tab1_0 = 100 # tab: amount of DAI to raise; 
    lot1_0 = 50 # lot: amount of collateral for sell
    colAuctionId = kick(appContract, tab1_0,lot1_0, client_3)
    print(colAuctionId)

    bid1_1 = 10
    lot1_1 = lot1_0
    tend(appContract,lot1_1,bid1_1,client_1)




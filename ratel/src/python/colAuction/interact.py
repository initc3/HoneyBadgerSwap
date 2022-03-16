import asyncio
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract, getAccount, players, blsPrime, sign_and_send

contract_name = 'colAuction'


# def toy(appContract,accout):
#     idx = reserveInput(web3, appContract, 1, account)[0]

#     web3.eth.defaultAccount = account.address
#     tx = appContract.functions.toy(idx).buildTransaction({
#         'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
#     })
#     tx_hash = sign_and_send(tx, web3, account)
#     receipt = web3.eth.get_transaction_receipt(tx_hash)
#     log = appContract.events.kick().processReceipt(receipt)
#     colAuctionId = log[0]['args']['colAuctionId']
#     while True:
#         time.sleep(1)
#         status = appContract.functions.status(colAuctionId).call()
#         if status == 1:
#             return colAuctionId


def kick(appContract,tab,lot,usr,gal,bid,account):
    idx = reserveInput(web3, appContract, 1, account)[0]
#    mask = asyncio.run(get_inputmasks(players(appContract), f'{idx}'))[0]
#    maskedValue = (value1 + mask) % blsPrime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.kick(idx, tab, lot, usr, gal, bid).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.kick().processReceipt(receipt)
    colAuctionId = log[0]['args']['colAuctionId']
    while True:
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 1:
            return colAuctionId

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)


    numAuct = 5
    AuctAddrs = []
    AuctAcc = []
    for aucMemID in range(numAuct):
        account = getAccount(web3, f'/opt/poa/keystore/aucMem_{aucMemID}/')
        AuctAcc.append(account)
        AuctAddrs.append(account.address)

    # colId = toy(appContract,)
    # print(colId)

       
    # usr: address to receive residual collateral after the auction
    # gal: address to receive raised DAI
    # bid: amount of DAI a bidder would like to pay
    # function kick(uint tab, uint lot, address usr, address gal, uint bid) public {
    tab1 = 100 # tab: amount of DAI to raise; 
    lot1 = 50 # lot: amount of collateral for sell
    usr1 = AuctAddrs[0].address # usr: address to receive residual collateral after the auction
    gal1 = AuctAddrs[1].address
    bid1 = 10
    colAuctionId = kick(appContract, tab1,lot1,usr1,gal1,bid1, AuctAcc[0])
import asyncio
from curses.ascii import SP
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import fp,parse_contract, getAccount, players, blsPrime, sign_and_send

contract_name = 'colAuction'


def createAuction(appContract,StartPrice,FloorPrice,totalAmt,account):
    idx1,idx2,idx3 = reserveInput(web3, appContract, 3, account)
    mask1,mask2,mask3 = asyncio.run(get_inputmasks(players(appContract), f'{idx1},{idx2},{idx3}'))
    maskedSP, maskedFP, maskedTM = (StartPrice + mask1) % blsPrime, (FloorPrice + mask2) % blsPrime, (totalAmt + mask3) % blsPrime
    
    web3.eth.defaultAccount = account.address
    tx = appContract.functions.createAuction(idx1,maskedSP,idx2,maskedFP,idx3,maskedTM).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.CreateAuction().processReceipt(receipt)
    colAuctionId = log[0]['args']['colAuctionId']
    while True:
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 1:
            return colAuctionId

# means I'll buy up to $amt if the prices reaches $price or below
def submitBids(appContract,colAuctionId,price,amt,account):
    status = appContract.functions.status(colAuctionId).call()
    if status == 3:
        return

    idx1, idx2 = reserveInput(web3, appContract, 2, account)
    mask1, mask2 = asyncio.run(get_inputmasks(players(appContract), f'{idx1},{idx2}'))
    maskedP, maskedAmt = (price + mask1) % blsPrime, (amt + mask2) % blsPrime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.submitBids(colAuctionId, idx1, maskedP, idx2, maskedAmt).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    web3.eth.wait_for_transaction_receipt(tx_hash)

    while True:
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status >= 2:
            return



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
    client_6 = getAccount(web3,f'/opt/poa/keystore/client_6/')
    client_7 = getAccount(web3,f'/opt/poa/keystore/client_7/')
    
    #test auction success

    print('==================================')

    
    totalAmt1 = 20
    StartPrice1 = 10
    FloorPrice1 = 1 

    colAuctionId1 = createAuction(appContract,StartPrice1,FloorPrice1,totalAmt1,client_1)
    print('new Auction id:',colAuctionId1)

    time.sleep(1)

    price2 = 5
    Amt2 = 10
    submitBids(appContract,colAuctionId1,price2,Amt2,client_2)
    print('finished input client_2')

    price3 = 3
    Amt3 = 6
    submitBids(appContract,colAuctionId1,price3,Amt3,client_3)
    print('finished input client_3')
    
    price4 = 7
    Amt4 = 7
    submitBids(appContract,colAuctionId1,price4,Amt4,client_4)
    print('finished input client_4')

    price5 = 2
    Amt5 = 9
    submitBids(appContract,colAuctionId1,price5,Amt5,client_5)
    print('finished input client_5')



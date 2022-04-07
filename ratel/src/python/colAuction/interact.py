import asyncio
from curses.ascii import SP
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import parse_contract, getAccount, players, blsPrime, sign_and_send

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
    log = appContract.events.InitAuction().processReceipt(receipt)
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

def scheduleCheck(appContract, colAuctionId, StartPrice, FloorPrice, totalAmt, account):
    curPrice = StartPrice

    while True:
        if curPrice < FloorPrice:
            print('auction failed!!!')
            return

        idx1 = reserveInput(web3, appContract, 1, account)[0]
        mask1 = asyncio.run(get_inputmasks(players(appContract), f'{idx1}'))[0]
        maskedCP = (curPrice + mask1) % blsPrime
    
        web3.eth.defaultAccount = account.address
        tx = appContract.functions.scheduleCheck(colAuctionId,idx1,maskedCP).buildTransaction({
            'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
        })
        tx_hash = sign_and_send(tx, web3, account)
        web3.eth.wait_for_transaction_receipt(tx_hash)

        for i in range(10):
            res = appContract.functions.colres(colAuctionId).call()
            if res != '':
                print(res)
                return
            time.sleep(1)
        
        curPrice = curPrice*0.99
        


def closeAuction(appContract, colAuctionId, account):
    # idx1, idx2, idx3= reserveInput(web3, appContract, 3, account)
    # mask1, mask2, mask3 = asyncio.run(get_inputmasks(players(appContract), f'{idx1},{idx2},{idx3}'))
    # maskedAmt, maskedSP, maskedLP = (AmtToSell + mask1) % blsPrime, (StartPrice + mask2) % blsPrime, (LowestPrice + mask3) % blsPrime
    
    web3.eth.defaultAccount = account.address
    tx = appContract.functions.closeAuction(colAuctionId).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    web3.eth.wait_for_transaction_receipt(tx_hash)

    while True:
        res = appContract.functions.colres(colAuctionId).call()
        if res != '':
            print(res)
            break
        time.sleep(1)



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
    scheduleCheck(appContract,StartPrice1,FloorPrice1,totalAmt1,client_1)
    print('new Auction id:',colAuctionId1)

    time.sleep(1)

    X2 = 5
    Amt2 = 10
    inputAuction(appContract,colAuctionId1,X2,Amt2,client_2)
    print('finished input client_2')

    X3 = 3
    Amt3 = 6
    inputAuction(appContract,colAuctionId1,X3,Amt3,client_3)
    print('finished input client_3')
    
    X4 = 7
    Amt4 = 7
    inputAuction(appContract,colAuctionId1,X4,Amt4,client_4)
    print('finished input client_4')

    X5 = 2
    Amt5 = 9
    inputAuction(appContract,colAuctionId1,X5,Amt5,client_5)
    print('finished input client_5')

    dutchAuctionSettle(appContract,colAuctionId1,AmtToSell1,StartPrice1,LowestPrice1,client_1)
    print('finished settle')

    #test auction failure(since the price is lower than Lowest Price)

    print('==================================')

    colAuctionId2 = initAuction(appContract,client_1)
    print('new Auction id:',colAuctionId2)

    X2 = 5
    Amt2 = 10
    inputAuction(appContract,colAuctionId2,X2,Amt2,client_2)
    print('finished input client_2')

    X3 = 3
    Amt3 = 6
    inputAuction(appContract,colAuctionId2,X3,Amt3,client_3)
    print('finished input client_3')
    
    X4 = 7
    Amt4 = 7
    inputAuction(appContract,colAuctionId2,X4,Amt4,client_4)
    print('finished input client_4')

    X5 = 2
    Amt5 = 9
    inputAuction(appContract,colAuctionId2,X5,Amt5,client_5)
    print('finished input client_5')

    X6 = 6
    Amt6 = 1
    inputAuction(appContract,colAuctionId2,X6,Amt6,client_6)
    print('finished input client_6')

    X7 = 2
    Amt7 = 7
    inputAuction(appContract,colAuctionId2,X7,Amt7,client_7)
    print('finished input client_7')

    AmtToSell2 = 10
    StartPrice2 = 10
    LowestPrice2 = 6
    dutchAuctionSettle(appContract,colAuctionId2,AmtToSell2,StartPrice2,LowestPrice2,client_1)
    print('finished settle')

    #auction failes (since AmtTosell > sum of all amt)

    print('==================================')

    colAuctionId3 = initAuction(appContract,client_1)
    print('new Auction id:',colAuctionId3)

    X2 = 5
    Amt2 = 10
    inputAuction(appContract,colAuctionId3,X2,Amt2,client_2)
    print('finished input client_2')

    X3 = 3
    Amt3 = 6
    inputAuction(appContract,colAuctionId3,X3,Amt3,client_3)
    print('finished input client_3')
    
    X4 = 7
    Amt4 = 7
    inputAuction(appContract,colAuctionId3,X4,Amt4,client_4)
    print('finished input client_4')

    X5 = 2
    Amt5 = 9
    inputAuction(appContract,colAuctionId3,X5,Amt5,client_5)
    print('finished input client_5')

    X6 = 2
    Amt6 = 5
    inputAuction(appContract,colAuctionId2,X6,Amt6,client_6)
    print('finished input client_6')



    AmtToSell3 = 40
    StartPrice3 = 10
    LowestPrice3 = 6
    dutchAuctionSettle(appContract,colAuctionId3,AmtToSell3,StartPrice3,LowestPrice3,client_1)
    print('finished settle')


import asyncio
from curses.ascii import SP
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr, token_addrs
from ratel.src.python.utils import fp, parse_contract, getAccount, players, prime, sign_and_send, threshold

contract_name = 'colAuction'

bids_cnt = []


def createAuction(appContract, StartPrice, FloorPrice, totalAmt, token, aucapp_addr, account):
    colAuctionlast = appContract.functions.colAuctionCnt().call()

    bids_cnt.append(0)

    #    idx1 = reserveInput(web3, appContract, 1, account)[0]
    #    mask1 = asyncio.run(get_inputmasks(players(appContract), f'{idx1}', threshold(appContract)))[0]
    #    maskedTM = (totalAmt + mask1) % prime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.createAuction(StartPrice, FloorPrice, totalAmt, token, aucapp_addr).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)

    while True:
        colAuctionId = appContract.functions.colAuctionCnt().call()
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status == 2 and colAuctionId != colAuctionlast:
            return colAuctionId


# means I'll buy up to $amt if the prices reaches $price or below
def submitBids(appContract, colAuctionId, price, amt, account):
    status = appContract.functions.status(colAuctionId).call()
    if status == 1:
        return

    cur_bidcnt = bids_cnt[colAuctionId - 1]
    #    print("curbid cnt",colAuctionId,cur_bidcnt)

    idx1, idx2 = reserveInput(web3, appContract, 2, account)
    mask1, mask2 = asyncio.run(get_inputmasks(players(appContract), f'{idx1},{idx2}', threshold(appContract)))
    maskedP, maskedAmt = (price + mask1) % prime, (amt + mask2) % prime

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.submitBids(colAuctionId, idx1, maskedP, idx2, maskedAmt).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)

    while True:
        time.sleep(1)
        status = appContract.functions.status(colAuctionId).call()
        if status - 2 > cur_bidcnt:
            bids_cnt[colAuctionId - 1] = status - 2
            return
        if status == 1:
            return


def initClient(appContract, account, token_addr):
    web3.eth.defaultAccount = account.address
    tx = appContract.functions.initClient(token_addr).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)


if __name__ == '__main__':
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)

    client_1 = getAccount(web3, f'/opt/poa/keystore/client_3/')
    client_2 = getAccount(web3, f'/opt/poa/keystore/client_4/')
    client_3 = getAccount(web3, f'/opt/poa/keystore/client_5/')
    client_4 = getAccount(web3, f'/opt/poa/keystore/client_6/')
    client_5 = getAccount(web3, f'/opt/poa/keystore/client_7/')
    client_6 = getAccount(web3, f'/opt/poa/keystore/client_8/')

    clients = [client_1, client_2, client_3, client_4, client_5, client_6]
    n_cli = len(clients)
    n_token = 4
    for i in range(n_cli):
        for token_id in range(n_token):
            print(i, token_id)
            initClient(appContract, clients[i], token_addrs[token_id])

    aucapp_addr = getAccount(web3, f'/opt/poa/keystore/client_2/').address

    # auction1 success

    totalAmt1 = 20
    StartPrice1 = 100
    FloorPrice1 = 10
    colAuctionId1 = createAuction(appContract, StartPrice1, FloorPrice1, totalAmt1, token_addrs[1], aucapp_addr,
                                  client_1)
    print('new Auction id:', colAuctionId1)

    # auction2 faild due to total amt > sum of all bidder's amt

    totalAmt2 = 40
    StartPrice2 = 100
    FloorPrice2 = 10
    colAuctionId2 = createAuction(appContract, StartPrice2, FloorPrice2, totalAmt2, token_addrs[2], aucapp_addr,
                                  client_1)
    print('new Auction id:', colAuctionId2)

    price11 = 60
    Amt11 = 2
    submitBids(appContract, colAuctionId1, price11, Amt11, client_1)
    print('finished input client_1 AuctionId:', colAuctionId1)

    submitBids(appContract, colAuctionId2, price11, Amt11, client_1)
    print('finished input client_1 AuctionId:', colAuctionId2)

    price12 = 50
    Amt12 = 10
    submitBids(appContract, colAuctionId1, price12, Amt12, client_2)
    print('finished input client_2 AuctionId:', colAuctionId1)

    submitBids(appContract, colAuctionId2, price12, Amt12, client_2)
    print('finished input client_2 AuctionId:', colAuctionId2)

    # auction3 failed due to the FloorPrice is too high

    totalAmt3 = 20
    StartPrice3 = 100
    FloorPrice3 = 50
    colAuctionId3 = createAuction(appContract, StartPrice3, FloorPrice3, totalAmt3, token_addrs[3], aucapp_addr,
                                  client_1)
    print('new Auction id:', colAuctionId3)

    submitBids(appContract, colAuctionId3, price11, Amt11, client_1)
    print('finished input client_1 AuctionId:', colAuctionId3)

    submitBids(appContract, colAuctionId3, price12, Amt12, client_2)
    print('finished input client_2 AuctionId:', colAuctionId3)

    price13 = 30
    Amt13 = 6
    submitBids(appContract, colAuctionId1, price13, Amt13, client_3)
    print('finished input client_3 AuctionId:', colAuctionId1)

    submitBids(appContract, colAuctionId2, price13, Amt13, client_3)
    print('finished input client_3 AuctionId:', colAuctionId2)

    submitBids(appContract, colAuctionId3, price13, Amt13, client_3)
    print('finished input client_3 AuctionId:', colAuctionId3)

    price14 = 70
    Amt14 = 7
    submitBids(appContract, colAuctionId1, price14, Amt14, client_4)
    print('finished input client_4 AuctionId:', colAuctionId1)

    submitBids(appContract, colAuctionId2, price14, Amt14, client_4)
    print('finished input client_4 AuctionId:', colAuctionId2)

    submitBids(appContract, colAuctionId3, price14, Amt14, client_4)
    print('finished input client_4 AuctionId:', colAuctionId3)

    price15 = 20
    Amt15 = 9
    submitBids(appContract, colAuctionId1, price15, Amt15, client_5)
    print('finished input client_5 AuctionId:', colAuctionId1)

    submitBids(appContract, colAuctionId2, price15, Amt15, client_5)
    print('finished input client_5 AuctionId:', colAuctionId2)

    submitBids(appContract, colAuctionId3, price15, Amt15, client_5)
    print('finished input client_5 AuctionId:', colAuctionId3)

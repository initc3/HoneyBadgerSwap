import asyncio
from curses.ascii import SP
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import fp,parse_contract, getAccount, players, prime, sign_and_send

contract_name = 'colAuction'

liveAuct = []
checkCnt = []

# means I'll buy up to $amt if the prices reaches $price or below
def scheduleCheck(appContract,colAuctionId,account):
    cur_check = checkCnt[colAuctionId-1]
#    print("schedule check aucId:",colAuctionId,"checkCnt:",cur_check)

    curTime = web3.eth.block_number
#    print("curTime: ",curTime)

    lastTime = appContract.functions.checkTime(colAuctionId).call()
    if(lastTime + 10 >= curTime):
        return


    web3.eth.defaultAccount = account.address
    tx = appContract.functions.scheduleCheck(colAuctionId).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)

    checkId = appContract.functions.checkNum(colAuctionId).call()
    if checkId >= cur_check:
        checkCnt[colAuctionId-1] = checkId
        
    

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)



    client_1 = getAccount(web3,f'/opt/poa/keystore/client_1/')

    cur_n = 0

    ccnt = 0

    while True:
        #refresh liveAuct[]
        while True:
            status = appContract.functions.status(cur_n+1).call()
            if status >= 1:
                cur_n += 1
                liveAuct.append(cur_n)
                checkCnt.append(0)
                print("cur Live Auct Id(adding):",liveAuct)
            else:
                break
        
        for aucId in liveAuct:
            status = appContract.functions.status(aucId).call()
            if status == 1:
                liveAuct.remove(aucId)
                print("cur Live Auct Id(removing):",liveAuct)
        
        if ccnt == 1000:
            ccnt = 0
            print("curLiveAuctionId:",liveAuct)

        for aucId in liveAuct:
#            print("start check:",aucId)
            scheduleCheck(appContract,aucId,client_1)

        time.sleep(5)
        ccnt = ccnt+1


        



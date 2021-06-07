import asyncio
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks
from ratel.src.python.deploy import url, parse_contract, appAddress, reserveInput, getAccount
from ratel.src.python.utils import fp

contract_name = 'review'

def initSession(appContract):
    tx_hash = appContract.functions.initSession(timeRegistration, numReviewer, reviewerAddrs).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    log = appContract.events.NewSession().processReceipt(receipt)
    sessionId = log[0]['args']['sessionId']
    return sessionId

def registerPaper(appContract, sessionId, paperNum):
    for i in range(paperNum):
        tx_hash = appContract.functions.registerPaper(sessionId).transact()
        web3.eth.wait_for_transaction_receipt(tx_hash)

def assignReviewer(appContract, sessionId, timeReview, reviewersPerPaper, paperNum):
    dueRegistration = appContract.functions.dueRegistration(sessionId).call()
    while True:
        time.sleep(1)
        blkNum = web3.eth.get_block_number()
        print('!!!! assignReviewer', dueRegistration, blkNum)
        if blkNum >= dueRegistration:
            break

    reviewers = []
    for i in range(paperNum):
        for j in range(reviewersPerPaper):
            reviewers.append(j)
    tx_hash = appContract.functions.assignReviewer(sessionId, timeReview, reviewersPerPaper, reviewers).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

    dueReview = appContract.functions.dueReview(sessionId).call()
    print('!!!! dueReview', dueReview)

def peerReview(appContract, sessionId, paperNum, reviewersPerPaper):
    for i in range(paperNum):
        for j in range(reviewersPerPaper):
            score = i + j
            print('!!!! peerReview', i, j, score)

            account = getAccount(web3, f'/opt/poa/keystore/server_{j}/')
            idx = reserveInput(web3, appContract, 1, account)[0]
            mask = asyncio.run(get_inputmasks(f'{idx}'))[0]
            maskedScore = score + mask

            tx = appContract.functions.peerReview(sessionId, i, idx, maskedScore).buildTransaction({'from': account.address, 'gas': 1000000, 'nonce': web3.eth.get_transaction_count(account.address)})
            signedTx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)
            web3.eth.send_raw_transaction(signedTx.rawTransaction)
            web3.eth.wait_for_transaction_receipt(signedTx.hash)

def calcResult(appContract, sessionId, threshold):
    dueReview = appContract.functions.dueReview(sessionId).call()
    while True:
        time.sleep(1)
        blkNum = web3.eth.get_block_number()
        print('!!!! calcResult', dueReview, blkNum)
        if blkNum >= dueReview:
            break

    account = getAccount(web3, f'/opt/poa/keystore/server_0/')
    idx = reserveInput(web3, appContract, 1, account)[0]
    mask = asyncio.run(get_inputmasks(f'{idx}'))[0]
    maskedThreshold = int(threshold * fp) + mask

    tx_hash = appContract.functions.calcResult(sessionId, idx, maskedThreshold).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    timeRegistration = 10
    timeReview = 30
    numReviewer = 2
    reviewerAddrs = []
    for reviewer in range(numReviewer):
        account = getAccount(web3, f'/opt/poa/keystore/server_{reviewer}/')
        reviewerAddrs.append(account.address)

    sessionId = initSession(appContract)

    reviewerNum = 2
    paperNum = 3
    registerPaper(appContract, sessionId, paperNum)

    assignReviewer(appContract, sessionId, timeReview, reviewerNum, paperNum)

    peerReview(appContract, sessionId, paperNum, reviewerNum)

    threshold = 1.5
    calcResult(appContract, sessionId, threshold)


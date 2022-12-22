import asyncio
import time
import json

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput, generate_zkrp_mul
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import (
    parse_contract,
    getAccount,
    players,
    prime,
    sign_and_send,
    threshold,
    get_zkrp,
)


contract_name = "rockPaperScissors"


def createGame(appContract, value1, account):
    print(f'**** CreateGame {value1}')

    tmp = 123
    proof1, commitment1, blinding1 = get_zkrp(value1, '>=', 1)
    
    proof2, commitment2, blinding2 = get_zkrp(value1*value1,'>=',0)

    idx1, idx2, idx3 = reserveInput(web3, appContract, 3, account)
    mask1, mask2, mask3 = asyncio.run(get_inputmasks(players(appContract), f'{idx1},{idx2},{idx3}', threshold(appContract)))
    maskedvalue1, maskedvalue2, maskedvalue3 = (value1 + mask1) % prime, (blinding1 + mask2) % prime, (blinding2 + mask3) % prime
    
    zkp1 = [idx2, maskedvalue2, proof1, commitment1]
    zkp2 = [idx3, maskedvalue3, proof2, commitment2]
    zkps = json.dumps([zkp1,zkp2])

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.createGame(idx1, maskedvalue1, zkps).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    receipt = sign_and_send(tx, web3, account)

    log = appContract.events.CreateGame().processReceipt(receipt)
    gameId = log[0]["args"]["gameId"]
    while True:
        time.sleep(1)
        status = appContract.functions.status(gameId).call()
        if status == 1:
            return gameId


def joinGame(appContract, gameId, value2, account):
    print(f'**** JoinGame {value2}')

    proof1, commitment1, blinding1 = get_zkrp(value2, '>=', 1)
    proof2, commitment2, blinding2 = get_zkrp(value2, '<=', 3)

    idx1, idx2, idx3 = reserveInput(web3, appContract, 3, account)
    mask1, mask2, mask3 = asyncio.run(get_inputmasks(players(appContract), f'{idx1},{idx2},{idx3}', threshold(appContract)))
    maskedvalue1, maskedvalue2, maskedvalue3 = (value2 + mask1) % prime, (blinding1 + mask2) % prime, (blinding2 + mask3) % prime

    zkp1 = [idx2,maskedvalue2,proof1,commitment1]
    zkp2 = [idx3,maskedvalue3,proof2,commitment2]
    zkps = json.dumps([zkp1,zkp2])

    web3.eth.defaultAccount = account.address
    tx = appContract.functions.joinGame(gameId, idx1, maskedvalue1, zkps).buildTransaction(
        {"nonce": web3.eth.get_transaction_count(web3.eth.defaultAccount)}
    )
    sign_and_send(tx, web3, account)

    while True:
        time.sleep(1)
        status = appContract.functions.status(gameId).call()
        if status == 2:
            return


def startRecon(appContract, gameId, account):
    web3.eth.defaultAccount = account.address
    tx = appContract.functions.startRecon(gameId).buildTransaction(
        {"nonce": web3.eth.get_transaction_count(web3.eth.defaultAccount)}
    )
    sign_and_send(tx, web3, account)

    while True:
        winner = appContract.functions.winners(gameId).call()
        if winner != "":
            print("!!!! winner", winner)
            break
        time.sleep(1)

if __name__ == "__main__":
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract(contract_name)
    appContract = web3.eth.contract(address=app_addr, abi=abi)

    client_1 = getAccount(web3, f"/opt/poa/keystore/client_1/")
    client_2 = getAccount(web3, f"/opt/poa/keystore/client_2/")
            
    gameId = createGame(appContract, 1, client_1)
    joinGame(appContract, gameId, 1, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 1, client_1)
    joinGame(appContract, gameId, 2, client_2)
    startRecon(appContract, gameId, client_2)

    gameId = createGame(appContract, 1, client_1)
    joinGame(appContract, gameId, 3, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 2, client_1)
    joinGame(appContract, gameId, 1, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 2, client_1)
    joinGame(appContract, gameId, 2, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 2, client_1)
    joinGame(appContract, gameId, 3, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 3, client_1)
    joinGame(appContract, gameId, 1, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 3, client_1)
    joinGame(appContract, gameId, 2, client_2)
    startRecon(appContract, gameId, client_1)

    gameId = createGame(appContract, 3, client_1)
    joinGame(appContract, gameId, 3, client_2)
    startRecon(appContract, gameId, client_1)
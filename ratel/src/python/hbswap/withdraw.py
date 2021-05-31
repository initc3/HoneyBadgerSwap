import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.deploy import parse_contract, tokenAddress, appAddress, url
from ratel.src.python.utils import fp


def approve(tokenContract, receiver, amt):
    tx_hash = tokenContract.functions.approve(receiver, int(amt * fp)).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

def withdraw(appContract, token, amt):
    amt = int(amt * fp)

    tx_hash = appContract.functions.secretWithdraw(token, amt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    # receipt = web3.eth.get_transaction_receipt(tx_hash)
    # log = appContract.events.SecretWithdraw().processReceipt(receipt)
    # print(log)

    while True:
        balance = appContract.functions.publicBalance(token, web3.eth.defaultAccount).call()
        print('balance', balance)
        if balance >= amt:
            break
        time.sleep(2)

    tx_hash = appContract.functions.publicWithdraw(token, amt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    balance = appContract.functions.publicBalance(token, web3.eth.defaultAccount).call()
    print('after balance', balance)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    # withdraw(appContract, ETH, 1)
    withdraw(appContract, tokenAddress, 1)
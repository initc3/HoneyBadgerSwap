import sys
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.deploy import parse_contract, url, app_addr, token_addrs
from ratel.src.python.utils import fp, getAccount, sign_and_send


def withdraw(appContract, tokenAddr, withdrawAmt):
    prevBalance = appContract.functions.publicBalance(tokenAddr, web3.eth.defaultAccount).call()
    print('!!!! prevBalance', prevBalance)
    tx = appContract.functions.secretWithdraw(tokenAddr, int(withdrawAmt * fp)).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)

    while True:
        balance = appContract.functions.publicBalance(tokenAddr, web3.eth.defaultAccount).call()
        print('!!!! balance', balance)
        if balance > prevBalance:
            break
        time.sleep(2)

    tx = appContract.functions.publicWithdraw(tokenAddr, int(withdrawAmt * fp)).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)
    balance = appContract.functions.publicBalance(tokenAddr, web3.eth.defaultAccount).call()
    print('!!!! after balance', balance)


if __name__=='__main__':
    client_id = int(sys.argv[1])
    token_id = int(sys.argv[2])
    withdrawAmt = int(sys.argv[3])
    print('token_id', token_id)

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    web3.eth.defaultAccount = account.address
    withdraw(appContract, token_addrs[token_id], withdrawAmt)
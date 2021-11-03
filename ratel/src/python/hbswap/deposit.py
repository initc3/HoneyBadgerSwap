import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.deploy import url, ETH, addrs
from ratel.src.python.utils import fp, decimal, getAccount, sign_and_send, parse_contract

def approve(tokenContract, receiver, amt):
    tx = tokenContract.functions.approve(receiver, int(amt * fp)).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)

def deposit(appContract, tokenAddr, depositAmt):
    if tokenAddr == ETH:
        tx = appContract.functions.publicDeposit(tokenAddr, int(depositAmt * fp)).buildTransaction({
            'value': int(depositAmt * decimal),
            'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
        })
        sign_and_send(tx, web3, account)
    else:
        abi, bytecode = parse_contract('Token')
        tokenContract = web3.eth.contract(address=tokenAddr, abi=abi)
        approve(tokenContract, appContract.address, int(depositAmt * decimal))

        tx = appContract.functions.publicDeposit(tokenAddr, int(depositAmt * fp)).buildTransaction({
            'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
        })
        sign_and_send(tx, web3, account)

    tx = appContract.functions.secretDeposit(tokenAddr, int(depositAmt * fp)).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)

if __name__=='__main__':
    client_id = int(sys.argv[1])
    tokenAddr = sys.argv[2]
    depositAmt = int(sys.argv[3])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=addrs[0], abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    web3.eth.defaultAccount = account.address
    deposit(appContract, tokenAddr, depositAmt)
    print('**** deposit finished')
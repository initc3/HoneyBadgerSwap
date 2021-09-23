import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.deploy import parse_contract, appAddress, url, ETH
from ratel.src.python.utils import fp, decimal

def approve(tokenContract, receiver, amt):
    tx_hash = tokenContract.functions.approve(receiver, int(amt * fp)).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

def deposit(appContract, tokenAddr, depositAmt):
    if tokenAddr == ETH:
        tx_hash = appContract.functions.publicDeposit(tokenAddr, int(depositAmt * fp)).transact({'from':web3.eth.defaultAccount, 'value': int(depositAmt * decimal)})
        web3.eth.wait_for_transaction_receipt(tx_hash)
    else:
        abi, bytecode = parse_contract('Token')
        tokenContract = web3.eth.contract(address=tokenAddr, abi=abi)
        approve(tokenContract, appContract.address, int(depositAmt * decimal))

        tx_hash = appContract.functions.publicDeposit(tokenAddr, int(depositAmt * fp)).transact()
        web3.eth.wait_for_transaction_receipt(tx_hash)

    tx_hash = appContract.functions.secretDeposit(tokenAddr, int(depositAmt * fp)).transact()
    print('**** deposit finished')
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    tokenAddr = sys.argv[1]
    depositAmt = int(sys.argv[2])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    deposit(appContract, tokenAddr, depositAmt)
from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.deploy import parse_contract, ETH, tokenAddress, appAddress, url
from ratel.src.python.utils import fp, decimal

def approve(tokenContract, receiver, amt):
    tx_hash = tokenContract.functions.approve(receiver, int(amt * fp)).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

def deposit(appContract, token, amt, tokenContract):
    if token == ETH:
        tx_hash = appContract.functions.publicDeposit(token, int(amt * fp)).transact({'from':web3.eth.defaultAccount, 'value': int(amt * decimal)})
        web3.eth.wait_for_transaction_receipt(tx_hash)
    else:
        approve(tokenContract, appContract.address, int(amt * decimal))
        tx_hash = appContract.functions.publicDeposit(token, int(amt * fp)).transact()
        web3.eth.wait_for_transaction_receipt(tx_hash)

    tx_hash = appContract.functions.secretDeposit(token, int(amt * fp)).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('Token')
    tokenContract = web3.eth.contract(address=tokenAddress, abi=abi)
    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    deposit(appContract, ETH, 10, tokenContract)
    deposit(appContract, tokenAddress, 10, tokenContract)
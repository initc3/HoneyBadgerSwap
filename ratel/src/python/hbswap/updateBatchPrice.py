import sys

from ratel.src.python.deploy import url, app_addr, token_addrs
from ratel.src.python.utils import getAccount, parse_contract, sign_and_send
from web3 import Web3
from web3.middleware import geth_poa_middleware


def updateBatchPrice(appContract):
    tx = appContract.functions.updateBatchPrice().buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    sign_and_send(tx, web3, account)


if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/admin/')
    web3.eth.defaultAccount = account.address

    updateBatchPrice(appContract)
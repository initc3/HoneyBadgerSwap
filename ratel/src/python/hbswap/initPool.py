from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.deploy import url, parse_contract, appAddress, tokenAddress, ETH
from ratel.src.python.utils import fp


def initPool(appContract, tokenA, tokenB, amtA, amtB):
    tx_hash = appContract.functions.initPool(tokenA, tokenB, int(amtA * fp), int(amtB * fp)).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('Test')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    initPool(appContract, ETH, tokenAddress, 1, 1)
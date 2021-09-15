import asyncio

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks
from ratel.src.python.deploy import url, parse_contract, appAddress, tokenAddress, ETH, reserveInput, getAccount
from ratel.src.python.utils import fp

def addLiquidity(appContract, tokenA, tokenB, amtA, amtB, account):
    amtA = int(amtA * fp)
    amtB = int(amtB * fp)
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(appContract, f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = amtA + maskA, amtB + maskB
    tx_hash = appContract.functions.addLiquidity(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('Hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/server_0/')
    addLiquidity(appContract, ETH, tokenAddress, 1, 1, account)
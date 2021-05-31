import asyncio

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks
from ratel.src.python.deploy import url, parse_contract, appAddress, tokenAddress, ETH, reserveInput
from ratel.src.python.utils import fp, blsPrime


def trade(appContract, tokenA, tokenB, amtA, amtB):
    amtA = int(amtA * fp)
    amtB = int(amtB * fp)
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2)
    maskA, maskB = asyncio.run(get_inputmasks(f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = (amtA + maskA) % blsPrime, (amtB + maskB) % blsPrime
    tx_hash = appContract.functions.trade(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    trade(appContract, ETH, tokenAddress, 0.01, -0.1)
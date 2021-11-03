import asyncio
import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, parse_contract, appAddress
from ratel.src.python.utils import fp, blsPrime, getAccount


# def addLiquidity(appContract, tokenA, tokenB, amtA, amtB, account):
#     amtA = int(amtA * fp)
#     amtB = int(amtB * fp)
#     idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
#     maskA, maskB = asyncio.run(get_inputmasks(appContract, f'{idxAmtA},{idxAmtB}'))
#     maskedAmtA, maskedAmtB = amtA + maskA, amtB + maskB
#     tx_hash = appContract.functions.addLiquidity(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
#     web3.eth.wait_for_transaction_receipt(tx_hash)

def addLiquidity(appContract, tokenA, tokenB, amtA, amtB, account):
    amtA = int(amtA * fp)
    amtB = int(amtB * fp)
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(appContract, f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = (amtA + maskA) % blsPrime, (amtB + maskB) % blsPrime
    tx_hash = appContract.functions.addLiquidity(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    print(receipt)

if __name__=='__main__':
    token_A_addr = sys.argv[1]
    token_B_addr = sys.argv[2]
    token_A_amt = int(sys.argv[3])
    token_B_amt = int(sys.argv[4])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    # addLiquidity(appContract, ETH, tokenAddress, 1, 1, account)
    account = getAccount(web3, f'/opt/poa/keystore/server_0/')
    addLiquidity(appContract, token_A_addr, token_B_addr, token_A_amt, token_B_amt, account)
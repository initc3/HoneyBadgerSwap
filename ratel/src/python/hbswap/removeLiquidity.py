import asyncio

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import get_inputmasks
from ratel.src.python.deploy import url, parse_contract, appAddress, tokenAddress, ETH, reserveInput, getAccount
from ratel.src.python.utils import fp

def removeLiquidity(appContract, tokenA, tokenB, amt, account):
    amt = int(amt * fp)
    idx = reserveInput(web3, appContract, 1, account)[0]
    print(idx)
    mask = asyncio.run(get_inputmasks(appContract, f'{idx}'))[0]
    maskedAmt = amt + mask
    tx_hash = appContract.functions.removeLiquidity(tokenA, tokenB, idx, maskedAmt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

if __name__=='__main__':
    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/server_0/')
    removeLiquidity(appContract, ETH, tokenAddress, 0.2, account)
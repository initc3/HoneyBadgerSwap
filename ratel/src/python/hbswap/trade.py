import asyncio
import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, addrs
from ratel.src.python.utils import fp, blsPrime, getAccount, sign_and_send, parse_contract


def trade(appContract, tokenA, tokenB, amtA, amtB, account):
    amtA = int(amtA * fp)
    amtB = int(amtB * fp)
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(appContract, f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = (amtA + maskA) % blsPrime, (amtB + maskB) % blsPrime
    tx = appContract.functions.trade(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    tx_hash = sign_and_send(tx, web3, account)
    receipt = web3.eth.get_transaction_receipt(tx_hash)
    print(receipt)

if __name__=='__main__':
    client_id = int(sys.argv[1])
    tokenA = sys.argv[2]
    tokenB = sys.argv[3]
    amtA = float(sys.argv[4])
    amtB = float(sys.argv[5])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=addrs[0], abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    web3.eth.defaultAccount = account.address

    trade(appContract, tokenA, tokenB, amtA, amtB, account)

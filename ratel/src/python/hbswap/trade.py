import asyncio
import sys
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.Client import get_inputmasks, reserveInput
from ratel.src.python.deploy import url, app_addr, token_addrs
from ratel.src.python.utils import fp, blsPrime, getAccount, sign_and_send, parse_contract, players


def trade(appContract, tokenA, tokenB, amtA, amtB, account):
    amtA = int(amtA * fp)
    amtB = int(amtB * fp)
    idxAmtA, idxAmtB = reserveInput(web3, appContract, 2, account)
    maskA, maskB = asyncio.run(get_inputmasks(players(appContract), f'{idxAmtA},{idxAmtB}'))
    maskedAmtA, maskedAmtB = (amtA + maskA) % blsPrime, (amtB + maskB) % blsPrime
    tx = appContract.functions.trade(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).buildTransaction({
        'nonce': web3.eth.get_transaction_count(web3.eth.defaultAccount)
    })
    receipt = sign_and_send(tx, web3, account)
    log = appContract.events.Trade().processReceipt(receipt)[0]
    print(log['args'])
    seqTrade = log['args']['seqTrade']
    with open('ratel/benchmark/data/gas.csv', 'a') as f:
        f.write(f"trade\t{seqTrade}\t"
                f"client\t{client_id}\t"
                f"tokenA\t{tokenA}\t"
                f"tokenB\t{tokenB}\t"
                f"gasUsed\t{receipt['gasUsed']}\t"
                f"{time.perf_counter()}\n")

if __name__=='__main__':
    client_id = int(sys.argv[1])
    tokenA = token_addrs[int(sys.argv[2])]
    tokenB = token_addrs[int(sys.argv[3])]
    amtA = float(sys.argv[4])
    amtB = float(sys.argv[5])
    repetition = int(sys.argv[6])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    web3.eth.defaultAccount = account.address

    for i in range(repetition):
        trade(appContract, tokenA, tokenB, amtA, amtB, account)
        trade(appContract, tokenA, tokenB, amtB, amtA, account)

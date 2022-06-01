import asyncio
import re
import sys
import time

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.Client import reserveInput, get_inputmasks
from ratel.src.python.deploy import token_addrs, url, app_addr
from ratel.src.python.utils import parse_contract, getAccount, fp, players, threshold, prime


async def trade(appContract, tokenA, tokenB, amtA, amtB, account, web3, nonce, idxAmtA, idxAmtB, maskA, maskB):
    amtA = int(amtA * fp)
    amtB = int(amtB * fp)
    maskedAmtA, maskedAmtB = (amtA + maskA) % prime, (amtB + maskB) % prime
    tx = appContract.functions.trade(tokenA, tokenB, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB).buildTransaction({
        'nonce': nonce
    })

    signed_tx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)
    tx_hash = web3.eth.send_raw_transaction(signed_tx.rawTransaction)

    return tx_hash


async def confirm(tx_hash, web3, appContract, timestamp, actual_time, init_time):
    while True:
        try:
            receipt = web3.eth.getTransactionReceipt(tx_hash)
            log = appContract.events.Trade().processReceipt(receipt)[0]
            print(log['args'])
            seqTrade = log['args']['seqTrade']
            dif = actual_time - init_time
            with open('ratel/benchmark/data/gas.csv', 'a') as f:
                f.write(f'{seqTrade}\t{actual_time}\t'
                        f'{timestamp}\t{dif}\t{dif - timestamp}\n')
            break

        except:
            await asyncio.sleep(1)


async def main(loop):
    client_id = 1
    tokenA = token_addrs[0]
    tokenB = token_addrs[1]
    amtA = 0.5
    amtB = -1

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)

    account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    web3.eth.defaultAccount = account.address

    l = len(timestamp_list)
    batch = 10
    indexes = reserveInput(web3, appContract, 2 * (l % batch), account)
    for i in range(l // batch):
        indexes.extend(reserveInput(web3, appContract, 2 * batch, account))

    st = f'{indexes[0]}'
    for index in indexes[1:]:
        st += f',{index}'
    masks = await get_inputmasks(players(appContract), st, threshold(appContract))

    nonce = web3.eth.get_transaction_count(web3.eth.defaultAccount)
    init_time = time.perf_counter()
    for i, timestamp in enumerate(timestamp_list):
        interval = timestamp - (time.perf_counter() - init_time)
        if interval > 0:
            time.sleep(interval)
        actual_time = time.perf_counter()

        tx_hash = await trade(appContract, tokenA, tokenB, amtA, amtB, account, web3, nonce, indexes[2 * i], indexes[2 * i + 1], masks[2 * i], masks[2 * i + 1])
        amtA, amtB = amtB, amtA
        nonce += 1

        loop.create_task(confirm(tx_hash, web3, appContract, timestamp, actual_time, init_time))

    await asyncio.gather(*asyncio.all_tasks())


if __name__ == '__main__':
    pool_name = sys.argv[1]
    duration = int(sys.argv[2])

    ratel / benchmark / src / swap / run.py

    loop = asyncio.get_event_loop()
    loop.run_until_complete(main(loop))
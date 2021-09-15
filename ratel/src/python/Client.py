import asyncio
import re

from aiohttp import ClientSession
from ratel.src.python.utils import players, http_port, http_host, get_inverse, blsPrime

def reserveInput(web3, appContract, num, account):
    tx = appContract.functions.reserveInput(num).buildTransaction({'from': account.address, 'gas': 1000000, 'nonce': web3.eth.get_transaction_count(account.address)})
    signedTx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)
    web3.eth.send_raw_transaction(signedTx.rawTransaction)
    web3.eth.wait_for_transaction_receipt(signedTx.hash)
    receipt = web3.eth.get_transaction_receipt(signedTx.hash)
    log = appContract.events.InputMask().processReceipt(receipt)
    return log[0]['args']['inpusMaskIndexes']

async def send_request(url):
    async with ClientSession() as session:
        async with session.get(url) as resp:
            json_response = await resp.json()
            return json_response

def interpolate(cur_players, shares):
    inputmask = 0
    for i in range(1, cur_players + 1):
        tot = 1
        for j in range(1, cur_players + 1):
            if i == j:
                continue
            tot = tot * j * get_inverse(j - i) % blsPrime
        inputmask = (inputmask + shares[i - 1] * tot) % blsPrime
    return inputmask

async def req_inputmask_shares(host, port, inputmask_idxes):
    url = f"http://{host}:{port}/inputmasks/{inputmask_idxes}"
    print(url)
    result = await send_request(url)
    return re.split(",", result["inputmask_shares"])

async def get_inputmasks(contract, inputmask_idxes):
    cur_players = players(contract)
    tasks = []
    for serverID in range(cur_players):
        task = asyncio.ensure_future(
            req_inputmask_shares(http_host, http_port + serverID, inputmask_idxes)
        )
        tasks.append(task)

    for task in tasks:
        await task

    inputmask_shares = []
    for task in tasks:
        inputmask_shares.append(task.result())

    inputmasks = []
    for i in range(len(tasks[0].result())):
        shares = []
        for j in range(cur_players):
            shares.append(int(inputmask_shares[j][i]))
        inputmasks.append(interpolate(cur_players, shares))
    return inputmasks
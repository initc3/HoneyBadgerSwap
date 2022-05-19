import asyncio
import re

from aiohttp import ClientSession
from ratel.src.python.utils import http_port, http_host, get_inverse, prime, sign_and_send


def reserveInput(web3, appContract, num, account):
    tx = appContract.functions.reserveInput(num).buildTransaction({'from': account.address, 'gas': 1000000, 'nonce': web3.eth.get_transaction_count(account.address)})
    receipt = sign_and_send(tx, web3, account)
    log = appContract.events.InputMask().processReceipt(receipt)
    return log[0]['args']['inpusMaskIndexes']


def reconstruction(shares):
    value = 0
    n = len(shares)
    for i in range(n):
        tot = 1
        for j in range(n):
            if i == j:
                continue
            tot = tot * shares[j][0] * get_inverse(shares[j][0] - shares[i][0]) % prime
        value = (value + shares[i][1] * tot) % prime
    return value


def interpolate(shares, t):
    value = reconstruction(shares[:t + 1])
    n = len(shares)
    for i in range(t + 2, n):
        check = reconstruction(shares[:i])
        if check != value:
            print('mac_fail')
            return 0
    return value % prime


def batch_interpolate(results, threshold):
    res = []
    num = len(results[0])
    players = len(results)
    for i in range(num):
        shares = []
        for j in range(players):
            result = int(results[j][i])
            if result != 0:
                shares.append((j + 1, result))
        res.append(interpolate(shares, threshold))
    return res


async def send_request(url):
    async with ClientSession() as session:
        async with session.get(url) as resp:
            json_response = await resp.json()
            return json_response


async def send_requests(players, request):
    tasks = []
    for server_id in range(players):
        task = send_request(f"http://{http_host}:{http_port + server_id}/{request}")
        tasks.append(task)

    results = await asyncio.gather(*tasks)
    return results


async def get_inputmasks(players, inputmask_idxes, threshold):
    request = f"inputmasks/{inputmask_idxes}"
    results = await send_requests(players, request)
    for i in range(len(results)):
        results[i] = re.split(",", results[i]["inputmask_shares"])

    inputmasks = batch_interpolate(results, threshold)

    return inputmasks

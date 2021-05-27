import asyncio
import re

from aiohttp import ClientSession

from ratel.src.python.utils import players, http_port, http_host, get_inverse, blsPrime


async def send_request(url):
    async with ClientSession() as session:
        async with session.get(url) as resp:
            json_response = await resp.json()
            return json_response

def interpolate(shares):
    inputmask = 0
    for i in range(1, players + 1):
        tot = 1
        for j in range(1, players + 1):
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

async def get_inputmasks(inputmask_idxes):
    tasks = []
    for serverID in range(players):
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
        for j in range(players):
            shares.append(int(inputmask_shares[j][i]))
        inputmasks.append(interpolate(shares))
    return inputmasks
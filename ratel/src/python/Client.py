import asyncio
import re
import json

from aiohttp import ClientSession
from ratel.src.python.utils import http_port, http_host, get_inverse, prime, sign_and_send
from zkrp_pyo3 import zkrp_prove_mul, zkrp_verify_mul


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

async def get_serverval(players, server_idxes, threshold):
    request = f"serverval/{server_idxes}"
    results = await send_requests(players, request)
    for i in range(len(results)):
        results[i] = re.split(",", results[i]["serverval_shares"])

    inputserverval = batch_interpolate(results, threshold)

    return inputserverval

async def get_zkrp_blinding_info(players, num, threshold):
    request_blinding = f"zkrp_blinding_shares/{num}"
    blinding_res = await send_requests(players, request_blinding)
    
    for i in range(len(blinding_res)):
        blinding_res[i] = re.split(',',blinding_res[i]["zkrp_blinding_shares"])
    blinding_prime = batch_interpolate(blinding_res, threshold)

    request_agg = f"zkrp_new_agg_com/{num}"
    comm_res = await send_requests(players, request_agg)
    comm_res = comm_res[0]["zkrp_blinding_info_2"]
    comm_res = re.split(";", comm_res)
    for i in range(len(comm_res)):
        comm_res[i] = re.split(',', comm_res[i][1:-1])
        for j in range(len(comm_res[i])):
            comm_res[i][j] = int(comm_res[i][j])

    return blinding_prime, comm_res

async def generate_zkrp_mul(players, x, y, threshold):
    blinding_prime_list, blinding_comm_list = await get_zkrp_blinding_info(players, 2, threshold)

    print('blinding prime list:', blinding_prime_list)
    print('blinding com list:', blinding_comm_list)

    rx_prime, ry_prime = blinding_prime_list[0], blinding_prime_list[1]
    cx_bytes, cy_bytes = blinding_comm_list[0], blinding_comm_list[1]

    rx_prime_bytes =  rx_prime.to_bytes((rx_prime.bit_length() + 7) // 8, 'little')
    ry_prime_bytes =  ry_prime.to_bytes((ry_prime.bit_length() + 7) // 8, 'little')

    mx_prime, my_prime, sx, sy_prime,  = zkrp_prove_mul(x, y, rx_prime_bytes,ry_prime_bytes)


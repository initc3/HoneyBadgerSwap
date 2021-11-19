import asyncio
import re
import sys

from web3 import Web3

from ratel.genfiles.python.hbswapRecover import recover
from ratel.src.python.Client import send_requests, batch_interpolate
from ratel.src.python.Server import Server
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import list_to_str, parse_contract

async def testing():
    await asyncio.sleep(3)
    seq_num_list = [seq] * repetition

    request = f'recoverdb/{list_to_str(seq_num_list)}'
    masked_shares = await send_requests(online_players, request)
    for i in range(len(masked_shares)):
        masked_shares[i] = re.split(",", masked_shares[i]["values"])
    keys = server.collect_keys(seq_num_list)
    masked_states = batch_interpolate(masked_shares)
    state_shares = server.recover_states(masked_states)
    server.restore_db(seq_num_list, keys, state_shares)

if __name__ == '__main__':
    serverID = int(sys.argv[1])
    online_players = int(sys.argv[2])
    seq = int(sys.argv[3])
    repetition = int(sys.argv[4])

    web3 = Web3(Web3.WebsocketProvider(url))
    abi, bytecode = parse_contract('hbswap')
    appContract = web3.eth.contract(address=app_addr, abi=abi)
    concurrency = 1
    test=True

    server = Server(
        serverID,
        web3,
        appContract,
        None,
        None,
        concurrency,
        recover,
        test,
    )

    loop = asyncio.get_event_loop()
    
    tasks = [
        loop.create_task(server.preprocessing()),
        loop.create_task(server.http_server()),
    ]
    
    if serverID == online_players:
        tasks.append(loop.create_task(testing()))

    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()
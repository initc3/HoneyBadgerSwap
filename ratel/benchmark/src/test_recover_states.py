import asyncio
import re
import sys
import time

from web3 import Web3

from ratel.genfiles.python.hbswapRecover import recover
from ratel.src.python.Client import send_requests, batch_interpolate
from ratel.src.python.Server import Server
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.utils import list_to_str, parse_contract, trade_key_num


async def testing():
    await asyncio.sleep(3)
    seq_num_list = [seq] * repetition

    request = f'recoverdb/{list_to_str(seq_num_list)}'

    with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        f.write(f'state\t{repetition * trade_key_num}\t'
                f'stage\t1\t'
                f'{time.perf_counter()}\n')

    keys = server.collect_keys(seq_num_list)

    with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        f.write(f'state\t{repetition * trade_key_num}\t'
                f'stage\t2\t'
                f'{time.perf_counter()}\n')

    task = loop.create_task(send_requests(online_players, request))

    await server.genInputMask(len(keys))
    await task

    masked_shares = task.result()

    with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        f.write(f'state\t{repetition * trade_key_num}\t'
                f'stage\t7\t'
                f'{time.perf_counter()}\n')

    for i in range(len(masked_shares)):
        masked_shares[i] = re.split(",", masked_shares[i]["values"])
    masked_states = batch_interpolate(masked_shares)

    with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        f.write(f'state\t{repetition * trade_key_num}\t'
                f'stage\t8\t'
                f'{time.perf_counter()}\n')

    state_shares = server.recover_states(masked_states)

    with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        f.write(f'state\t{repetition * trade_key_num}\t'
                f'stage\t9\t'
                f'{time.perf_counter()}\n')

    server.restore_db(seq_num_list, keys, state_shares)

    with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        f.write(f'state\t{repetition * trade_key_num}\t'
                f'stage\t10\t'
                f'{time.perf_counter()}\n')

    print('**** test finished!')

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
        # loop.create_task(server.preprocessing()),
        loop.create_task(server.http_server()),
    ]
    
    if serverID == online_players:
        tasks.append(loop.create_task(testing()))

    loop.run_until_complete(asyncio.wait(tasks))
    loop.close()
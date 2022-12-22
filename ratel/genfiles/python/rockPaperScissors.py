import asyncio
import json
from ratel.src.python.utils import location_sharefile, prog, mpcPort, prime, sz, int_to_hex, hex_to_int, recover_input, fp, mark_finish, read_db, write_db, bytes_to_int, bytes_to_list, bytes_to_dict, int_to_bytes, list_to_bytes, dict_to_bytes, execute_cmd, MultiAcquire, sign_and_send, verify_proof
from ratel.benchmark.src.test_mpc import run_online

async def monitor(server):
    blkNum = server.web3.eth.get_block_number()
    while True:
        curBlkNum = server.web3.eth.get_block_number()
        if curBlkNum - blkNum > server.confirmation:
            logs = []

            eventFilter = server.contract.events.CreateGame.createFilter(fromBlock=blkNum, toBlock=curBlkNum - server.confirmation)
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append((log['blockNumber'], log['transactionIndex'], 'CreateGame', log))
            eventFilter = server.contract.events.JoinGame.createFilter(fromBlock=blkNum, toBlock=curBlkNum - server.confirmation)
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append((log['blockNumber'], log['transactionIndex'], 'JoinGame', log))
            eventFilter = server.contract.events.StartRecon.createFilter(fromBlock=blkNum, toBlock=curBlkNum - server.confirmation)
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append((log['blockNumber'], log['transactionIndex'], 'StartRecon', log))
            eventFilter = server.contract.events.GenInputMask.createFilter(fromBlock=blkNum, toBlock=curBlkNum - server.confirmation)
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append((log['blockNumber'], log['transactionIndex'], 'GenInputMask', log))

            logs.sort(key=lambda s:(s[0], s[1]))
            for log in logs:
                server.loop.create_task(eval(f'run{log[2]}')(server, log[3]))
            blkNum = curBlkNum - server.confirmation + 1
        else:
            await asyncio.sleep(1)


async def runGenInputMask(server, log):
    input_mask_cnt = log['args']['inputMaskCnt']
    input_mask_version = log['args']['inputMaskVersion']
    await server.gen_input_mask(input_mask_cnt, input_mask_version)


async def runCreateGame(server, log):
    seqCreateGame = log['args']['seqCreateGame']
    gameId = log['args']['gameId']
    player1 = log['args']['player1']
    idxValue1 = log['args']['idxValue1']
    maskedValue1 = log['args']['maskedValue1']
    zkpstmt0 = log['args']['zkpstmt0']

    value1 = recover_input(server.db, maskedValue1, idxValue1)

    readKeys = []
    writeKeys = [f'gameBoard_{gameId}']
    readKeys =  [k.lower() for k in readKeys]
    writeKeys =  [k.lower() for k in writeKeys]

    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0

    tasks = []
    port = mpcPort(seqCreateGame, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    for key in writeKeys:
        if key not in readKeys:
            tasks.append(server.dbLock[key].acquire())
    await asyncio.wait(tasks)

    for key in writeKeys:
        server.dbLockCnt[key] += 1

    tmpRange = 1
    pfVarvalue10 = value1 - tmpRange
    assert(await verify_proof(server, pfVarvalue10, zkpstmt0 ))

    game = {
        'player1': player1,
        'value1': value1,
    }
    print('**** game', game)
    game = dict_to_bytes(game)
    write_db(server, f'gameBoard_{gameId}', game, 0)

    curStatus = 1

    tx = server.contract.functions.statusSet(curStatus, gameId).buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})
    sign_and_send(tx, server.web3, server.account)
    print(server.contract.functions.status(gameId).call())

    mark_finish(server, seqCreateGame)


async def runJoinGame(server, log):
    seqJoinGame = log['args']['seqJoinGame']
    gameId = log['args']['gameId']
    player2 = log['args']['player2']
    idxValue2 = log['args']['idxValue2']
    maskedValue2 = log['args']['maskedValue2']
    zkpstmt0 = log['args']['zkpstmt0']
    zkpstmt1 = log['args']['zkpstmt1']

    value2 = recover_input(server.db, maskedValue2, idxValue2)

    readKeys = [f'gameBoard_{gameId}']
    writeKeys = [f'gameBoard_{gameId}']
    readKeys =  [k.lower() for k in readKeys]
    writeKeys =  [k.lower() for k in writeKeys]

    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0
    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0

    tasks = []
    port = mpcPort(seqJoinGame, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    for key in writeKeys:
        if key not in readKeys:
            tasks.append(server.dbLock[key].acquire())
    await asyncio.wait(tasks)

    for key in readKeys:
        server.dbLockCnt[key] += 1
    for key in writeKeys:
        server.dbLockCnt[key] += 1


    tmpRange =  1
    pfVarvalue20 = value2 - tmpRange
    assert(await verify_proof(server, pfVarvalue20, zkpstmt0 ))
    tmpRange =  3
    pfVarvalue21 = (prime + tmpRange - value2) % prime
    assert(await verify_proof(server, pfVarvalue21, zkpstmt1 ))

    value1 = read_db(server, f'gameBoard_{gameId}', 0)
    value1 = bytes_to_dict(value1)
    game = value1

    game['player2'] = player2
    game['value2'] = value2

    print('**** game', game)

    game = dict_to_bytes(game)
    write_db(server, f'gameBoard_{gameId}', game, 0)

    curStatus = 2

    tx = server.contract.functions.statusSet(curStatus, gameId).buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})
    sign_and_send(tx, server.web3, server.account)
    print(server.contract.functions.status(gameId).call())

    mark_finish(server, seqJoinGame)


async def runStartRecon(server, log):
    seqStartRecon = log['args']['seqStartRecon']
    gameId = log['args']['gameId']

    readKeys = [f'gameBoard_{gameId}']
    writeKeys = []
    readKeys =  [k.lower() for k in readKeys]
    writeKeys =  [k.lower() for k in writeKeys]

    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0

    tasks = []
    port = mpcPort(seqStartRecon, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    await asyncio.wait(tasks)

    for key in readKeys:
        server.dbLockCnt[key] += 1

    value1 = read_db(server, f'gameBoard_{gameId}', 0)
    value1 = bytes_to_dict(value1)
    game = value1

    value1 = game['value1']
    value2 = game['value2']

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(
            int_to_hex(value1)
            + int_to_hex(value2)
        )

    await run_online(server.serverID, port, server.players, server.threshold, 'rockPaperScissorsStartRecon1', seqStartRecon)

    input_arg_num = 2
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        result = hex_to_int(f.read(sz))

    if result > 2:
        result -= prime
    print('****', result)
    if result == 0:
        print('**** tie')
        winner = 'tie'
    elif result == 1 or result == -2:
        print('**** winner-player1')
        winner = 'player1'
    else:
        print('**** winner-player2')
        winner = 'player2'


    tx = server.contract.functions.winnersSet(winner, gameId).buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})
    sign_and_send(tx, server.web3, server.account)
    print(server.contract.functions.winners(gameId).call())

    mark_finish(server, seqStartRecon)


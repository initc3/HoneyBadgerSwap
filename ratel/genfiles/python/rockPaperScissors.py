import asyncio
from zkrp_pyo3 import pedersen_aggregate, pedersen_commit, zkrp_verify
from ratel.src.python.utils import (
    location_sharefile,
    prog,
    mpcPort,
    prime,
    sz,
    int_to_hex,
    hex_to_int,
    recover_input,
    fp,
    replay,
    mark_finish,
    read_db,
    write_db,
    bytes_to_int,
    bytes_to_list,
    bytes_to_dict,
    int_to_bytes,
    list_to_bytes,
    dict_to_bytes,
    execute_cmd,
    players,
)


async def monitor(server):
    blkNum = server.web3.eth.get_block_number()
    while True:
        curBlkNum = server.web3.eth.get_block_number()
        if curBlkNum - blkNum > server.confirmation:
            logs = []

            eventFilter = server.contract.events.CreateGame.createFilter(
                fromBlock=blkNum, toBlock=curBlkNum - server.confirmation
            )
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append(
                    (log["blockNumber"], log["transactionIndex"], "CreateGame", log)
                )
            eventFilter = server.contract.events.JoinGame.createFilter(
                fromBlock=blkNum, toBlock=curBlkNum - server.confirmation
            )
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append(
                    (log["blockNumber"], log["transactionIndex"], "JoinGame", log)
                )
            eventFilter = server.contract.events.StartRecon.createFilter(
                fromBlock=blkNum, toBlock=curBlkNum - server.confirmation
            )
            _logs = eventFilter.get_all_entries()
            for log in _logs:
                logs.append(
                    (log["blockNumber"], log["transactionIndex"], "StartRecon", log)
                )

            logs.sort(key=lambda s: (s[0], s[1]))
            for i in range(replay):
                for log in logs:
                    if i == 0 or log[2] == "Trade":
                        server.loop.create_task(eval(f"run{log[2]}")(server, log[3]))
            blkNum = curBlkNum - server.confirmation + 1
        else:
            await asyncio.sleep(1)


async def runCreateGame(server, log):
    seqCreateGame = log["args"]["seqCreateGame"]
    gameId = log["args"]["gameId"]
    player1 = log["args"]["player1"]
    idxValue1 = log["args"]["idxValue1"]
    maskedValue1 = log["args"]["maskedValue1"]
    idxBlinding = log["args"]["idxBlinding"]
    maskedBlinding = log["args"]["maskedBlinding"]
    proof = log["args"]["proof"]
    commitment = log["args"]["commitment"]

    # TODO:
    # proof, commitment, blinding_ = zkrp_prove(2022, 32)
    assert zkrp_verify(
        proof, commitment
    ), "[Error]: Committed secret value does not pass range proof verification!"

    value1 = recover_input(server.db, maskedValue1, idxValue1)
    blinding = recover_input(server.db, maskedBlinding, idxBlinding)

    # TODO: where is the blinding mask created? we also need to share it.
    value1_bytes = list(value1.to_bytes(32, byteorder="little"))
    blinding_bytes = list(blinding.to_bytes(32, byteorder="little"))

    share_commitment = pedersen_commit(value1_bytes, blinding_bytes)

    # TODO: create the function to commit to the unmasked secret shares.
    # TODO: we also need to change the current zkrp interface to allow specifying r and choose range to prove.

    server.zkrpShares[f"{idxValue1}"] = share_commitment
    results = await server.get_zkrp_shares(players(server.contract), f"{idxValue1}")
    # print(")))))))", results)
    agg_commitment = pedersen_aggregate(
        results, [x + 1 for x in list(range(server.players))]
    )

    # print("((((((((", agg_commitment, commitment)
    assert agg_commitment == commitment

    readKeys = []
    writeKeys = [f"gameBoard_{gameId}"]

    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0

    tasks = []
    for key in writeKeys:
        if key not in readKeys:
            tasks.append(server.dbLock[key].acquire())
    port = mpcPort(seqCreateGame, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    await asyncio.wait(tasks)

    for key in writeKeys:
        server.dbLockCnt[key] += 1

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(int_to_hex(value1))

    cmd = f"{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {prime} rockPaperScissorsCreateGame1"
    await execute_cmd(cmd)

    input_arg_num = 1
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        valid = hex_to_int(f.read(sz))

    print("**** valid", valid)
    if valid == 1:
        game = {
            "player1": player1,
            "value1": value1,
        }
        print("**** game", game)
        game = dict_to_bytes(game)
        write_db(server, f"gameBoard_{gameId}", game)

        curStatus = 1

        tx = server.contract.functions.statusSet(curStatus, gameId).buildTransaction(
            {
                "from": server.account.address,
                "gas": 1000000,
                "nonce": server.web3.eth.get_transaction_count(server.account.address),
            }
        )
        signedTx = server.web3.eth.account.sign_transaction(
            tx, private_key=server.account.privateKey
        )
        server.web3.eth.send_raw_transaction(signedTx.rawTransaction)
        server.web3.eth.wait_for_transaction_receipt(signedTx.hash)
        print(server.contract.functions.status(gameId).call())

    mark_finish(server, seqCreateGame, port)


async def runJoinGame(server, log):
    seqJoinGame = log["args"]["seqJoinGame"]
    gameId = log["args"]["gameId"]
    player2 = log["args"]["player2"]
    idxValue2 = log["args"]["idxValue2"]
    maskedValue2 = log["args"]["maskedValue2"]

    value2 = recover_input(server.db, maskedValue2, idxValue2)

    readKeys = [f"gameBoard_{gameId}"]
    writeKeys = [f"gameBoard_{gameId}"]

    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0
    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0

    tasks = []
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    for key in writeKeys:
        if key not in readKeys:
            tasks.append(server.dbLock[key].acquire())
    port = mpcPort(seqJoinGame, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    await asyncio.wait(tasks)

    for key in readKeys:
        server.dbLockCnt[key] += 1
    for key in writeKeys:
        server.dbLockCnt[key] += 1

    value1 = read_db(server, f"gameBoard_{gameId}")
    value1 = bytes_to_dict(value1)
    game = value1

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(int_to_hex(value2))

    cmd = f"{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {prime} rockPaperScissorsJoinGame1"
    await execute_cmd(cmd)

    input_arg_num = 1
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        valid = hex_to_int(f.read(sz))

    print("**** valid", valid)
    if valid == 1:
        game["player2"] = player2
        game["value2"] = value2

        print("**** game", game)

        game = dict_to_bytes(game)
        write_db(server, f"gameBoard_{gameId}", game)

        curStatus = 2

        tx = server.contract.functions.statusSet(curStatus, gameId).buildTransaction(
            {
                "from": server.account.address,
                "gas": 1000000,
                "nonce": server.web3.eth.get_transaction_count(server.account.address),
            }
        )
        signedTx = server.web3.eth.account.sign_transaction(
            tx, private_key=server.account.privateKey
        )
        server.web3.eth.send_raw_transaction(signedTx.rawTransaction)
        server.web3.eth.wait_for_transaction_receipt(signedTx.hash)
        print(server.contract.functions.status(gameId).call())

    mark_finish(server, seqJoinGame, port)


async def runStartRecon(server, log):
    seqStartRecon = log["args"]["seqStartRecon"]
    gameId = log["args"]["gameId"]

    readKeys = [f"gameBoard_{gameId}"]
    writeKeys = []

    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
            server.dbLockCnt[key] = 0

    tasks = []
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    port = mpcPort(seqStartRecon, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    await asyncio.wait(tasks)

    for key in readKeys:
        server.dbLockCnt[key] += 1

    value1 = read_db(server, f"gameBoard_{gameId}")
    value1 = bytes_to_dict(value1)
    game = value1

    value1 = game["value1"]
    value2 = game["value2"]

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(int_to_hex(value1) + int_to_hex(value2))

    cmd = f"{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {prime} rockPaperScissorsStartRecon1"
    await execute_cmd(cmd)

    input_arg_num = 2
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        result = hex_to_int(f.read(sz))

    if result > 2:
        result -= prime
    print("****", result)
    if result == 0:
        print("**** tie")
        winner = "tie"
    elif result == 1 or result == -2:
        print("**** winner-player1")
        winner = "player1"
    else:
        print("**** winner-player2")
        winner = "player2"

    tx = server.contract.functions.winnersSet(winner, gameId).buildTransaction(
        {
            "from": server.account.address,
            "gas": 1000000,
            "nonce": server.web3.eth.get_transaction_count(server.account.address),
        }
    )
    signedTx = server.web3.eth.account.sign_transaction(
        tx, private_key=server.account.privateKey
    )
    server.web3.eth.send_raw_transaction(signedTx.rawTransaction)
    server.web3.eth.wait_for_transaction_receipt(signedTx.hash)
    print(server.contract.functions.winners(gameId).call())

    mark_finish(server, seqStartRecon, port)

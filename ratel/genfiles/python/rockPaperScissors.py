import ast
import asyncio
import json
import time
from ratel.src.python.utils import (
    location_sharefile,
    prog,
    mpcPort,
    blsPrime,
    sz,
    int_to_hex,
    hex_to_int,
    recover_input,
    fp,
    replay,
    players,
)

# TODO: Manually added code for zkrp
# TODO: this can only be run under the virtual environment. can we move the package to the whole environment?
# from zkrp_pyo3 import zkrp_prove, zkrp_verify
from ratel.src.zkrp_pyo3.zkrp_pyo3 import (
    pedersen_aggregate,
    pedersen_commit,
    pedersen_open,
    zkrp_prove,
    zkrp_verify,
)


async def monitor(server, loop):
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
            for log in logs:
                for _ in range((1 if log[2] != "Trade" else replay)):
                    loop.create_task(eval(f"run{log[2]}")(server, log[3]))
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

    server.zkrpShares[f"{idxValue1}"].append(share_commitment)
    results = await server.get_zkrp_shares(players(server.contract), f"{idxValue1}")
    print("#####", results)

    print("((((((Server players:", server.players)
    agg_commitment = pedersen_aggregate(results, list(range(server.players)))

    # assert agg_commitment == commitment

    await server.dbLock["access"].acquire()
    readKeys = []
    writeKeys = [f"gameBoard_{gameId}"]
    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
    server.dbLock["access"].release()

    tasks = []
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    for key in writeKeys:
        tasks.append(server.dbLock[key].acquire())
    port = mpcPort(seqCreateGame, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    await asyncio.wait(tasks)

    for key in readKeys:
        server.dbLock[key].release()

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(int_to_hex(value1))

    cmd = f"{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {blsPrime} rockPaperScissorsCreateGame1"
    proc = await asyncio.create_subprocess_shell(
        cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE
    )
    stdout, stderr = await proc.communicate()
    print(f"[{cmd!r} exited with {proc.returncode}]")
    if stdout:
        print(f"[stdout]\n{stdout.decode()}")
    if stderr:
        print(f"[stderr]\n{stderr.decode()}")

    input_arg_num = 1
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        valid = hex_to_int(f.read(sz))
    server.portLock[port].release()

    print("**** valid", valid)
    if valid == 1:
        game = {
            "player1": player1,
            "value1": value1,
        }
        print("**** game", game)
        game = str(game)
        game = bytes(game, encoding="utf-8")
        server.db.Put(f"gameBoard_{gameId}".lower().encode(), game)
        server.dbLock[f"gameBoard_{gameId}"].release()
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

    await server.dbLock["execHistory"].acquire()
    try:
        execHistory = server.db.Get(f"execHistory".encode())
    except KeyError:
        execHistory = bytes(0)
    try:
        execHistory = execHistory.decode(encoding="utf-8")
        execHistory = dict(ast.literal_eval(execHistory))
    except:
        execHistory = {}
    execHistory[f"seqCreateGame"] = True
    execHistory = str(execHistory)
    execHistory = bytes(execHistory, encoding="utf-8")
    server.db.Put(f"execHistory".encode(), execHistory)
    server.dbLock["execHistory"].release()


async def runJoinGame(server, log):
    seqJoinGame = log["args"]["seqJoinGame"]
    gameId = log["args"]["gameId"]
    player2 = log["args"]["player2"]
    idxValue2 = log["args"]["idxValue2"]
    maskedValue2 = log["args"]["maskedValue2"]

    value2 = recover_input(server.db, maskedValue2, idxValue2)

    await server.dbLock["access"].acquire()
    readKeys = []
    writeKeys = [f"gameBoard_{gameId}"]
    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
    server.dbLock["access"].release()

    tasks = []
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    for key in writeKeys:
        tasks.append(server.dbLock[key].acquire())
    port = mpcPort(seqJoinGame, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    await asyncio.wait(tasks)

    try:
        value1 = server.db.Get(f"gameBoard_{gameId}".lower().encode())
    except KeyError:
        value1 = bytes(0)
    try:
        value1 = value1.decode(encoding="utf-8")
        value1 = dict(ast.literal_eval(value1))
    except:
        value1 = {}
    game = value1
    for key in readKeys:
        server.dbLock[key].release()

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(int_to_hex(value2))

    cmd = f"{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {blsPrime} rockPaperScissorsJoinGame1"
    proc = await asyncio.create_subprocess_shell(
        cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE
    )
    stdout, stderr = await proc.communicate()
    print(f"[{cmd!r} exited with {proc.returncode}]")
    if stdout:
        print(f"[stdout]\n{stdout.decode()}")
    if stderr:
        print(f"[stderr]\n{stderr.decode()}")

    input_arg_num = 1
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        valid = hex_to_int(f.read(sz))
    server.portLock[port].release()

    print("**** valid", valid)
    if valid == 1:
        game["player2"] = player2
        game["value2"] = value2
        print("**** game", game)
        game = str(game)
        game = bytes(game, encoding="utf-8")
        server.db.Put(f"gameBoard_{gameId}".lower().encode(), game)
        server.dbLock[f"gameBoard_{gameId}"].release()
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

    await server.dbLock["execHistory"].acquire()
    try:
        execHistory = server.db.Get(f"execHistory".encode())
    except KeyError:
        execHistory = bytes(0)
    try:
        execHistory = execHistory.decode(encoding="utf-8")
        execHistory = dict(ast.literal_eval(execHistory))
    except:
        execHistory = {}
    execHistory[f"seqJoinGame"] = True
    execHistory = str(execHistory)
    execHistory = bytes(execHistory, encoding="utf-8")
    server.db.Put(f"execHistory".encode(), execHistory)
    server.dbLock["execHistory"].release()


async def runStartRecon(server, log):
    seqStartRecon = log["args"]["seqStartRecon"]
    gameId = log["args"]["gameId"]

    await server.dbLock["access"].acquire()
    readKeys = [f"gameBoard_{gameId}"]
    writeKeys = []
    for key in readKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
    for key in writeKeys:
        if key not in server.dbLock.keys():
            server.dbLock[key] = asyncio.Lock()
    server.dbLock["access"].release()

    tasks = []
    for key in readKeys:
        tasks.append(server.dbLock[key].acquire())
    for key in writeKeys:
        tasks.append(server.dbLock[key].acquire())
    port = mpcPort(seqStartRecon, server.concurrency)
    tasks.append(server.portLock[port].acquire())
    await asyncio.wait(tasks)

    try:
        value1 = server.db.Get(f"gameBoard_{gameId}".lower().encode())
    except KeyError:
        value1 = bytes(0)
    try:
        value1 = value1.decode(encoding="utf-8")
        value1 = dict(ast.literal_eval(value1))
    except:
        value1 = {}
    game = value1
    value1 = game["value1"]
    value2 = game["value2"]
    for key in readKeys:
        server.dbLock[key].release()

    file = location_sharefile(server.serverID, port)
    with open(file, "wb") as f:
        f.write(int_to_hex(value1) + int_to_hex(value2))

    cmd = f"{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {blsPrime} rockPaperScissorsStartRecon1"
    proc = await asyncio.create_subprocess_shell(
        cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE
    )
    stdout, stderr = await proc.communicate()
    print(f"[{cmd!r} exited with {proc.returncode}]")
    if stdout:
        print(f"[stdout]\n{stdout.decode()}")
    if stderr:
        print(f"[stderr]\n{stderr.decode()}")

    input_arg_num = 2
    with open(file, "rb") as f:
        f.seek(input_arg_num * sz)
        result = hex_to_int(f.read(sz))
    server.portLock[port].release()

    if result > 2:
        result -= blsPrime
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

    await server.dbLock["execHistory"].acquire()
    try:
        execHistory = server.db.Get(f"execHistory".encode())
    except KeyError:
        execHistory = bytes(0)
    try:
        execHistory = execHistory.decode(encoding="utf-8")
        execHistory = dict(ast.literal_eval(execHistory))
    except:
        execHistory = {}
    execHistory[f"seqStartRecon"] = True
    execHistory = str(execHistory)
    execHistory = bytes(execHistory, encoding="utf-8")
    server.db.Put(f"execHistory".encode(), execHistory)
    server.dbLock["execHistory"].release()

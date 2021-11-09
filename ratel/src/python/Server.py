import aiohttp_cors
import ast
import asyncio
import os
import re
import subprocess

from aiohttp import web
from ratel.src.python.Client import send_request, reserveInput
from ratel.src.python.utils import key_inputmask, spareShares, players, threshold, batchShares, blsPrime, \
    location_inputmask, http_host, http_port, reconstruct, mpc_port, location_db, openDB, getAccount, \
    confirmation

class Server:
    def __init__(self, serverID, web3, contract, init_players, init_threshold, concurrency):
        self.serverID = serverID
        self.db = openDB(location_db(serverID))
        self.host = http_host
        self.http_port = http_port + serverID
        self.contract = contract
        self.web3 = web3
        self.account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')
        self.confirmation = confirmation
        self.players = init_players
        self.threshold = init_threshold
        self.concurrency = concurrency

        self.totInputMask = 0 #self.contract.functions.inputMaskCnt().call()

        self.portLock = {}
        for i in range(concurrency):
            self.portLock[mpc_port + i * 100] = asyncio.Lock()

        self.dbLock  = {}
        self.dbLock['access'] = asyncio.Lock()
        self.dbLock['execHistory'] = asyncio.Lock()

    async def http_server(self):
        async def handler_inputmask(request):
            print(f"s{self.serverID} request: {request}")
            mask_idxes = re.split(',', request.match_info.get("mask_idxes"))
            res = ''
            for mask_idx in mask_idxes:
                res += f"{',' if len(res) > 0 else ''}{int.from_bytes(bytes(self.db.Get(key_inputmask(mask_idx))), 'big')}"
            data = {
                "inputmask_shares": res,
            }
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        async def handler_recover_db(request):
            print(f"s{self.serverID} request: {request}")
            keys = re.split(',', request.match_info.get("keys"))
            mask_idx = int(keys[0])
            keys = keys[1:]
            num = len(keys)
            print('mask_idx, num', mask_idx, num)
            print('keys', keys)
            res = ''
            for key in keys:
                try:
                    secret = int.from_bytes(bytes(self.db.Get(key.encode())), 'big')
                    input_mask = int.from_bytes(bytes(self.db.Get(key_inputmask(mask_idx))), 'big')
                    masked_secret = (secret + input_mask) % blsPrime
                    print('masked_secret', masked_secret)
                    res += f"{',' if len(res) > 0 else ''}{masked_secret}"
                except KeyError:
                    res += f"{',' if len(res) > 0 else ''}"
                mask_idx += 1
            data = {
                "values": res,
            }
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        app = web.Application()

        cors = aiohttp_cors.setup(app, defaults={
            "*": aiohttp_cors.ResourceOptions(
                allow_credentials=True,
                expose_headers="*",
                allow_headers="*",
            )
        })

        resource = cors.add(app.router.add_resource("/inputmasks/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_inputmask))
        resource = cors.add(app.router.add_resource("/recoverdb/{keys}"))
        cors.add(resource.add_route("GET", handler_recover_db))

        print('Starting http server...')
        runner = web.AppRunner(app)
        await runner.setup()
        site = web.TCPSite(runner, host=self.host, port=self.http_port)
        await site.start()
        await asyncio.sleep(100 * 3600)

    async def init(self, recover, apptask):
        async def prepare(recover, apptask):
            isServer = self.contract.functions.isServer(self.account.address).call()
            if not isServer:
                self.registerServer()
                await self.recoverHistory(recover)

            tasks = [
                self.preprocessing(),
                self.monitorNewServer(),
                self.http_server(),
                apptask,
            ]
            await asyncio.gather(*tasks)

        tasks = [
            prepare(recover, apptask),
            self.monitorGenInputMask(),
        ]
        await asyncio.gather(*tasks)

    def genInputMask(self):
        print('Generating new inputmasks...')

        env = os.environ.copy()
        cmd = ['./random-shamir.x', '-i', f'{self.serverID}', '-N', f'{players(self.contract)}', '-T', f'{threshold(self.contract)}', '--nshares', f'{batchShares}']
        task = subprocess.Popen(cmd, env=env)
        task.wait()

        file = location_inputmask(self.serverID)
        with open(file, 'r') as f:
            for line in f.readlines():
                key = key_inputmask(self.totInputMask)
                share = int(line) % blsPrime
                self.db.Put(key, share.to_bytes((share.bit_length() + 7) // 8, 'big'))
                self.totInputMask += 1

        print(f'Total inputmask number: {self.totInputMask}\n')

    async def preprocessing(self):
        while True:
            if self.contract.functions.isInputMaskReady().call() > self.contract.functions.T().call() and self.contract.functions.isServer(self.account.address).call():
                cnt = self.contract.functions.inputMaskCnt().call()
                if cnt + spareShares >= self.totInputMask:
                    self.genInputMask()
            await asyncio.sleep(60)

    async def monitorGenInputMask(self):
        blkNum = self.web3.eth.get_block_number()
        while True:
            await asyncio.sleep(5)
            curBlkNum = self.web3.eth.get_block_number()
            if curBlkNum - blkNum > self.confirmation:
                eventFilter = self.contract.events.GenInputMask.createFilter(fromBlock=blkNum, toBlock=curBlkNum - self.confirmation)
                logs = eventFilter.get_all_entries()
                blkNum = curBlkNum - self.confirmation + 1
                for log in logs:
                    inputMaskCnt = log['args']['inputMaskCnt']
                    committeeChangeCnt = log['args']['committeeChangeCnt']

                    self.totInputMask = inputMaskCnt
                    self.genInputMask()

                    tx = self.contract.functions.setReady(committeeChangeCnt).buildTransaction({'from': self.account.address, 'gas': 1000000, 'nonce': self.web3.eth.get_transaction_count(self.account.address)})
                    signedTx = self.web3.eth.account.sign_transaction(tx, private_key=self.account.privateKey)
                    self.web3.eth.send_raw_transaction(signedTx.rawTransaction)
                    self.web3.eth.wait_for_transaction_receipt(signedTx.hash)
                    print('!!!! isInputMaskReady', self.contract.functions.isInputMaskReady().call())

    async def monitorNewServer(self):
        blkNum = self.web3.eth.get_block_number()
        while True:
            await asyncio.sleep(5)
            curBlkNum = self.web3.eth.get_block_number()
            if curBlkNum - blkNum > self.confirmation:
                eventFilter = self.contract.events.NewServer.createFilter(fromBlock=blkNum, toBlock=curBlkNum - self.confirmation)
                logs = eventFilter.get_all_entries()
                blkNum = curBlkNum - self.confirmation + 1
                for log in logs:
                    newServer = log['args']['server']

                    self.players += 1

                    tx = self.contract.functions.addServer(newServer).buildTransaction({'from': self.account.address, 'gas': 1000000, 'nonce': self.web3.eth.get_transaction_count(self.account.address)})
                    signedTx = self.web3.eth.account.sign_transaction(tx, private_key=self.account.privateKey)
                    self.web3.eth.send_raw_transaction(signedTx.rawTransaction)
                    self.web3.eth.wait_for_transaction_receipt(signedTx.hash)
                    print('!!!! votes', self.contract.functions.votes(newServer).call())

    def registerServer(self):
        tx = self.contract.functions.registerServer().buildTransaction({'from': self.account.address, 'gas': 1000000, 'nonce': self.web3.eth.get_transaction_count(self.account.address)})
        signedTx = self.web3.eth.account.sign_transaction(tx, private_key=self.account.privateKey)
        self.web3.eth.send_raw_transaction(signedTx.rawTransaction)
        self.web3.eth.wait_for_transaction_receipt(signedTx.hash)

    async def recoverHistory(self, recover):
        while True:
            isServer = self.contract.functions.isServer(self.account.address).call()
            print('isServer', isServer)
            if isServer:
                break
            await asyncio.sleep(1)
        while True:
            committeeChangeCnt = self.contract.functions.committeeChangeCnt().call()
            numCommittee = self.contract.functions.numCommittee(self.account.address).call()
            print('committeeChangeCnt', committeeChangeCnt, 'numCommittee', numCommittee)
            if numCommittee == committeeChangeCnt:
                break
            await asyncio.sleep(1)

        # check which op is missing and collect missing keys/values
        opCnt = self.contract.functions.opCnt().call()
        print('!!!! opCnt', opCnt)
        try:
            execHistory = self.db.Get(f'execHistory'.encode())
        except KeyError:
            execHistory = bytes(0)

        try:
            execHistory = execHistory.decode(encoding='utf-8')
            execHistory = dict(ast.literal_eval(execHistory))
        except:
            execHistory = {}

        request_keys = {}
        for i in range(opCnt):
            print('missing opSeq', i)
            if not i in execHistory:
                keys = recover(self.contract, i)
                for key in keys:
                    request_keys[key] = True
        if len(request_keys):
            if 'execHistory' in request_keys.keys():
                del request_keys['execHistory']
            print('request_keys', [*request_keys])
            mask_idxes = reserveInput(self.web3, self.contract, len([*request_keys]), self.account)
            print('mask_idxes', mask_idxes)
            keys = str(mask_idxes[0])
            for key in [*request_keys]:
                keys += f',{key.lower()}'

            # fetch share from other servers
            shares = []
            for serverID in range(players(self.contract)):
                if serverID != self.serverID:
                    url = f"http://{http_host}:{http_port + serverID}/recoverdb/{keys}"
                    result = await send_request(url)
                    for i, share in enumerate(re.split(",", result["values"])):
                        if len(share) <= 0:
                            continue
                        if (len(shares) <= i):
                            shares.append([])
                        shares[i].append(int(share))
            print('shares', shares)
            mask_idx = mask_idxes[0]
            for key, _shares in zip(keys, shares):
                masked_value = reconstruct(_shares, len(_shares))
                input_mask = int.from_bytes(bytes(self.db.Get(key_inputmask(mask_idx))), 'big')
                share = (masked_value - input_mask) % blsPrime
                mask_idx += 1
                self.db.Put(key.encode(), share.to_bytes((share.bit_length() + 7) // 8, 'big'))

            # mark op as executed
            for i in range(opCnt):
                if not i in execHistory:
                    execHistory[i] = True
            execHistory = str(execHistory)
            execHistory = bytes(execHistory, encoding='utf-8')
            self.db.Put(f'execHistory'.encode(), execHistory)


import aiohttp_cors
import ast
import asyncio
import os
import re
import subprocess

from aiohttp import web
from ratel.src.python.Client import send_requests, batch_interpolate
from ratel.src.python.utils import key_inputmask, spareShares, players, threshold, blsPrime, \
    location_inputmask, http_host, http_port, mpc_port, location_db, openDB, getAccount, \
    confirmation, shareBatchSize, list_to_str


class Server:
    def __init__(self, serverID, web3, contract, init_players, init_threshold, concurrency, recover, test=False):
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

        self.recover = recover

        self.test = test

        self.input_mask_queue_tail = 0 #self.contract.functions.inputMaskCnt().call() if db has not been cleared

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
            seq_num_list = re.split(',', request.match_info.get("list"))

            keys = self.collect_keys(seq_num_list)
            masked_shares = self.mask_shares(keys)

            res = list_to_str(masked_shares)

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
        resource = cors.add(app.router.add_resource("/recoverdb/{list}"))
        cors.add(resource.add_route("GET", handler_recover_db))

        print('Starting http server...')
        runner = web.AppRunner(app)
        await runner.setup()
        site = web.TCPSite(runner, host=self.host, port=self.http_port)
        await site.start()
        await asyncio.sleep(100 * 3600)

    async def init(self, apptask):
        async def prepare(apptask):
            isServer = self.contract.functions.isServer(self.account.address).call()
            if not isServer:
                self.registerServer()
                await self.recoverHistory()

            tasks = [
                self.preprocessing(),
                self.monitorNewServer(),
                self.http_server(),
                apptask,
            ]
            await asyncio.gather(*tasks)

        tasks = [
            prepare(apptask),
            self.monitorGenInputMask(shareBatchSize),
        ]
        await asyncio.gather(*tasks)

    def genInputMask(self, shareBatchSize):
        print('Generating new inputmasks...')

        env = os.environ.copy()
        cmd = ['./random-shamir.x', '-i', f'{self.serverID}', '-N', f'{players(self.contract)}', '-T', f'{threshold(self.contract)}', '--nshares', f'{shareBatchSize}']
        task = subprocess.Popen(cmd, env=env)
        task.wait()

        file = location_inputmask(self.serverID)
        with open(file, 'r') as f:
            for line in f.readlines():
                key = key_inputmask(self.input_mask_queue_tail)
                share = int(line) % blsPrime
                self.db.Put(key, share.to_bytes((share.bit_length() + 7) // 8, 'big'))
                self.input_mask_queue_tail += 1

        print(f'Total inputmask number: {self.input_mask_queue_tail}\n')

    def check_input_mask_availability(self):
        input_mask_queue_head = self.contract.functions.inputMaskCnt().call()
        if input_mask_queue_head + spareShares >= self.input_mask_queue_tail:
            self.genInputMask(shareBatchSize)

    async def preprocessing(self):
        while True:
            if self.contract.functions.isInputMaskReady().call() > self.contract.functions.T().call() and self.contract.functions.isServer(self.account.address).call():
                self.check_input_mask_availability()
            await asyncio.sleep(60)

    async def monitorGenInputMask(self, shareBatchSize):
        blkNum = self.web3.eth.get_block_number()
        while True:
            await asyncio.sleep(5)
            curBlkNum = self.web3.eth.get_block_number()
            if curBlkNum - blkNum > self.confirmation:
                eventFilter = self.contract.events.GenInputMask.createFilter(fromBlock=blkNum, toBlock=curBlkNum - self.confirmation)
                logs = eventFilter.get_all_entries()
                blkNum = curBlkNum - self.confirmation + 1
                for log in logs:
                    input_mask_queue_head = log['args']['inputMaskCnt']
                    committeeChangeCnt = log['args']['committeeChangeCnt']

                    self.input_mask_queue_tail = input_mask_queue_head
                    self.genInputMask(shareBatchSize)

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

    async def recoverHistory(self):
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

        #TODO: test below
        seq_num_list = self.check_missing_tasks()
        request = f'recoverdb/{list_to_str(seq_num_list)}'
        masked_shares = await send_requests(players(self.contract), request)
        for i in range(len(masked_shares)):
            masked_shares[i] = re.split(",", masked_shares[i]["values"])
        keys = self.collect_keys(seq_num_list)
        masked_states = batch_interpolate(masked_shares)
        state_shares = self.recover_states(masked_states)
        self.restore_db(seq_num_list, keys, state_shares)

    def check_missing_tasks(self):
        try:
            execHistory = self.db.Get(f'execHistory'.encode())
        except KeyError:
            execHistory = bytes(0)

        try:
            execHistory = execHistory.decode(encoding='utf-8')
            execHistory = dict(ast.literal_eval(execHistory))
        except:
            execHistory = {}

        opCnt = self.contract.functions.opCnt().call()
        seq_num_list = []
        for seq in range(opCnt):
            if not seq in execHistory:
                print('missing opSeq', seq)
                seq_num_list.append(seq)
        return seq_num_list

    def collect_keys(self, seq_num_list):
        if not self.test:
            seq_num_list = list(set(seq_num_list))

        keys = []
        for seq_num in seq_num_list:
            keys.extend(self.recover(self.contract, int(seq_num)))

        if not self.test:
            keys = list(set(keys))

        return keys

    def mask_shares(self, keys):
        masked_shares = []

        for key in keys:
            masked_state_share = 0
            try:
                secret = int.from_bytes(bytes(self.db.Get(key.lower().encode())), 'big')

                self.check_input_mask_availability()
                input_mask_share = int.from_bytes(bytes(self.db.Get(key_inputmask(self.input_mask_queue_tail - 1))), 'big')
                self.input_mask_queue_tail -= 1
                masked_state_share = (secret + input_mask_share) % blsPrime

            except KeyError:
                print(f'Do not have the state {key}')

            masked_shares.append(masked_state_share)

        return masked_shares

    def recover_states(self, masked_states):
        state_shares = []

        for masked_state in masked_states:
            input_mask = int.from_bytes(bytes(self.db.Get(key_inputmask(self.input_mask_queue_tail - 1))), 'big')
            self.input_mask_queue_tail -= 1
            state_share = (masked_state - input_mask) % blsPrime
            state_shares.append(state_share)

        return state_shares

    def restore_db(self, seq_num_list, keys, values):
        assert len(keys) == len(values)

        for key, value in zip(keys, values):
            self.db.Put(key.encode(), value.to_bytes((value.bit_length() + 7) // 8, 'big'))

        try:
            execHistory = self.db.Get(f'execHistory'.encode())
        except KeyError:
            execHistory = bytes(0)

        try:
            execHistory = execHistory.decode(encoding='utf-8')
            execHistory = dict(ast.literal_eval(execHistory))
        except:
            execHistory = {}

        for seq in seq_num_list:
                execHistory[seq] = True

        execHistory = str(execHistory)
        execHistory = bytes(execHistory, encoding='utf-8')
        self.db.Put(f'execHistory'.encode(), execHistory)
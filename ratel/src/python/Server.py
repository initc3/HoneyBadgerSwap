import json

import aiohttp_cors
import ast
import asyncio
import json
import re
import time
import math

from aiohttp import web, ClientSession
from collections import defaultdict
from zkrp_pyo3 import pedersen_aggregate, pedersen_commit, zkrp_verify, zkrp_prove

from ratel.src.python.Client import send_requests, batch_interpolate
from ratel.src.python.utils import key_inputmask_index, key_serverval_index, key_zkrp_blinding_index, \
    key_zkrp_blinding_commitment_index, key_zkrp_agg_commitment_index, spareShares, prime, \
    location_inputmask, http_host, http_port, mpc_port, location_db, openDB, getAccount, \
    confirmation, shareBatchSize, list_to_str, trade_key_num, INPUTMASK_SHARES_DIR, execute_cmd, sign_and_send, \
    key_inputmask_version, list_to_bytes, bytes_to_list


class Server:
    def __init__(self, serverID, web3, contract, init_players, init_threshold, concurrency, recover):#, test=False):
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

        self.portLock = {}
        for i in range(-1, concurrency):
            self.portLock[mpc_port + i * 100] = asyncio.Lock()

        self.dbLock = {}
        self.dbLockCnt = {}

        self.loop = asyncio.get_event_loop()

        self.local_input_mask_cnt = 0
        self.local_zkrp_blinding_share_cnt = 0
        self.local_zkrp_blinding_com_cnt = 0
        self.used_zkrp_blinding_share = 0
        self.used_zkrp_blinding_com = 0

        self.zkrp_blinding_commitment = []

        # self.test = test
        #
        # self.input_mask_queue_tail = 0
        # try:
        #     self.input_mask_queue_tail = int.from_bytes(bytes(self.db.Get(f'input_mask_queue_tail'.encode())), 'big')
        # except KeyError:
        #     pass
        # print('**** input_mask_queue_tail', self.input_mask_queue_tail)

        self.zkrpShares = {}


    async def get_zkrp_shares(self, players, inputmask_idxes):
        request = f"zkrp_share_idxes/{inputmask_idxes}"
        results = await send_requests(players, request)
        parsed_results = []
        for i in range(len(results)):
            parsed_results.append(json.loads(results[i]["zkrp_share_idx"]))

        return parsed_results

    async def http_server(self):
        async def handler_inputmask(request):
            print(f"s{self.serverID} request: {request}")
            mask_idxes = re.split(",", request.match_info.get("mask_idxes"))
            res = ""
            for mask_idx in mask_idxes:
                res += f"{',' if len(res) > 0 else ''}{int.from_bytes(bytes(self.db.Get(key_inputmask_index(mask_idx))), 'big')}"
            data = {
                "inputmask_shares": res,
            }
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        # async def handler_recover_db(request):
        #     print(f"s{self.serverID} request: {request}")
        #     seq_num_list = re.split(',', request.match_info.get("list"))
        #
        #     with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        #         f.write(f'state\t{len(seq_num_list * trade_key_num)}\t'
        #                 f'stage\t3\t'
        #                 f'{time.perf_counter()}\t'
        #                 f's-{self.serverID}\n')
        #
        #     keys = self.collect_keys(seq_num_list)
        #     masked_shares = await self.mask_shares(keys)
        #
        #     with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
        #         f.write(f'state\t{len(seq_num_list * trade_key_num)}\t'
        #                 f'stage\t6\t'
        #                 f'{time.perf_counter()}\t'
        #                 f's-{self.serverID}\n')
        #
        #     res = list_to_str(masked_shares)
        #
        #     data = {
        #         "values": res,
        #     }
        #     print(f"s{self.serverID} response: {res}")
        #     return web.json_response(data)

        async def handler_mpc_verify(request):
            print(f"s{self.serverID} request: s{request} request from {request.remote}")
            mask_idx = re.split(',', request.match_info.get("mask_idxes"))[0]

            while mask_idx not in self.zkrpShares.keys():
                await asyncio.sleep(1)

            data = {
                "zkrp_share_idx": json.dumps(self.zkrpShares[mask_idx]),
            }
            return web.json_response(data)

        async def handler_serverval(request):
            print(f"s{self.serverID} request: {request}")
            mask_idxes = re.split(",", request.match_info.get("mask_idxes"))
            print("mask_idxes:",mask_idxes)

            res = ""
            for mask_idx in mask_idxes:
                t1 = key_serverval_index(mask_idx)
                print('t1:',t1)
                int.from_bytes(bytes(self.db.Get(t1)), 'big')
                res += f"{',' if len(res) > 0 else ''}{int.from_bytes(bytes(self.db.Get(key_serverval_index(mask_idx))), 'big')}"
            data = {
                "serverval_shares": res,
            }
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        async def handler_zkrp_blinding_shares(request):
            print(f"s{self.serverID} request: {request}")
            need_num = int(re.split(",", request.match_info.get("mask_idxes"))[0])
            while self.used_zkrp_blinding_share + need_num > self.local_zkrp_blinding_share_cnt:
                await asyncio.sleep(1)
                await self.gen_zkrp_blinding_shares(100)

            res = ""
            for i in range(need_num):
                real_idx = self.used_zkrp_blinding_share + i
                res += f"{',' if len(res) > 0 else ''}{int.from_bytes(bytes(self.db.Get(key_zkrp_blinding_index(real_idx))), 'big')}"
            data = {
                "zkrp_blinding_shares": res,
            }
            self.used_zkrp_blinding_share += need_num
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        async def handler_zkrp_blinding_commitment_shares(request):
            print(f"s{self.serverID} request: {request}")
            mask_idxes = re.split(",", request.match_info.get("mask_idxes"))

            res = ""
            for mask_idx in mask_idxes:
                cur_lis = json.loads(self.db.Get(key_zkrp_blinding_commitment_index(mask_idx)).decode())
                # tmp_str = f'{cur_lis}'
                res += f"{',' if len(res) > 0 else ''}{cur_lis}"

            data = {
                "zkrp_blinding_commitment_shares": res,
            }
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        # async def handler_zkrp_blinding_info_1(request):
            # print(f"s{self.serverID} request: {request}")
            # need_num = int(re.split(",", request.match_info.get("mask_idxes"))[0])
            # while self.used_zkrp_blinding_share + need_num > self.local_zkrp_blinding_share_cnt:
                # await asyncio.sleep(1)
                # await self.gen_zkrp_blinding_shares(100)

            # res = ""
            # for i in range(need_num):
                # real_idx = self.used_zkrp_blinding_share + i
                # tmp_str = json.loads(self.db.Get(key_zkrp_blinding_commitment_index(real_idx)).decode())
                # res += f"{';' if len(res) > 0 else ''}{tmp_str}"

            # data = {
                # "zkrp_blinding_info_1": res,
            # }
            # print(f"s{self.serverID} response: {res}")
            # return web.json_response(data)

        async def handler_zkrp_blinding_info_2(request):
            print(f"s{self.serverID} request: {request}")
            need_num = int(re.split(",", request.match_info.get("mask_idxes"))[0])
            while self.used_zkrp_blinding_com + need_num > self.local_zkrp_blinding_com_cnt:
                await asyncio.sleep(1)
            
            print('used num:',self.used_zkrp_blinding_com)
            print('need num:',need_num)
            print('local blinding:',self.local_zkrp_blinding_com_cnt)

            res = ""
            for i in range(need_num):
                real_idx = self.used_zkrp_blinding_com + i
                tmp_str = json.loads(self.db.Get(key_zkrp_agg_commitment_index(real_idx)).decode())
                res += f"{';' if len(res) > 0 else ''}{tmp_str}"

            self.used_zkrp_blinding_com += need_num
            data = {
                "zkrp_blinding_info_2": res,
            }
            print(f"s{self.serverID} response: {res}")
            return web.json_response(data)

        app = web.Application()

        cors = aiohttp_cors.setup(
            app,
            defaults={
                "*": aiohttp_cors.ResourceOptions(
                    allow_credentials=True,
                    expose_headers="*",
                    allow_headers="*",
                )
            },
        )

        resource = cors.add(app.router.add_resource("/inputmasks/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_inputmask))
        # resource = cors.add(app.router.add_resource("/recoverdb/{list}"))
        # cors.add(resource.add_route("GET", handler_recover_db))
        resource = cors.add(app.router.add_resource("/zkrp_share_idxes/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_mpc_verify))
        resource = cors.add(app.router.add_resource("/serverval/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_serverval))
        resource = cors.add(app.router.add_resource("/zkrp_blinding_shares/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_zkrp_blinding_shares))
        resource = cors.add(app.router.add_resource("/zkrp_blinding_commitment_shares/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_zkrp_blinding_commitment_shares))
        # resource = cors.add(app.router.add_resource("/zkrp_new_blinding_shares/{mask_idxes}"))
        # cors.add(resource.add_route("GET", handler_zkrp_blinding_info_1))
        resource = cors.add(app.router.add_resource("/zkrp_new_agg_com/{mask_idxes}"))
        cors.add(resource.add_route("GET", handler_zkrp_blinding_info_2))

        print("Starting http server...")
        runner = web.AppRunner(app)
        await runner.setup()
        site = web.TCPSite(runner, host=self.host, port=self.http_port)
        await site.start()
        await asyncio.sleep(100 * 3600)


    #TODO: modify the following
    # async def init(self, apptask):
    #     async def prepare(apptask):
    #         # isServer = self.contract.functions.isServer(self.account.address).call()
    #         # if not isServer:
    #         #     self.registerServer()
    #         #     await self.recoverHistory()
    #
    #         tasks = [
    #             self.preprocessing(),
    #             # self.monitorNewServer(),
    #             self.http_server(),
    #             apptask,
    #         ]
    #         await asyncio.gather(*tasks)
    #
    #     tasks = [
    #         prepare(apptask),
    #         self.monitorGenInputMask(shareBatchSize),
    #     ]
    #     await asyncio.gather(*tasks)

    async def init(self, monitor):
        tasks = [
            monitor,
            self.http_server(),
            # self.preprocess_zkrp_blinding(),
            self.preprocessing()
        ]
        await asyncio.gather(*tasks)

    async def gen_zkrp_blinding_shares(self, share_batch_size=shareBatchSize):
        print(f'Generating blinding shares... s-{self.serverID}')

        cmd = f'./random-shamir.x -i {self.serverID} -N {self.players} -T {self.threshold} --nshares {share_batch_size} --prep-dir {INPUTMASK_SHARES_DIR} -P {prime}'
        await execute_cmd(cmd)

        cur_zkrp_blinding_cnt = self.local_zkrp_blinding_share_cnt
        file = location_inputmask(self.serverID, self.players)
        with open(file, "r") as f:
            i = 0
            for line in f.readlines():
                if i % 2 == 0:
                    share = int(line) % prime
                    self.db.Put(key_zkrp_blinding_index(cur_zkrp_blinding_cnt,0), share.to_bytes((share.bit_length() + 7) // 8, 'big'))
                else:
                    share_prime = int(line) % prime
                    self.db.Put(key_zkrp_blinding_index(cur_zkrp_blinding_cnt,1), share_prime.to_bytes((share_prime.bit_length() + 7) // 8, 'big'))

                    value_bytes = list(share.to_bytes(32, byteorder='little'))
                    blinding_bytes = list(share_prime.to_bytes(32, byteorder='little'))
                    share_commitment = pedersen_commit(value_bytes,blinding_bytes)
                    self.db.Put(key_zkrp_blinding_commitment_index(cur_zkrp_blinding_cnt), json.dumps(share_commitment).encode())

                    cur_zkrp_blinding_cnt += 1
                i = i + 1

        self.local_zkrp_blinding_share_cnt = cur_zkrp_blinding_cnt
        print(f'Total zkrp blinding shares number: {self.local_zkrp_blinding_share_cnt}\n')

    async def gen_input_mask(self, input_mask_cnt, input_mask_version, share_batch_size=shareBatchSize):
        print(f'Generating new inputmasks... s-{self.serverID}')

        cmd = f'./random-shamir.x -i {self.serverID} -N {self.players} -T {self.threshold} --nshares {share_batch_size} --prep-dir {INPUTMASK_SHARES_DIR} -P {prime}'
        await execute_cmd(cmd)

        file = location_inputmask(self.serverID, self.players)
        with open(file, "r") as f:
            for line in f.readlines():
                share = int(line) % prime
                self.db.Put(key_inputmask_index(input_mask_cnt), share.to_bytes((share.bit_length() + 7) // 8, 'big'))
                self.db.Put(key_inputmask_version(input_mask_cnt), input_mask_version.to_bytes((input_mask_version.bit_length() + 7) // 8, 'big'))
                input_mask_cnt += 1

        # self.db.Put(f'input_mask_queue_tail'.encode(), self.input_mask_queue_tail.to_bytes((self.input_mask_queue_tail.bit_length() + 7) // 8, 'big'))

        self.local_input_mask_cnt = input_mask_cnt
        print(f'Total inputmask number: {self.local_input_mask_cnt}\n')

    # async def check_input_mask_availability(self):
    #     input_mask_queue_head = self.contract.functions.inputMaskCnt().call()
    #     if input_mask_queue_head + spareShares >= self.input_mask_queue_tail:
    #         await self.genInputMask(shareBatchSize)


    async def gen_batch_zkrp_blinding(self):
        origin_cnt = self.local_zkrp_blinding_share_cnt

        ##### (1) generating the zkrp blinding shares #####
        await self.gen_zkrp_blinding_shares(100)
        while origin_cnt == self.local_zkrp_blinding_share_cnt:
            await asyncio.sleep(1)
            print('waiting for gen zkrp blind share')

        ##### (2) interpolate the blinding commitment #####
        cur_zkrp_com = self.local_zkrp_blinding_com_cnt
        for cur_zkrp_idx in range(50):
            request = f"zkrp_blinding_commitment_shares/{cur_zkrp_com}"
            results = await send_requests(self.players, request)
            for i in range(len(results)):
                tmp_str = results[i]["zkrp_blinding_commitment_shares"]
                results[i] = re.split(',', tmp_str[1:-1])
            for i in range(len(results)):
                for j in range(len(results[i])):
                    results[i][j] = int(results[i][j])

            agg_commitment = pedersen_aggregate(results, [x + 1 for x in list(range(self.players))])
            self.db.Put(key_zkrp_agg_commitment_index(cur_zkrp_com), json.dumps(agg_commitment).encode())
            cur_zkrp_com = cur_zkrp_com + 1
        self.local_zkrp_blinding_com_cnt = cur_zkrp_com
        print('zkrp blinding generated!')


    async def preprocess_zkrp_blinding(self):
        await self.gen_batch_zkrp_blinding()

    async def preprocessing(self):
        ### TODO: remove the following
        if (self.serverID != 0):
            return

        while True:
            input_mask_cnt = self.contract.functions.inputMaskCnt().call()
            if input_mask_cnt + spareShares >= self.local_input_mask_cnt:
                print(f'Request to generate input masks....')
                tx = self.contract.functions.genInputMask(self.local_input_mask_cnt).buildTransaction(
                    {'from': self.account.address, 'gas': 1000000,
                     'nonce': self.web3.eth.get_transaction_count(self.account.address)})
                sign_and_send(tx, self.web3, self.account)
            await asyncio.sleep(600)

        # while True:
        #     if self.contract.functions.isInputMaskReady().call() > self.contract.functions.T().call() and self.contract.functions.isServer(self.account.address).call():
        #         await self.check_input_mask_availability()
        #     await asyncio.sleep(60)


    # async def monitorGenInputMask(self, shareBatchSize):
    #     blkNum = self.web3.eth.get_block_number()
    #     while True:
    #         await asyncio.sleep(5)
    #         curBlkNum = self.web3.eth.get_block_number()
    #         if curBlkNum - blkNum > self.confirmation:
    #             eventFilter = self.contract.events.GenInputMask.createFilter(fromBlock=blkNum, toBlock=curBlkNum - self.confirmation)
    #             logs = eventFilter.get_all_entries()
    #             blkNum = curBlkNum - self.confirmation + 1
    #             for log in logs:
    #                 input_mask_queue_head = log['args']['inputMaskCnt']
    #                 committeeChangeCnt = log['args']['committeeChangeCnt']
    #
    #                 self.input_mask_queue_tail = input_mask_queue_head
    #                 await self.genInputMask(shareBatchSize)
    #
    #                 tx = self.contract.functions.setReady(committeeChangeCnt).buildTransaction({'from': self.account.address, 'gas': 1000000, 'nonce': self.web3.eth.get_transaction_count(self.account.address)})
    #                 sign_and_send(tx, self.web3, self.account)
    #                 print('!!!! isInputMaskReady', self.contract.functions.isInputMaskReady().call())
    #
    #
    # async def monitorNewServer(self):
    #     blkNum = self.web3.eth.get_block_number()
    #     while True:
    #         await asyncio.sleep(5)
    #         curBlkNum = self.web3.eth.get_block_number()
    #         if curBlkNum - blkNum > self.confirmation:
    #             eventFilter = self.contract.events.NewServer.createFilter(fromBlock=blkNum, toBlock=curBlkNum - self.confirmation)
    #             logs = eventFilter.get_all_entries()
    #             blkNum = curBlkNum - self.confirmation + 1
    #             for log in logs:
    #                 newServer = log['args']['server']
    #
    #                 self.players += 1
    #
    #                 tx = self.contract.functions.addServer(newServer).buildTransaction({'from': self.account.address, 'gas': 1000000, 'nonce': self.web3.eth.get_transaction_count(self.account.address)})
    #                 sign_and_send(tx, self.web3, self.account)
    #                 print('!!!! votes', self.contract.functions.votes(newServer).call())
    #
    # def registerServer(self):
    #     tx = self.contract.functions.registerServer().buildTransaction({'from': self.account.address, 'gas': 1000000, 'nonce': self.web3.eth.get_transaction_count(self.account.address)})
    #     sign_and_send(tx, self.web3, self.account)
    #
    #
    # async def recoverHistory(self):
    #     while True:
    #         isServer = self.contract.functions.isServer(self.account.address).call()
    #         print('isServer', isServer)
    #         if isServer:
    #             break
    #         await asyncio.sleep(1)
    #     while True:
    #         committeeChangeCnt = self.contract.functions.committeeChangeCnt().call()
    #         numCommittee = self.contract.functions.numCommittee(self.account.address).call()
    #         print('committeeChangeCnt', committeeChangeCnt, 'numCommittee', numCommittee)
    #         if numCommittee == committeeChangeCnt:
    #             break
    #         await asyncio.sleep(1)
    #
    #     #TODO: test below
    #     seq_num_list = self.check_missing_tasks()
    #     request = f'recoverdb/{list_to_str(seq_num_list)}'
    #     masked_shares = await send_requests(self.players, request)
    #     for i in range(len(masked_shares)):
    #         masked_shares[i] = re.split(",", masked_shares[i]["values"])
    #     keys = self.collect_keys(seq_num_list)
    #     masked_states = batch_interpolate(masked_shares)
    #     state_shares = self.recover_states(masked_states)
    #     self.restore_db(seq_num_list, keys, state_shares)
    #
    #
    # def check_missing_tasks(self):
    #     try:
    #         execHistory = self.db.Get(f'execHistory'.encode())
    #     except KeyError:
    #         execHistory = bytes(0)
    #
    #     try:
    #         execHistory = execHistory.decode(encoding='utf-8')
    #         execHistory = dict(ast.literal_eval(execHistory))
    #     except:
    #         execHistory = {}
    #
    #     opCnt = self.contract.functions.opCnt().call()
    #     seq_num_list = []
    #     for seq in range(opCnt):
    #         if not seq in execHistory:
    #             print('missing opSeq', seq)
    #             seq_num_list.append(seq)
    #     return seq_num_list
    #
    #
    # def collect_keys(self, seq_num_list):
    #     if not self.test:
    #         seq_num_list = list(set(seq_num_list))
    #
    #     keys = []
    #     for seq_num in seq_num_list:
    #         keys.extend(self.recover(self.contract, int(seq_num)))
    #
    #     if not self.test:
    #         keys = list(set(keys))
    #
    #     return keys
    #
    #
    # async def mask_shares(self, keys):
    #     masked_shares = []
    #
    #     with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
    #         f.write(f'state\t{len(keys)}\t'
    #                 f'stage\t4\t'
    #                 f'{time.perf_counter()}\t'
    #                 f's-{self.serverID}\n')
    #
    #     await self.genInputMask(len(keys))
    #
    #     with open(f'ratel/benchmark/data/recover_states.csv', 'a') as f:
    #         f.write(f'state\t{len(keys)}\t'
    #                 f'stage\t5\t'
    #                 f'{time.perf_counter()}\t'
    #                 f's-{self.serverID}\n')
    #
    #     for key in keys:
    #         masked_state_share = 0
    #         try:
    #             secret = int.from_bytes(bytes(self.db.Get(key.lower().encode())), 'big')
    #
    #             input_mask_share = int.from_bytes(bytes(self.db.Get(key_inputmask_index(self.input_mask_queue_tail - 1))), 'big')
    #             self.input_mask_queue_tail -= 1
    #             masked_state_share = (secret + input_mask_share) % prime
    #
    #         except KeyError:
    #             print(f'Do not have the state {key}')
    #
    #         masked_shares.append(masked_state_share)
    #
    #     self.db.Put(f'input_mask_queue_tail'.encode(), self.input_mask_queue_tail.to_bytes((self.input_mask_queue_tail.bit_length() + 7) // 8, 'big'))
    #
    #     return masked_shares
    #
    #
    # def recover_states(self, masked_states):
    #     state_shares = []
    #
    #     for masked_state in masked_states:
    #         input_mask = int.from_bytes(bytes(self.db.Get(key_inputmask_index(self.input_mask_queue_tail - 1))), 'big')
    #         self.input_mask_queue_tail -= 1
    #         state_share = (masked_state - input_mask) % prime
    #         state_shares.append(state_share)
    #
    #     self.db.Put(f'input_mask_queue_tail'.encode(), self.input_mask_queue_tail.to_bytes((self.input_mask_queue_tail.bit_length() + 7) // 8, 'big'))
    #
    #     return state_shares
    #
    #
    # def restore_db(self, seq_num_list, keys, values):
    #     assert len(keys) == len(values)
    #
    #     for key, value in zip(keys, values):
    #         self.db.Put(key.encode(), value.to_bytes((value.bit_length() + 7) // 8, 'big'))
    #
    #     try:
    #         execHistory = self.db.Get(f'execHistory'.encode())
    #     except KeyError:
    #         execHistory = bytes(0)
    #
    #     try:
    #         execHistory = execHistory.decode(encoding='utf-8')
    #         execHistory = dict(ast.literal_eval(execHistory))
    #     except:
    #         execHistory = {}
    #
    #     for seq in seq_num_list:
    #             execHistory[seq] = True
    #
    #     execHistory = str(execHistory)
    #     execHistory = bytes(execHistory, encoding='utf-8')
    #     self.db.Put(f'execHistory'.encode(), execHistory)

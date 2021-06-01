import asyncio
import json
import os
import subprocess
import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware

from ratel.src.python.utils import threshold, spareShares, players, batchShares, blsPrime, \
    location_inputmask, key_inputmask

url = 'ws://0.0.0.0:8546'
ETH = '0x0000000000000000000000000000000000000000'
tokenAddress = '0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2'
appAddress = '0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385'
confirmation = 2

def parse_contract(name):
    contract = json.load(open(f'ratel/genfiles/build/contracts/{name}.json'))
    return contract['abi'], contract['bytecode']

def getAccount(web3, keystoreDir):
    for filename in os.listdir(keystoreDir):
        with open(keystoreDir + filename) as keyfile:
            encryptedKey = keyfile.read()
            privateKey = web3.eth.account.decrypt(encryptedKey, '')
            return web3.eth.account.privateKeyToAccount(privateKey)

def reserveInput(web3, appContract, num, account):
    # tx_hash = appContract.functions.reserveInput(num).transact()
    # web3.eth.wait_for_transaction_receipt(tx_hash)
    tx = appContract.functions.reserveInput(num).buildTransaction({'from': account.address, 'gas': 1000000, 'nonce': web3.eth.get_transaction_count(account.address)})
    signedTx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)
    web3.eth.send_raw_transaction(signedTx.rawTransaction)
    web3.eth.wait_for_transaction_receipt(signedTx.hash)
    receipt = web3.eth.get_transaction_receipt(signedTx.hash)
    log = appContract.events.InputMask().processReceipt(receipt)
    return log[0]['args']['inpusMaskIndexes']

async def preprocessing(db, contract, serverID):
    tot = contract.functions.inputmaskCnt().call()
    while True:
        cnt = contract.functions.inputmaskCnt().call()
        if cnt + spareShares >= tot:
            print('Generating new inputmasks...')

            env = os.environ.copy()
            cmd = ['./random-shamir.x', '-i', f'{serverID}', '-N', f'{players}', '-T', f'{threshold}', '--nshares', f'{batchShares}']
            task = subprocess.Popen(cmd, env=env)
            task.wait()

            file = location_inputmask(serverID)
            with open(file, 'r') as f:
                idx = tot
                for line in f.readlines():
                    key = key_inputmask(idx)
                    share = int(line) % blsPrime
                    db.Put(key, share.to_bytes((share.bit_length() + 7) // 8, 'big'))
                    idx += 1

            tot += batchShares
            print(f'Total inputmask number: {tot}\n')

        await asyncio.sleep(60)

if __name__=='__main__':
    appName = sys.argv[1]

    web3 = Web3(Web3.WebsocketProvider(url))

    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    abi, bytecode = parse_contract('Token')
    tx_hash = web3.eth.contract(
        abi=abi,
        bytecode=bytecode).constructor().transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    tokenAddress = web3.eth.waitForTransactionReceipt(tx_hash)['contractAddress']
    print(f'Deployed to: {tokenAddress}\n')
    tokenContract = web3.eth.contract(address=tokenAddress, abi=abi)

    abi, bytecode = parse_contract(appName)
    servers = []
    for serverID in range(4):
        account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')
        servers.append(account.address)
    tx_hash = web3.eth.contract(
        abi=abi,
        bytecode=bytecode).constructor(servers, threshold).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    appAddress = web3.eth.waitForTransactionReceipt(tx_hash)['contractAddress']
    print(f'Deployed to: {appAddress}\n')
    appContract = web3.eth.contract(address=appAddress, abi=abi)





import json
import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.Server import getAccount

url = 'ws://0.0.0.0:8546'
ETH = '0x0000000000000000000000000000000000000000'
tokenAddress = '0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2'
appAddress = '0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385'

def parse_contract(name):
    contract = json.load(open(f'ratel/genfiles/build/contracts/{name}.json'))
    return contract['abi'], contract['bytecode']

if __name__=='__main__':
    appName = sys.argv[1]
    init_players = int(sys.argv[2])
    init_threshold = int(sys.argv[3])

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
    for serverID in range(init_players):
        account = getAccount(web3, f'/opt/poa/keystore/server_{serverID}/')
        servers.append(account.address)
    tx_hash = web3.eth.contract(
        abi=abi,
        bytecode=bytecode).constructor(servers, init_threshold).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)
    appAddress = web3.eth.waitForTransactionReceipt(tx_hash)['contractAddress']
    print(f'Deployed to: {appAddress}\n')
    appContract = web3.eth.contract(address=appAddress, abi=abi)





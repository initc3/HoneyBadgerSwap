import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.utils import getAccount, parse_contract

url = 'ws://0.0.0.0:8546'

app_addr = '0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2'

token_addrs = [
    '0x0000000000000000000000000000000000000000',
    '0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385', #1
    '0x8C89e5D2bCc0e4C26E3295d48d052E11bd03C06A', #2
    '0x9B183bb82Ce0DEf8d3A21dAD67dc84686B23de54', #3
    '0xa49010A6BCC8d3446aB7F55e8B4b3165d1cBAa71', #4
    '0xA32282D118c4747284110d25bD03247747528C44', #5
    '0xcbA6d4dc88890dBA0DC1671E5Ace3aB960CAC4C9', #6
    '0x0ea0398CF8F9eE889648A954077D04AD690b1AA5', #7
    '0x0A659f92A12C309D5360FF2e5EF23DD2292A1a0d', #8
    '0x19AffC1f95048521f3381F2A657FbE5899D02742', #9
    '0xE5dccE1ECA989cFEef98dE6A00212098007678a6', #10
]

if __name__=='__main__':
    appName = sys.argv[1]
    token_num = int(sys.argv[2])
    init_players = int(sys.argv[3])
    init_threshold = int(sys.argv[4])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.eth.defaultAccount = web3.eth.accounts[0]
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

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
    print(f'Deployed app to: {appAddress}\n')
    appContract = web3.eth.contract(address=appAddress, abi=abi)

    for token_id in range(1, token_num + 1):
        abi, bytecode = parse_contract('Token')
        tx_hash = web3.eth.contract(
            abi=abi,
            bytecode=bytecode).constructor().transact()
        web3.eth.wait_for_transaction_receipt(tx_hash)
        tokenAddress = web3.eth.waitForTransactionReceipt(tx_hash)['contractAddress']
        print(f'Deployed token {token_id} to: {tokenAddress}\n')
        tokenContract = web3.eth.contract(address=tokenAddress, abi=abi)
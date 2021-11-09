import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.utils import getAccount, parse_contract

url = 'ws://0.0.0.0:8546'

app_addr = '0xA0072d34984CC8de81b48923DE7d32e2AbC23265'

token_addrs = [
    '0x0000000000000000000000000000000000000000', #ETH
    '0xea53C26EA09eDdbf07B71902d08507b2ebB7DB96', #1
    '0xE38147Fc18547f25CE6962f5b677480E9bB10070', #2
    '0x28CE9c4B2de6e80dad6EFad64CA8cC78b59216dB', #3
    '0x7dA904c96276a6d1b388B694409394155FD3dB3E', #4
    '0xca43919ba9076d57710b2cC874ef9C5B99199387', #5
    '0x0DEC0818D2288920A76f2939739d50F05492f911', #6
    '0x3E43E3f698EB536E923Fa3dDD118f134da7Bb391', #7
    '0x64D86eF85f0c4C4FcAb4D28573574e384549f642', #8
    '0xBb557A2977A95be7b0Fa80007Df924362143E820', #9
    '0xD80FE8516Fd33568b3B764935DA41F9C48a4DA1C', #10
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
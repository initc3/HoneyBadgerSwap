import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.deploy import url, token_addrs
from ratel.src.python.utils import getAccount, parse_contract


def transferEther(client_account, server_account, amt):
    tx_hash = web3.eth.send_transaction({
        'to': client_account.address,
        'from': server_account.address,
        'value': amt
    })
    web3.eth.wait_for_transaction_receipt(tx_hash)

    print(f'**** client {client_id} ether balance {web3.eth.get_balance(client_account.address)}')

def transferToken(token_index, client_account, amt):
    token_addr = token_addrs[token_index]
    abi, bytecode = parse_contract('Token')
    tokenContract = web3.eth.contract(address=token_addr, abi=abi)

    tx_hash = tokenContract.functions.approve(client_account.address, amt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

    tx_hash = tokenContract.functions.transfer(client_account.address, amt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

    print(f'**** client {client_id} token {token_index} balance {tokenContract.functions.balanceOf(client_account.address).call()}')

def refill(client_id, token_id):
    client_account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    server_account = getAccount(web3, f'/opt/poa/keystore/server_0/')
    web3.eth.defaultAccount = server_account.address

    amt = int(1e8 * 1e18)

    transferEther(client_account, server_account, amt)
    transferToken(token_id, client_account, amt)

if __name__=='__main__':
    client_num = int(sys.argv[1])
    token_num = int(sys.argv[2])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    for client_id in range(1, client_num + 1):
        for token_id in range(1, token_num + 1):
            refill(client_id, token_id)


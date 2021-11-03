import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.deploy import url, addrs
from ratel.src.python.utils import getAccount, parse_contract


def transferEther():
    tx_hash = web3.eth.send_transaction({
        'to': client_account.address,
        'from': server_account.address,
        'value': amt
    })
    web3.eth.wait_for_transaction_receipt(tx_hash)
    print('**** Ether balance', web3.eth.get_balance(client_account.address))

def transferToken(token_index):
    token_addr = addrs[token_index]
    abi, bytecode = parse_contract('Token')
    tokenContract = web3.eth.contract(address=token_addr, abi=abi)

    tx_hash = tokenContract.functions.approve(client_account.address, amt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

    tx_hash = tokenContract.functions.transfer(client_account.address, amt).transact()
    web3.eth.wait_for_transaction_receipt(tx_hash)

    print('**** Token balance', token_index, tokenContract.functions.balanceOf(client_account.address).call())

if __name__=='__main__':
    client_id = int(sys.argv[1])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    client_account = getAccount(web3, f'/opt/poa/keystore/client_{client_id}/')
    server_account = getAccount(web3, f'/opt/poa/keystore/server_0/')
    web3.eth.defaultAccount = server_account.address

    amt = int(1e8 * 1e18)

    transferEther()
    transferToken(1)



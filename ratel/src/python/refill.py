import sys

from web3 import Web3
from web3.middleware import geth_poa_middleware
from ratel.src.python.deploy import url, token_addrs
from ratel.src.python.utils import getAccount, parse_contract


def transferEther(client_account, server_account, amt):
    web3.eth.send_transaction({
        'to': client_account.address,
        'from': server_account.address,
        'value': amt
    })


def transferToken(token_index, client_account, amt):
    token_addr = token_addrs[token_index]
    abi, bytecode = parse_contract('Token')
    tokenContract = web3.eth.contract(address=token_addr, abi=abi)

    tokenContract.functions.approve(client_account.address, amt).transact()

    tokenContract.functions.transfer(client_account.address, amt).transact()


def refill(receiver, token_id):
    client_account = getAccount(web3, f'/opt/poa/keystore/{receiver}/')
    server_account = getAccount(web3, f'/opt/poa/keystore/admin/')
    web3.eth.defaultAccount = server_account.address

    amt = int(1e8 * 1e18)

    if token_id == 0:
        transferEther(client_account, server_account, amt)
    else:
        transferToken(token_id, client_account, amt)


if __name__=='__main__':
    receiver = sys.argv[1]
    token_id = int(sys.argv[2])

    web3 = Web3(Web3.WebsocketProvider(url))
    web3.middleware_onion.inject(geth_poa_middleware, layer=0)

    refill(receiver, token_id)


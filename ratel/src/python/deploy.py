import json

from web3 import Web3
from web3.middleware import geth_poa_middleware

url = 'ws://0.0.0.0:8546'
web3 = Web3(Web3.WebsocketProvider(url))

def deploy_contract():
    tx_hash = web3.eth.contract(
        abi=abi,
        bytecode=bytecode).constructor().transact()

    tx_receipt = web3.eth.wait_for_transaction_receipt(tx_hash)
    print(tx_receipt)
    address = web3.eth.waitForTransactionReceipt(tx_hash)['contractAddress']
    return address

def deposit():
    pass

f = open('ratel/genfiles/build/contracts/Test.json')
data = json.load(f)

web3.eth.defaultAccount = web3.eth.accounts[0]

web3.middleware_onion.inject(geth_poa_middleware, layer=0)
abi = data['abi']
bytecode = data['bytecode']
address = deploy_contract()
print(f'Deployed to: {address}\n')
# address = '0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2'

myContract = web3.eth.contract(address=address, abi=abi)

print(myContract.functions.publicBalance('0x0000000000000000000000000000000000000000', web3.eth.defaultAccount).call())
tx_hash = myContract.functions.publicDeposit('0x0000000000000000000000000000000000000000',2**16).transact({'from':web3.eth.defaultAccount, 'value':10**15})
tx_receipt = web3.eth.wait_for_transaction_receipt(tx_hash)
print(tx_receipt)

print(myContract.functions.publicBalance('0x0000000000000000000000000000000000000000', web3.eth.defaultAccount).call())
tx_hash = myContract.functions.secretDeposit('0x0000000000000000000000000000000000000000',2**16).transact()
tx_receipt = web3.eth.wait_for_transaction_receipt(tx_hash)
print(tx_receipt)

print(myContract.functions.publicBalance('0x0000000000000000000000000000000000000000', web3.eth.defaultAccount).call())
tx_hash = myContract.functions.publicDeposit('0x0000000000000000000000000000000000000000',2**16).transact({'from':web3.eth.defaultAccount, 'value':10**15})
tx_receipt = web3.eth.wait_for_transaction_receipt(tx_hash)
print(tx_receipt)

print(myContract.functions.publicBalance('0x0000000000000000000000000000000000000000', web3.eth.defaultAccount).call())
tx_hash = myContract.functions.secretDeposit('0x0000000000000000000000000000000000000000',2**16).transact()
tx_receipt = web3.eth.wait_for_transaction_receipt(tx_hash)
print(tx_receipt)



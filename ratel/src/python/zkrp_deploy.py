import sys

from web3 import Web3
from ratel.src.python.utils import getAccount, parse_contract
from ratel.src.python.deploy import url, app_addr
from ratel.src.python.Server import Server
from ratel.src.python.utils import parse_contract, repeat_experiment
from ratel.genfiles.python.rockPaperScissorsRecover import recover


if __name__=='__main__':
    with open('ratel/genfiles/tmp.txt', 'r') as f:
        sum_zkrp = int(f.read())
    print(sum_zkrp)


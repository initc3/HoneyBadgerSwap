#!/usr/bin/env bash

poaDir=/opt/hbswap/poa
dataDir=/opt/hbswap/poa/data
keyStore=/opt/hbswap/poa/keystore/server_0

geth --datadir $dataDir init $poaDir/genesis.json

geth \
    --datadir $dataDir \
    --keystore $keyStore \
    --mine --allow-insecure-unlock --unlock 0 \
    --password $poaDir/empty_password.txt \
    --rpc --rpcaddr 0.0.0.0 --rpccorsdomain '*' --rpcapi admin,debug,eth,miner,net,personal,shh,txpool,web3 \
    --ws --wsaddr 0.0.0.0 --wsorigins '*' --wsapi admin,debug,eth,miner,net,personal,shh,txpool,web3 \
    --syncmode full \
    --ipcpath "$dataDir/geth.ipc"
    #2>> $dataDir/geth.log &

#!/usr/bin/env bash

POADIR=${POADIR:-/opt/poa}
DATADIR=${DATADIR:-/opt/poa/data}
KEYSTORE=${KEYSTORE:-/opt/poa/keystore/server_0}

geth --datadir $DATADIR init $POADIR/genesis.json

geth \
    --datadir $DATADIR \
    --keystore $KEYSTORE \
    --mine --allow-insecure-unlock --unlock 0 \
    --password $POADIR/empty_password.txt \
    --rpc \
    --rpcaddr 0.0.0.0 \
    --rpccorsdomain '*' \
    --rpcapi admin,debug,eth,miner,net,personal,shh,txpool,web3 \
    --ws \
    --wsaddr 0.0.0.0 \
    --wsorigins '*' \
    --wsapi admin,debug,eth,miner,net,personal,shh,txpool,web3 \
    --syncmode full \
    --ipcpath "$DATADIR/geth.ipc" \
    2>> $DATADIR/geth.log &

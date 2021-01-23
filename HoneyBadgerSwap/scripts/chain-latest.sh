#!/usr/bin/env bash

POADIR=${POADIR:-/opt/poa}
DATADIR=${POA_DATADIR:-/opt/poa/data}
KEYSTORE=${POA_KEYSTORE:-/opt/poa/keystore/server_0}

rm -rf $DATADIR
mkdir $DATADIR

geth --datadir $DATADIR init $POADIR/genesis.json

geth \
    --datadir $DATADIR \
    --keystore $KEYSTORE \
    --mine --allow-insecure-unlock --unlock 0 \
    --password $POADIR/empty_password.txt \
    --http \
    --http.addr 0.0.0.0 \
    --http.corsdomain '*' \
    --http.api admin,debug,eth,miner,net,personal,shh,txpool,web3 \
    --ws \
    --ws.addr 0.0.0.0 \
    --ws.origins '*' \
    --ws.api admin,debug,eth,miner,net,personal,shh,txpool,web3 \
    --syncmode full \
    --ipcpath "$DATADIR/geth.ipc"

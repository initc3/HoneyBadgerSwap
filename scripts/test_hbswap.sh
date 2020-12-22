#!/usr/bin/env bash
set -e

# bash Scripts/hbswap/scripts/test_hbswap.sh

players=4
threshold=1
port=5000

prog="malicious-shamir-party.x"

prepare() {
    make -j 8 $prog
#    Scripts/setup-online.sh $players
    Scripts/setup-ssl.sh $players

    kill -9 $(lsof -i:5000 -t) || true
    rm ./Persistence/* || true
    mkdir ./Persistence || true
}

compile() {
    ./compile.py -v -C -F 256 hbswap_init
    ./compile.py -v -C -F 256 hbswap_inputmask
    ./compile.py -v -C -F 256 hbswap_trade
}

run() {
    #./Scripts/mal-shamir.sh hbswap
    ./$prog -N $players -T $threshold -p 0 -pn $port $1 &
    ./$prog -N $players -T $threshold -p 1 -pn $port $1 &
    ./$prog -N $players -T $threshold -p 2 -pn $port $1
}

trade() {
    echo "inputmask-----------------------------------------------"
    run hbswap_inputmask
    echo "client-----------------------------------------------"
    python3 Scripts/hbswap/python/hbswap_client.py $1 $2
    echo "trade-----------------------------------------------"
    run hbswap_trade
    echo "org-----------------------------------------------"
    python3 Scripts/hbswap/python/server_org_file.py
}

prepare
compile

run hbswap_init
trade 1.1 -2.5
trade -1.1 2

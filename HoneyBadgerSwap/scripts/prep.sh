#!/usr/bin/env bash
set -e

players=3
threshold=1
port=5000

prog="malicious-shamir-party.x"

prepare() {
  Scripts/setup-ssl.sh $players

  pkill -f $prog || true
  pkill -f start_server.py || true
  pkill -f server.go || true

  rm -rf Persistence
  rm -rf Scripts/hbswap/data
  rm -rf Scripts/hbswap/log
  rm -rf Scripts/hbswap/db

  mkdir Persistence
  mkdir Scripts/hbswap/data
  mkdir Scripts/hbswap/log
  mkdir Scripts/hbswap/db
}

compile() {
  ./compile.py -v -C -F 256 $1
}

run() {
    ./$prog -N $players -T $threshold -p 0 -pn $port $1 &
    ./$prog -N $players -T $threshold -p 1 -pn $port $1 &
    ./$prog -N $players -T $threshold -p 2 -pn $port $1
}

org() {
  mv 'Persistence/Transactions-P0.data' 'Scripts/hbswap/data/Pool-P0.data'
  mv 'Persistence/Transactions-P1.data' 'Scripts/hbswap/data/Pool-P1.data'
  mv 'Persistence/Transactions-P2.data' 'Scripts/hbswap/data/Pool-P2.data'
}

prepare

compile hbswap_init
compile hbswap_trade_prep
compile hbswap_trade

run hbswap_init
org

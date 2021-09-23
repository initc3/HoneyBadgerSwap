#!/usr/bin/env bash

clear() {
  pkill -f python3 || true
  rm -rf /opt/hbswap/db/*
}

app=$1
players=$2
threshold=$3

bash chain-latest.sh &
sleep 3

clear
python3 -m ratel.src.python.deploy $app $players $threshold
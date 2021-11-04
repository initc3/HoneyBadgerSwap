#!/usr/bin/env bash

clear() {
  pkill -f python3 || true
  rm -rf /opt/hbswap/db/*
}

app=$1
token_num=$2
players=$3
threshold=$4

bash chain-latest.sh &
sleep 3

clear
python3 -m ratel.src.python.deploy $app $token_num $players $threshold
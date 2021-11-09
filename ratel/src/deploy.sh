#!/usr/bin/env bash

app=$1
token_num=$2
players=$3
threshold=$4

bash chain-latest.sh &
sleep 3

clear
python3 -m ratel.src.python.deploy $app $token_num $players $threshold

token_id=0
for (( server_id = 0; server_id < $players; server_id++ )) do
  python3 -m ratel.src.python.refill server_$server_id $token_id
done
#!/usr/bin/env bash
#####
##### ./ratel/benchmark/src/test_concurrent_trade_start.sh
#####

set -e
set -x

##### fixed parameter
players=4
threshold=1
token_A_id=0
#####
client_num=$1
concurrency=$2
token_num=$client_num

clear() {
  pkill -f python || true
  pkill -f ./malicious-shamir-party.x || true
  rm -rf /opt/hbswap/db/*
}

clear

bash ratel/src/deploy.sh hbswap $token_num $players $threshold

bash ratel/src/run.sh hbswap 0,1,2,3 $players $threshold $concurrency

for (( client_id = 1; client_id <= $client_num; client_id++ )) do
  token_B_id=$client_id
  python3 -m ratel.src.python.refill client_$client_id $token_A_id
  python3 -m ratel.src.python.refill client_$client_id $token_B_id
done

for (( client_id = 1; client_id <= $client_num; client_id++ )) do
  token_B_id=$client_id
  python3 -m ratel.src.python.hbswap.deposit $client_id $token_A_id 10000
  python3 -m ratel.src.python.hbswap.deposit $client_id $token_B_id 10000
done

for (( client_id = 1; client_id <= $client_num; client_id++ )) do
  token_B_id=$client_id
  python3 -m ratel.src.python.hbswap.initPool $client_id $token_A_id $token_B_id 1000 1000 &
  sleep 2
done
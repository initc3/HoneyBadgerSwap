#!/usr/bin/env bash
#####
##### ./ratel/benchmark/src/swap/test_real_data_trade_start.sh
#####

set -e
set -x

source ratel/src/utils.sh

##### fixed parameter
threshold=1
token_A_id=0
test=0
#####

##### get argv
players=$1
client_num=$2
concurrency=$3
#####
token_num=$client_num

bash ratel/src/deploy.sh hbswap $token_num $players $threshold

refill $players

ids=$(create_ids $players)
bash ratel/src/run.sh hbswap $ids $players $threshold $concurrency $test

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
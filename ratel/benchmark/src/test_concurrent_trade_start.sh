#!/usr/bin/env bash
#####
##### ./ratel/benchmark/src/test_concurrent_trade.sh [num_repetition]
#####

set -e
set -x

players=4
threshold=1

client_num=3
token_num=$client_num

concurrency=2

bash ratel/src/deploy.sh hbswap $token_num $players $threshold

bash ratel/src/run.sh hbswap 0,1,2,3 $players $threshold $concurrency

python3 -m ratel.src.python.refill $client_num $token_num

token_A_id=0
for (( client_id = 1; client_id <= $client_num; client_id++ )) do
  token_B_id=$client_id
  python3 -m ratel.src.python.hbswap.deposit $client_id $token_A_id 10000
  python3 -m ratel.src.python.hbswap.deposit $client_id $token_B_id 10000
  python3 -m ratel.src.python.hbswap.initPool $client_id $token_A_id $token_B_id 1000 1000
done
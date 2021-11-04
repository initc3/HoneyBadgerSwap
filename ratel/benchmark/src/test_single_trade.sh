#!/usr/bin/env bash
#####
##### ./ratel/benchmark/src/test_single_trade.sh [num_repetition]
#####

set -e
set -x

rm ratel/benchmark/data/latency.csv || true
rm ratel/benchmark/data/gas.csv || true

players=4
threshold=1

client_num=1
token_num=$client_num

concurrency=1

bash ratel/src/deploy.sh hbswap $token_num $players $threshold

bash ratel/src/run.sh hbswap 0,1,2,3 $players $threshold $concurrency

python3 -m ratel.src.python.refill $client_num $token_num

token_A_id=0
token_B_id=1
client_id=1

python3 -m ratel.src.python.hbswap.deposit $client_id $token_A_id 10000
python3 -m ratel.src.python.hbswap.deposit $client_id $token_B_id 10000

python3 -m ratel.src.python.hbswap.initPool $client_id $token_A_id $token_B_id 1000 1000

rep=$1
python3 -m ratel.src.python.hbswap.trade $client_id $token_A_id $token_B_id 0.5 -1 $rep
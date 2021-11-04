#!/usr/bin/env bash
#####
##### ./ratel/benchmark/src/test_concurrent_trade.sh [num_repetition]
#####

set -e
set -x

rm ratel/benchmark/data/latency.csv || true
rm ratel/benchmark/data/gas.csv || true

players=4
threshold=1

client_num=3

concurrency=3

bash ratel/src/run.sh hbswap 0,1,2,3 $players $threshold $concurrency

sleep 10

token_A_id=0

rep=$1
for (( client_id = 1; client_id <= $client_num; client_id++ )) do
  token_B_id=$client_id
  python3 -m ratel.src.python.hbswap.trade $client_id $token_A_id $token_B_id 0.5 -1 $rep &
done
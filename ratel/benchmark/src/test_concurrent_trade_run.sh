#!/usr/bin/env bash
#####
##### ./ratel/benchmark/src/test_concurrent_trade_run.sh
#####

set -e
set -x

source ratel/src/utils.sh

##### fixed parameter
threshold=1
token_A_id=0
test=0
#####

mkdir -p ratel/benchmark/data

players=$1
client_num=$2
concurrency=$3
rep=$4

ids=$(create_ids $players)
bash ratel/src/run.sh hbswap $ids $players $threshold $concurrency $test

python3 -m ratel.benchmark.src.set_up_offline_data $players $threshold $concurrency

for ((server_id = 0; server_id < $players; server_id++ )) do
  rm ratel/benchmark/data/latency_$server_id.csv || true
done
rm ratel/benchmark/data/gas.csv || true
rm ratel/benchmark/data/fig.pdf || true

sleep 10

for (( client_id = 1; client_id <= $client_num; client_id++ )) do
  token_B_id=$client_id
  python3 -m ratel.src.python.hbswap.trade $client_id $token_A_id $token_B_id 0.5 -1 $rep &
  sleep 1
done

#!/usr/bin/env bash

set -x

players=$1
concurrency=$2

for ((i = 0; i < $concurrency; i += 1)) do
  for ((server_id = 0; server_id < $players; server_id += 1)) do
    port=$((5000 + $i * 100))
    echo $port
    cp Persistence/Transactions-P$server_id-$port.data ratel/benchmark/data/sharefiles/Transactions-P$server_id-$port.data
  done
done
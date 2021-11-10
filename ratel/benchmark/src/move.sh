#!/usr/bin/env bash

set -x

for ((server_id = 0; server_id < 4; server_id++)) do
  for ((port = 5000; port < 7000; port += 100)) do
    cp Persistence/Transactions-P$server_id-$port.data ratel/benchmark/data/sharefiles/Transactions-P$server_id-$port.data
  done
done
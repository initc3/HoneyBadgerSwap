#!/usr/bin/env bash

source ratel/src/utils.sh

set -e
set -x

##### fixed parameter
players=3
threshold=1
batch_num=10
#####

kill_python
kill_mpc


rm ratel/benchmark/data/inputmask_generation_latency_*.csv || true

for ((server_id = 0; server_id < $players; server_id++ )) do
  python3 -m ratel.benchmark.src.test_inputmask_generation $server_id $players $threshold $batch_num &
done
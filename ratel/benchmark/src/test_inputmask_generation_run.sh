#!/usr/bin/env bash

##### fixed parameter
players=4
#####

.utils
clear
rm ratel/benchmark/data/inputmask_generation_latency_*.csv || true

for ((server_id = 0; server_id < $players; server_id++ )) do
  python3 -m ratel.benchmark.src.test_inputmask_generation $server_id &
done
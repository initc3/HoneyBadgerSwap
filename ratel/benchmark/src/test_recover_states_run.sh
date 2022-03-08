#!/usr/bin/env bash
source ratel/src/utils.sh

set -e
set -x

seq=3

online_players=$1
repetion=$2

kill_python
kill_mpc

rm ratel/benchmark/data/recover_states.csv || true

for ((server_id = 0; server_id < $online_players; server_id++ )) do
  python3 -m ratel.benchmark.src.test_recover_states $server_id $(($online_players-1)) $seq $repetion &
done
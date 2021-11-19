#!/usr/bin/env bash
source ./ratel/benchmark/src/utils.sh

set -e
set -x

app=$1
IDs=$2
players=$3
threshold=$4
concurrency=$5
test=$6

kill_python
kill_mpc

IFS=','
read -a strarr <<< "$IDs"

for id in "${strarr[@]}";
do
  python3 -m ratel.src.python.$app.run $id $players $threshold $concurrency $test &
done

sleep 3
#!/usr/bin/env bash
source ratel/src/utils.sh

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

mkdir -p ratel/log

for id in "${strarr[@]}";
do
  python3 -m ratel.src.python.$app.run $id $players $threshold $concurrency $test > ratel/log/server_$id.log 2>&1 &
done

sleep 3
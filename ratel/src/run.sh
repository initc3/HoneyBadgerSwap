#!/usr/bin/env bash

set -e
set -x

app=$1
IDs=$2
players=$3
threshold=$4
concurrency=$5

IFS=','

read -a strarr <<< "$IDs"


for id in "${strarr[@]}";
do
  pkill -f "python3 -m ratel.src.python.$app.run $id" || true
done

for id in "${strarr[@]}";
do
  python3 -m ratel.src.python.$app.run $id $players $threshold $concurrency &
done
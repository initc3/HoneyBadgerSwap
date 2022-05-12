#!/usr/bin/env bash

bash chain-latest.sh &
sleep 3

pkill -f python3 || true

rm -rf /opt/hbswap/db/*
python3 -m ratel.src.python.deploy fabcar

python3 -m ratel.src.python.fabcar.run 0 &
python3 -m ratel.src.python.fabcar.run 1 &
python3 -m ratel.src.python.fabcar.run 2 &
python3 -m ratel.src.python.fabcar.run 3 &

# bash ratel/src/python/fabcar/run.sh
# python3 -m ratel.src.python.fabcar.fabcar
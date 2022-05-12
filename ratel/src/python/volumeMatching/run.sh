#!/usr/bin/env bash

bash chain-latest.sh &
sleep 3

pkill -f python3 || true

rm -rf /opt/hbswap/db/*
python3 -m ratel.src.python.deploy VolumeMatching

python3 -m ratel.src.python.volumeMatching.run 0 &
python3 -m ratel.src.python.volumeMatching.run 1 &
python3 -m ratel.src.python.volumeMatching.run 2 &
python3 -m ratel.src.python.volumeMatching.run 3 &

# bash ratel/src/python/volumeMatching/run.sh
# python3 -m ratel.src.python.volumeMatching.deposit
# python3 -m ratel.src.python.volumeMatching.submitBid
# python3 -m ratel.src.python.volumeMatching.volumeMatch
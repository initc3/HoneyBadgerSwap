#!/usr/bin/env bash

bash chain-latest.sh &
sleep 3

pkill -f python3 || true

rm -rf /opt/hbswap/db/*
python3 -m ratel.src.python.deploy review

python3 -m ratel.src.python.review.run 0 &
python3 -m ratel.src.python.review.run 1 &
python3 -m ratel.src.python.review.run 2 &
python3 -m ratel.src.python.review.run 3 &

# bash ratel/src/python/review/run.sh
# python3 -m ratel.src.python.review.review
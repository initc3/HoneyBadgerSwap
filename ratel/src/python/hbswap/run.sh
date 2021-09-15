#!/usr/bin/env bash

bash chain-latest.sh &
sleep 3

pkill -f python3 || true

rm -rf /opt/hbswap/db/*
python3 -m ratel.src.python.deploy hbswap

python3 -m ratel.src.python.hbswap.run 0 &
python3 -m ratel.src.python.hbswap.run 1 &
python3 -m ratel.src.python.hbswap.run 2 &
# python3 -m ratel.src.python.hbswap.run 3 &

# bash ratel/src/python/hbswap/run.sh
# python3 -m ratel.src.python.hbswap.deposit
# python3 -m ratel.src.python.hbswap.withdraw
# python3 -m ratel.src.python.hbswap.initPool
# python3 -m ratel.src.python.hbswap.addLiquidity
# python3 -m ratel.src.python.hbswap.removeLiquidity
# python3 -m ratel.src.python.hbswap.trade

# python3 -m ratel.src.python.hbswap.recover 3
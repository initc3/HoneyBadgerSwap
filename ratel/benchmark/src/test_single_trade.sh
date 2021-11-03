#!/usr/bin/env bash

rm ratel/benchmark/data/log.csv

bash ratel/src/deploy.sh hbswap 4 1

bash ratel/src/run.sh hbswap 0,1,2,3 4 1

python3 -m ratel.src.python.refill 1

python3 -m ratel.src.python.hbswap.deposit 1 0x0000000000000000000000000000000000000000 10000
python3 -m ratel.src.python.hbswap.deposit 1 0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385 10000

python3 -m ratel.src.python.hbswap.initPool 1 0x0000000000000000000000000000000000000000 0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385 1000 1000

rep=$1
for ((i=1;i<=$rep;i++)); do
  echo '!!!!' $i
  sleep 1
  python3 -m ratel.src.python.hbswap.trade 1 0x0000000000000000000000000000000000000000 0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385 0.5 -1
  sleep 1
  python3 -m ratel.src.python.hbswap.trade 1 0x0000000000000000000000000000000000000000 0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385 -1 0.5
done
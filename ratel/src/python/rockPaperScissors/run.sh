#!/usr/bin/env bash

bash chain-latest.sh &
sleep 3

pkill -f python3 || true

rm -rf /opt/hbswap/db/*
python3 -m ratel.src.python.deploy RockPaperScissors

python3 -m ratel.src.python.rockPaperScissors.run 0 &
python3 -m ratel.src.python.rockPaperScissors.run 1 &
python3 -m ratel.src.python.rockPaperScissors.run 2 &
python3 -m ratel.src.python.rockPaperScissors.run 3 &

# bash ratel/src/python/rockPaperScissors/run.sh
# python3 -m ratel.src.python.rockPaperScissors.createGame
# python3 -m ratel.src.python.rockPaperScissors.joinGame
# python3 -m ratel.src.python.rockPaperScissors.startRecon
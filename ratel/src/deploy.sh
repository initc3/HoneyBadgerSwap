#!/usr/bin/env bash
source ratel/src/utils.sh

set -e
set -x

app=$1
token_num=$2
players=$3
threshold=$4

kill_python
kill_mpc
clear_db

bash chain-latest.sh

python3 -m ratel.src.python.deploy $app $token_num $players $threshold
#!/usr/bin/env bash
source ratel/src/utils.sh

set -e
set -x

##### fixed parameter
threshold=1
token_num=1
concurrency=1
client_id=1
token_A_id=0
token_B_id=1
test=1
#####

online_players=$1

bash ratel/src/deploy.sh hbswap $token_num $online_players $threshold

refill $online_players

ids=$(create_ids $online_players)
bash ratel/src/run.sh hbswap $ids $online_players $threshold $concurrency $test

python3 -m ratel.src.python.refill client_$client_id $token_A_id
python3 -m ratel.src.python.refill client_$client_id $token_B_id

python3 -m ratel.src.python.hbswap.deposit $client_id $token_A_id 10000
python3 -m ratel.src.python.hbswap.deposit $client_id $token_B_id 10000

python3 -m ratel.src.python.hbswap.initPool $client_id $token_A_id $token_B_id 1000 1000

python3 -m ratel.src.python.hbswap.trade $client_id $token_A_id $token_B_id 0.5 -1 1





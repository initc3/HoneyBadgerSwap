#!/usr/bin/env bash

##### fixed parameter
players=4
threshold=1
token_num=1
#####

.utils
clear

bash ratel/src/deploy.sh hbswap $token_num $players $threshold

./ratel/benchmark/src/test_inputmask_generation_run.sh
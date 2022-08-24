#!/usr/bin/env bash

./latency-control.sh stop

./ratel/benchmark/src/test_concurrent_trade_start.sh 3 3 10

sleep 20

./latency-control.sh start 200 50

./ratel/benchmark/src/test_concurrent_trade_run.sh 3 3 10 1

#sleep 1000
#
#python3 -m ratel.benchmark.src.test_mpc 3 1 10


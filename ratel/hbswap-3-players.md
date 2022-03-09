

`bash setup-ssl.sh 3`

`bash ratel/src/compile.sh hbswap`

`./ratel/benchmark/src/test_concurrent_trade_start.sh 3 1 1`

`./ratel/benchmark/src/test_concurrent_trade_run.sh 3 1 1 1`

`python3 -m ratel.benchmark.src.plot ratel/benchmark/data`

`./latency-control.sh stop`

`./latency-control.sh start 200 50`

`python3 -m ratel.benchmark.src.test_mpc 3`
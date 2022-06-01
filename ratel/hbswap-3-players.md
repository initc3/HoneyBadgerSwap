

`docker-compose build --no-cache --pull`

`docker-compose up -d`

`docker exec -it honeybadgerswap_dev_1 bash`

`bash setup-ssl.sh 3`

`bash ratel/src/compile.sh hbswap 0 1`

```
./ratel/benchmark/src/test_concurrent_trade_start.sh [players] [client_num] [concurrency]

./ratel/benchmark/src/test_concurrent_trade_start.sh 3 1 1
./ratel/benchmark/src/test_concurrent_trade_start.sh 3 2 2
./ratel/benchmark/src/test_concurrent_trade_start.sh 3 4 4
./ratel/benchmark/src/test_concurrent_trade_start.sh 3 8 8
./ratel/benchmark/src/test_concurrent_trade_start.sh 3 16 16
```

```
./ratel/benchmark/src/test_concurrent_trade_run.sh [players] [client_num] [concurrency] [rep]

./ratel/benchmark/src/test_concurrent_trade_run.sh 3 1 1 20
./ratel/benchmark/src/test_concurrent_trade_run.sh 3 2 2 20
./ratel/benchmark/src/test_concurrent_trade_run.sh 3 4 4 20
./ratel/benchmark/src/test_concurrent_trade_run.sh 3 8 8 20
./ratel/benchmark/src/test_concurrent_trade_run.sh 3 16 16 20
```

`python3 -m ratel.benchmark.src.trade_throughput ratel/benchmark/data`

```
python3 -m ratel.benchmark.src.trade_latency [players] [dir]
python3 -m ratel.benchmark.src.trade_latency 3 ratel/benchmark/data
```

`python3 -m ratel.benchmark.src.trade_plot`

`./latency-control.sh stop`

`./latency-control.sh start 200 50`

```
./compile.py -v -C -F 128 ratel/genfiles/mpc/hbswapTrade1.mpc
python3 -m ratel.benchmark.src.test_mpc 3 1 1
python3 -m ratel.benchmark.src.test_mpc 3 1 10
```

```
./ratel/benchmark/src/move.sh 3 1
```

```
./ratel/benchmark/src/test_inputmask_generation_run.sh
python3 -m ratel.benchmark.src.test_inputmask_generation_plot ratel/benchmark/data 3 1 10
```

```
./latency-control.sh stop
ratel/benchmark/src/swap/test_real_data_trade_start.sh 3 1 1
./latency-control.sh start 200 50
ratel/benchmark/src/swap/test_real_data_trade_run.sh 3 1 1

python3 -m ratel.benchmark.src.swap.analyze 3 ratel/benchmark/data
python3 -m ratel.benchmark.src.swap.collect 3 ratel/benchmark/data
python3 -m ratel.benchmark.src.swap.simulate 172800 432000 traderjoev2_USDC.e_WAVAX
```

```
python3 -m ratel.src.python.hbswap.updateBatchPrice
```

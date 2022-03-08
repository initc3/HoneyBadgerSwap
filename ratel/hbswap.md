# HoneyBadgerSwap

`docker-compose build`

`docker-compose up -d`

`docker exec -it honeybadgerswap_dev_1 bash`

`bash ratel/src/deploy.sh hbswap 1 4 1`

```
python3 -m ratel.src.python.refill server_0 0 &
python3 -m ratel.src.python.refill server_1 0 &
python3 -m ratel.src.python.refill server_2 0 &
python3 -m ratel.src.python.refill server_3 0 &
python3 -m ratel.src.python.refill client_1 0 &
python3 -m ratel.src.python.refill client_1 1
```

`bash ratel/src/compile.sh hbswap`

`bash ratel/src/run.sh hbswap 0,1,2,3 4 1 1 0`

```
python3 -m ratel.src.python.hbswap.deposit 1 0 10000
python3 -m ratel.src.python.hbswap.deposit 1 1 10000
```

`python3 -m ratel.src.python.hbswap.initPool 1 0 1 1000 1000`

`python3 -m ratel.src.python.hbswap.trade 1 0 1 0.5 -1 1`

`python3 -m ratel.src.python.hbswap.withdraw `

###Introduce latency
```bash
./latency-control.sh start 200 50
./latency-control.sh stop
```

###Test single trade
```
./ratel/benchmark/src/test_concurrent_trade_start.sh 1 1
./ratel/benchmark/src/test_concurrent_trade_run.sh 1 1 1
```

###Test concurrent trades
```
./latency-control.sh stop
./ratel/benchmark/src/test_concurrent_trade_start.sh 10 10
./ratel/benchmark/src/test_concurrent_trade_start.sh 2 1

./latency-control.sh start 200 50
./ratel/benchmark/src/test_concurrent_trade_run.sh 10 10 10
./ratel/benchmark/src/test_concurrent_trade_run.sh 10 5 5
```

###Test MP-SPDZ concurrency
```bash
python3 -m ratel.benchmark.src.test_mpc
```

###Test recover states
```bash
./ratel/benchmark/src/test_recover_states_start.sh 4
./ratel/benchmark/src/test_recover_states_run.sh 4 5
```


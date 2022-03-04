# HoneyBadgerSwap

`docker-compose build`

`docker-compose up -d`

`docker exec -it honeybadgerswap_dev_1 bash`

`bash ratel/src/compile.sh`

Start local private blockchain and deploy application contract:
```
bash ratel/src/deploy.sh [app]
```

Test recovery mechanism (run the following commands one by one):
```
bash ratel/src/deploy.sh hbswap 3 1

bash ratel/src/run.sh hbswap 0,1,2

python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 
python3 -m ratel.src.python.hbswap.deposit 0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2 1 

bash ratel/src/run.sh hbswap 3

python3 -m ratel.src.python.hbswap.initPool
```

Test concurrency:
```
bash ratel/src/deploy.sh hbswap 4 1

bash ratel/src/run.sh hbswap 0,1,2,3 4 1

python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &
python3 -m ratel.src.python.hbswap.deposit 0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2 1 &
python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &
python3 -m ratel.src.python.hbswap.deposit 0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2 1 &

python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &
python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &
python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &

```

Introduce latency:
```bash
./latency-control.sh start 200 50
./latency-control.sh stop
```

Test single trade
```
./ratel/benchmark/src/test_concurrent_trade_start.sh 1 1
./ratel/benchmark/src/test_concurrent_trade_run.sh 1 1 1
```

Test concurrent trade
```
./latency-control.sh stop
./ratel/benchmark/src/test_concurrent_trade_start.sh 10 10
./ratel/benchmark/src/test_concurrent_trade_start.sh 2 1

./latency-control.sh start 200 50
./ratel/benchmark/src/test_concurrent_trade_run.sh 10 10 10
./ratel/benchmark/src/test_concurrent_trade_run.sh 10 5 5
```

Test MP-SPDZ concurrency
```bash
python3 -m ratel.benchmark.src.test_mpc
```

Test recover states
```bash
./ratel/benchmark/src/test_recover_states_start.sh 4
./ratel/benchmark/src/test_recover_states_run.sh 4 5
```


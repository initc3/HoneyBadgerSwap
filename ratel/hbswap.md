# HoneyBadgerSwap

`docker-compose build`

`docker-compose up -d`

`docker exec -it honeybadgerswap_dev_1 bash`

`bash setup-ssl.sh 4`

`bash ratel/src/compile.sh hbswap 0 0`

`bash ratel/src/deploy.sh hbswap 1 4 1`

```
python3 -m ratel.src.python.refill server_0 0 &
python3 -m ratel.src.python.refill server_1 0 &
python3 -m ratel.src.python.refill server_2 0 &
python3 -m ratel.src.python.refill server_3 0 &
python3 -m ratel.src.python.refill client_1 0 &
python3 -m ratel.src.python.refill client_1 1
```

`bash ratel/src/run.sh hbswap 0,1,2,3 4 1 1 0`

```
python3 -m ratel.src.python.hbswap.deposit 1 0 10000
python3 -m ratel.src.python.hbswap.deposit 1 1 10000
```

`python3 -m ratel.src.python.hbswap.initPool 1 0 1 1000 1000`

`python3 -m ratel.src.python.hbswap.trade 1 0 1 0.5 -1 1`

`python3 -m ratel.src.python.hbswap.withdraw 1 0 2000`



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
./ratel/benchmark/src/test_concurrent_trade_start.sh 4 1 1
./ratel/benchmark/src/test_concurrent_trade_run.sh 4 1 1 1
```

Test concurrent trade
```
./latency-control.sh stop
./ratel/benchmark/src/test_concurrent_trade_start.sh 4 10 10
./ratel/benchmark/src/test_concurrent_trade_start.sh 4 2 2

./latency-control.sh start 200 50
./ratel/benchmark/src/test_concurrent_trade_run.sh 4 10 10 10
./ratel/benchmark/src/test_concurrent_trade_run.sh 4 10 5 5
./ratel/benchmark/src/test_concurrent_trade_run.sh 4 2 2 1
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

To give the proof zkrp(amtA * amtB < 0)

In server side, idxAmtA, maskedAmtA, idxAmtB, maskedAmtB; 

recover_input: amtA_shares, amtB_shares;











# def recover_input(db, masked_value, idx): # return: int
#     try:
#         input_mask_share = db.Get(key_inputmask_index(idx))
#     except KeyError:
#         input_mask_share = bytes(0)
#     input_mask_share = int.from_bytes(input_mask_share, 'big')
#     return (masked_value - input_mask_share) % prime



----------------------------------------------------------------

prove zkrp(amtA < 0)

client: generate idxAmtA, maskedAmtA,    get_zkrp()->proofValue, commitmentValue, idxBlindingValue, maskedBlindingValue

server:

amtA_share = maskedAmtA - input_mask_share_with_key_idxAmtA : recover_input (idxAmtA,maskedAmtA)

blidingValue_share = maskedBlindingValue - input_mask_share_with_key_idxBlindingValue: recover_input ()

compute g^{amtA_share} * h^{blindingValue_share}



every server broadcast share of commitment & 

reconstruct commitment Value & 

check whether the aggregate_commitment_from_servers == commitment_given_by_client




# RockPaperScissors

### Build docker image:
`docker-compose -f docker-compose-dev.yml build --no-cache`

### Start docker container:
`docker-compose -f docker-compose-dev.yml up -d`

### Enter docker container:
`docker exec -it honeybadgerswap_dev_1 bash`

### Compile ratel program:
`bash ratel/src/compile.sh`

### Start local private blockchain and deploy application contract:
```
bash ratel/src/deploy.sh [app_name] [token_num] [MPC_server_number] [threshold]

bash ratel/src/deploy.sh rockPaperScissors 0 4 1
```

### Transfer Ether(token_id=0) to MPC servers and clients for them to pay transaction fee
```
python3 -m ratel.src.python.refill [user_name (see available choices in poa/keystore/)] [token_id]

python3 -m ratel.src.python.refill server_0 0
python3 -m ratel.src.python.refill server_1 0
python3 -m ratel.src.python.refill server_2 0
python3 -m ratel.src.python.refill server_3 0

python3 -m ratel.src.python.refill client_1 0
python3 -m ratel.src.python.refill client_2 0
```

### Start MPC servers to monitor events emitted by application contract and take MPC  tasks:
```
bash ratel/src/run.sh [app_name] [MPC_server_IDs] [MPC_server_number] [threshold] [concurrency] [test_flag]

bash ratel/src/run.sh rockPaperScissors 0,1,2,3 4 1 1 0
```

### Interact with app contract
```
python3 -m ratel.src.python.rockPaperScissors.interact 
```


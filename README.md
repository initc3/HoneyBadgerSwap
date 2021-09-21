# HoneyBadgerSwap

`docker-compose -f docker-compose-dev.yml build --no-cache`

`docker-compose -f docker-compose-dev.yml up -d`

`docker exec -it honeybadgerswap_dev_1 bash`

`bash ratel/src/compile.sh`

Start local private blockchain and deploy application contract:
```
bash ratel/src/deploy.sh [app]
```

Test recovery mechanism:
```
bash ratel/src/deploy.sh hbswap
pkill -f python3 || true
bash ratel/src/run.sh hbswap 0,1,2
python3 -m ratel.src.python.hbswap.deposit
bash ratel/src/run.sh hbswap 3
python3 -m ratel.src.python.hbswap.initPool
```


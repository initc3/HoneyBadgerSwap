# HoneyBadgerSwap

`docker-compose -f docker-compose-dev.yml build --no-cache`

`docker-compose -f docker-compose-dev.yml up -d`

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
python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &
python3 -m ratel.src.python.hbswap.deposit 0x0000000000000000000000000000000000000000 1 &

```


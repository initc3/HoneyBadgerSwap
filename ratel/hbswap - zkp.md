# HoneyBadgerSwap

`docker-compose build`

`docker-compose up -d`

`docker exec -it honeybadgerswap_dev_1 bash`

`bash setup-ssl.sh 4 /opt/ssl`

```
curl https://sh.rustup.rs -sSf | sh  # enter directly
```

```
source "$HOME/.cargo/env"
cd ratel/src/zkrp_pyo3
pip install .
cd ../../..
```

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

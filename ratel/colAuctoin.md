# Collateral auction

This collateral auction application in the MPC sidechain framework provides privacy of the auction details. Every bidders submits bids in secret form. Therefore, each person knows only the amount of collateral he or she wants to purchase. Over time, the price decreases until the auction ends or the price falls below the floor price. At last, each auction will return whether it ends successfully.


### prerequesite

#### clone the repo
```
git clone https://github.com/xmhuangzhen/HoneyBadgerSwap.git
cd HoneyBadgerSwap
```
#### install docker-engine
```
sudo apt-get remove docker docker.io containerd runc
sudo apt-get update
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

#### install docker-compose
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

### running the application

#### pull docker image & enter the container
```
sudo docker pull initc3/honeybadgerswap:8fc1863
sudo docker tag initc3/honeybadgerswap:8fc1863 hbswap:latest
sudo docker-compose up -d
sudo docker exec -it honeybadgerswap_dev_1 bash
```

#### setup ssl connection
```
bash setup-ssl.sh 4
```

#### compile the ratel
```
bash ratel/src/compile.sh colAuction
```

#### Start local private blockchain and deploy application contract
```
bash ratel/src/deploy.sh colAuction 0 4 1
```
#### Transfer Ether(token_id=0) to MPC servers and clients for them to pay transaction fee

```
python3 -m ratel.src.python.refill server_0 0 \
& python3 -m ratel.src.python.refill server_1 0 \
& python3 -m ratel.src.python.refill server_2 0 \
& python3 -m ratel.src.python.refill server_3 0 
```

```
python3 -m ratel.src.python.refill client_1 0 \
& python3 -m ratel.src.python.refill client_2 0 \
& python3 -m ratel.src.python.refill client_3 0 \
& python3 -m ratel.src.python.refill client_4 0 \
& python3 -m ratel.src.python.refill client_5 0
```
#### Start MPC servers to monitor events emitted by application contract and take MPC tasks:

```
bash ratel/src/run.sh colAuction 0,1,2,3 4 1 1 0
```

#### Interact with app contract
```
python3 -m ratel.src.python.colAuction.check & 
python3 -m ratel.src.python.colAuction.interact 
```



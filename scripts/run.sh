#!/usr/bin/env bash
set -e

eth_host=${1:-localhost}
leader_host=${2:-localhost}

. Scripts/hbswap/scripts/utils.sh

prepare_nodes() {
  echo 'Preparing mpc nodes...'
  bash Scripts/hbswap/scripts/prep.sh
  echo 'Finished preparing mpc nodes'
}

start_local_network() {
  echo 'Staring local network...'
  pkill -f geth | true
  rm -rf Scripts/hbswap/poa/data
  bash Scripts/hbswap/scripts/chain.sh
  echo 'Finished staring local network'
}

deploy_contract() {
  echo 'Deploying contracts...'
  go run Scripts/hbswap/go/deploy/deploy.go $eth_host
  echo 'Finished deploying contracts'
}

shut_down() {
  echo 'Shuting down previous instances...'
  pkill -f random-shamir.x || true
  pkill -f $prog || true
  pkill -f start_server.py || true
  pkill -f server.go || true

  rm -rf Scripts/hbswap/log
  mkdir Scripts/hbswap/log

  sleep 2
  echo 'Finished shuting down previous instances'
}

start_servers() {
  echo 'Starting servers...'
  mpcserver 0 &
  mpcserver 1 &
  mpcserver 2 &
  mpcserver 3 &
  httpserver 0 &
  httpserver 1 &
  httpserver 2 &
  httpserver 3 &
  echo 'Finished starting servers'
}

deposit() {
  go run Scripts/hbswap/go/client/deposit.go $1 $2 $3 $4 $5 $eth_host
  go run Scripts/hbswap/go/client/secret_deposit.go $1 $2 $3 $4 $5 $eth_host
}

init_pool() {
  deposit $1 $2 $3 $4 $5
  go run Scripts/hbswap/go/client/init_pool.go $1 $2 $3 $4 $5 $eth_host
}

add_liquidity() {
  deposit $1 $2 $3 $4 $5
  go run Scripts/hbswap/go/client/add_liquidity.go $1 $2 $3 $4 $5
}

remove_liquidity() {
  go run Scripts/hbswap/go/client/remove_liquidity.go $1 $2 $3 $4
}

withdraw() {
  go run Scripts/hbswap/go/client/secret_withdraw.go $1 $2 $3 $4 $5 $eth_host
  go run Scripts/hbswap/go/client/withdraw.go $1 $2 $3 $4 $5 $eth_host
}

trade() {
  go run Scripts/hbswap/go/client/trade.go $1 $2 $3 $4 $5 $eth_host
}

httpserver() {
  python3 Scripts/hbswap/python/server/start_server.py $1
}

mpcserver() {
  go run Scripts/hbswap/go/server/server.go $1 $eth_host $leader_host > Scripts/hbswap/log/mpc_server_$1.log 2>&1
}

prepare_nodes

start_local_network

deploy_contract

shut_down

start_servers

init_pool 0 $eth $token_1 10 20
#add_liquidity 0 $eth $token_1 10 10
#remove_liquidity 0 $eth $token_1 10
#init_pool 0 $eth $token_2 10 20

deposit 0 $eth $token_1 10 10
#deposit 0 $eth $token_2 10 10

#python3 Scripts/hbswap/python/client/req_balance.py $eth $user
#python3 Scripts/hbswap/python/client/req_balance.py $token_1 $user
#python3 Scripts/hbswap/python/client/req_balance.py $token_2 $user

trade 0 $eth $token_1 1.1 -2.5
#python3 Scripts/hbswap/python/client/req_price.py 1
trade 0 $eth $token_1 -1.1 2
#python3 Scripts/hbswap/python/client/req_price.py 2
#
#trade 0 $eth $token_2 1.1 -2.5
#python3 Scripts/hbswap/python/client/req_price.py 3
#trade 0 $eth $token_2 -1.1 2
#python3 Scripts/hbswap/python/client/req_price.py 4

#withdraw 0 $eth $token_1 9 9
#withdraw 0 $eth $token_2 9 9







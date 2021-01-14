#!/usr/bin/env bash
set -e

#host=${1:-localhost}
node_id=$1
eth_hostname=$2
leader_hostname=$3

# Place the data where MP-SPDZ expects it
setup_data() {
    rm -rf Scripts/hbswap/log Scripts/hbswap/data/* Scripts/hbswap/poa/data /opt/hbswap/db
    mkdir -p Persistence Player-Data Scripts/hbswap/log
    mkdir -p /opt/hbswap/db /opt/hbswap/inputmask-shares /opt/hbswap/preprocessing-data
    # Copy the private key, where MP-SPDZ expects it to be (under Player-Data/).
    cp /opt/hbswap/secrets/P$node_id.key Player-Data/
    # Copy the public keys of all players
    cp /opt/hbswap/public-keys/* Player-Data/
}


httpserver() {
  python3 Scripts/hbswap/python/server/start_server.py $1
}

mpcserver() {
  go run Scripts/hbswap/go/server/server.go $1 $eth_hostname $leader_hostname
  #go run Scripts/hbswap/go/server/server.go -serverid=$1 -ethhost=$eth_hostname -leaderhost=$leader_hostname
  #go run Scripts/hbswap/go/server/server.go $1 $eth_hostname $leader_hostname > Scripts/hbswap/log/mpc_server_$1.log 2>&1
}

setup_data
httpserver $node_id & mpcserver $node_id

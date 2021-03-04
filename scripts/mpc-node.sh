#!/usr/bin/env bash
set -e

#host=${1:-localhost}
node_id=$1
eth_hostname=$2
leader_hostname=$3
#network=${4:-privatenet}
config=${4:-/opt/hbswap/conf/server.toml}
go_code_path=/go/src/github.com/initc3/HoneyBadgerSwap/src/go

# Place the data where MP-SPDZ expects it
setup_data() {
    rm -rf /opt/hbswap/db
    mkdir -p Persistence Player-Data
    mkdir -p /opt/hbswap/db /opt/hbswap/inputmask-shares /opt/hbswap/preprocessing-data
    # Copy the public keys of all players
    cp /opt/hbswap/public-keys/* Player-Data/
    # Symlink to the private key, to where MP-SPDZ expects it to be (under Player-Data/).
    ln -s /run/secrets/P$node_id.key Player-Data/P$node_id.key
}


httpserver() {
  python -m honeybadgerswap.server.start_server $1
}

mpcserver() {
  #go run $go_code_path/server/server.go -n $network $1 $leader_hostname $eth_hostname
  go run $go_code_path/server/server.go -config $config $1 $leader_hostname $eth_hostname
}

setup_data
httpserver $node_id & mpcserver $node_id

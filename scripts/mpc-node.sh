#!/usr/bin/env bash
set -e

#host=${1:-localhost}
node_id=$1
eth_hostname=$2
leader_hostname=$3
go_code_path=/go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap/go

# Place the data where MP-SPDZ expects it
setup_data() {
    rm -rf /opt/hbswap/db
    mkdir -p Persistence Player-Data
    mkdir -p /opt/hbswap/db /opt/hbswap/inputmask-shares /opt/hbswap/preprocessing-data
    # Copy the private key, where MP-SPDZ expects it to be (under Player-Data/).
    cp /opt/hbswap/secrets/P$node_id.key Player-Data/
    # Copy the public keys of all players
    cp /opt/hbswap/public-keys/* Player-Data/
}


httpserver() {
  python -m honeybadgerswap.server.start_server $1
}

mpcserver() {
  go run $go_code_path/server/server.go $1 $eth_hostname $leader_hostname
}

setup_data
httpserver $node_id & mpcserver $node_id

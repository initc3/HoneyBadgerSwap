#!/usr/bin/env bash
set -e

node_id=$1
config=${2:-/opt/hbswap/conf/server.toml}
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
  python -m honeybadgerswap.server.start_server $node_id
}

mpcserver() {
  go run $go_code_path/server/server.go -config $config -id $node_id > /usr/src/hbswap/log/mpc_server_$node_id.log 2>&1
}

setup_data
httpserver & mpcserver

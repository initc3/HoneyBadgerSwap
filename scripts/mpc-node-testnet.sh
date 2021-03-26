#!/usr/bin/env bash
set -e

config=${1:-/opt/hbswap/conf/server.toml}

# Place the data where MP-SPDZ expects it
setup_data() {
    rm -rf /opt/hbswap/db
    mkdir -p Persistence Player-Data
    mkdir -p /opt/hbswap/db /opt/hbswap/inputmask-shares /opt/hbswap/preprocessing-data
    # Copy the public keys of all players
    cp /opt/hbswap/public-keys/* Player-Data/
    # Symlink to the private key, to where MP-SPDZ expects it to be (under Player-Data/).
    ln -s /run/secrets/P$NODE_ID.key Player-Data/P$NODE_ID.key
}


httpserver() {
    #python -m honeybadgerswap.server.main
    hbswap-start-httpserver
    #uvicorn honeybadgerswap.server.main:app --host 0.0.0.0 --port 8080
}

mpcserver() {
  ./mpcserver -config $config -id $NODE_ID > /usr/src/hbswap/log/mpc_server_$NODE_ID.log 2>&1
}

setup_data
httpserver & mpcserver

#!/usr/bin/env bash
set -e

###Command: bash Scripts/hbswap/scripts/run.sh

trade() {
  go run Scripts/hbswap/go/trade/trade.go $1 $2
}

httpserver() {
  python3 Scripts/hbswap/python/start_server.py $1
}

mpcserver() {
  go run Scripts/hbswap/go/server/server.go $1 > Scripts/hbswap/log/mpc_server_$1.log 2>&1
}

rm -rf Scripts/hbswap/poa/data
bash Scripts/hbswap/scripts/chain.sh

go run Scripts/hbswap/go/deploy/deploy.go

bash Scripts/hbswap/scripts/prep.sh

httpserver 0 &
httpserver 1 &
httpserver 2 &

mpcserver 0 &
mpcserver 1 &
mpcserver 2 &

trade 1.1 -2.5
trade -1.1 2






#!/usr/bin/env bash
set -e

deposit() {
  go run Scripts/hbswap/go/client/deposit.go $1 $2 $3
}

withdraw() {
  go run Scripts/hbswap/go/client/withdraw.go $1 $2 $3
}

secret_deposit() {
  go run Scripts/hbswap/go/client/secret_deposit.go $1 $2 $3
}

secret_withdraw() {
  go run Scripts/hbswap/go/client/secret_withdraw.go $1 $2 $3
}

trade() {
  go run Scripts/hbswap/go/client/trade.go $1 $2 $3
}

httpserver() {
  python3 Scripts/hbswap/python/server/start_server.py $1
}

mpcserver() {
  go run Scripts/hbswap/go/server/server.go $1 > Scripts/hbswap/log/mpc_server_$1.log 2>&1
}

pkill -f geth | true
rm -rf Scripts/hbswap/poa/data
bash Scripts/hbswap/scripts/chain.sh

go run Scripts/hbswap/go/deploy/deploy.go

deposit 0 10 10

bash Scripts/hbswap/scripts/prep.sh

httpserver 0 &
httpserver 1 &
httpserver 2 &
httpserver 3 &

mpcserver 0 &
mpcserver 1 &
mpcserver 2 &
mpcserver 3 &

secret_deposit 0 10 10

trade 0 1.1 -2.5
#trade 0 -1.1 2
#
#secret_withdraw 0 9 9
#
#withdraw 0 9 9






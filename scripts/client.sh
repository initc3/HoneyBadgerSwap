#!/usr/bin/env bash
set -e

eth_host=${1:-localhost}
network=${2:-privatenet}

deposit() {
  go run Scripts/hbswap/go/client/deposit.go -n $network $1 $2 $3 $eth_host
}

withdraw() {
  go run Scripts/hbswap/go/client/withdraw.go -n $network $1 $2 $3 $eth_host
}

secret_deposit() {
  go run Scripts/hbswap/go/client/secret_deposit.go -n $network $1 $2 $3 $eth_host
}

secret_withdraw() {
  go run Scripts/hbswap/go/client/secret_withdraw.go -n $network $1 $2 $3 $eth_host
}

trade() {
  go run Scripts/hbswap/go/client/trade.go -n $network $1 $2 $3 $eth_host
}

secret_deposit 0 10 10

trade 0 1.1 -2.5
trade 0 -1.1 2

secret_withdraw 0 9 9

withdraw 0 9 9

#!/usr/bin/env bash

set -e

eth_host=${1:-localhost}
network=${2:-privatenet}
go_code_path=/go/src/github.com/initc3/HoneyBadgerSwap/src/go

. eth-data.sh

init_pool() {
  go run $go_code_path/client/deposit.go -n $network $1 $2 $3 $4 $5 $eth_host
  #go run $go_code_path/client/secret_deposit.go -n $network $1 $2 $3 $4 $5 $eth_host
  #go run $go_code_path/client/init_pool.go -n $network $1 $2 $3 $4 $5 $eth_host
}

init_pool 0 $eth $token_1 10 20

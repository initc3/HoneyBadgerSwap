#!/usr/bin/env bash

set -e

eth_host=${1:-localhost}

. scripts/eth-data.sh

init_pool() {
  go run go/client/deposit.go $1 $2 $3 $4 $5 $eth_host
  go run go/client/secret_deposit.go $1 $2 $3 $4 $5 $eth_host
  go run go/client/init_pool.go $1 $2 $3 $4 $5 $eth_host
}

init_pool 0 $eth $token_1 10 20
#init_pool 0 $eth $token_2 10 20

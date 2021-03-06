#!/usr/bin/env bash

set -e

config=${1:-/opt/hbswap/conf/client.toml}
go_code_path=/go/src/github.com/initc3/HoneyBadgerSwap/src/go

. eth-data.sh

init_pool() {
  go run $go_code_path/client/deposit.go -config $config $1 $2 $3 $4 $5
  go run $go_code_path/client/secret_deposit.go -config $config $1 $2 $3 $4 $5
  go run $go_code_path/client/init_pool.go -config $config $1 $2 $3 $4 $5
}

init_pool 0 $eth $token_1 10 20

#!/usr/bin/env bash

set -e

config=${1:-/opt/hbswap/conf/client.toml}
go_code_path=/go/src/github.com/initc3/HoneyBadgerSwap/src/go

. eth-data.sh

deposit() {
  go run $go_code_path/client/secret_deposit.go -config $config $1 $2 $3 $4 $5
}

deposit 0 $eth $token_1 10 10

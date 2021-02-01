#!/usr/bin/env bash

set -e

eth_host=${1:-localhost}
go_code_path=/go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap/go

. eth-data.sh

deposit() {
  go run $go_code_path/client/secret_deposit.go $1 $2 $3 $4 $5 $eth_host
}

deposit 0 $eth $token_1 10 10

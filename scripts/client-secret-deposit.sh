#!/usr/bin/env bash

set -e

eth_host=${1:-localhost}

. scripts/eth-data.sh

deposit() {
  go run go/client/secret_deposit.go $1 $2 $3 $4 $5 $eth_host
}

deposit 0 $eth $token_1 10 10

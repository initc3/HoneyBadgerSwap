#!/usr/bin/env bash
set -e

config=${4:-/opt/hbswap/conf/client.toml}
go_code_path=/go/src/github.com/initc3/HoneyBadgerSwap/src/go
#user=$2
#amt_A=$3
#amt_B=$4


. utils.sh

trade() {
    go run $go_code_path/client/trade.go -config $config $1 $2 $3 $4 $5
}

#trade 0 $eth $token_1 1.1 -2.5
#trade $user $eth $token_1 $amt_A $amt_B

trade 0 $eth $token_1 -1.1 2

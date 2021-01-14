#!/usr/bin/env bash
set -e

eth_host=${1:-localhost}
#user=$2
#amt_A=$3
#amt_B=$4


. Scripts/hbswap/scripts/utils.sh

trade() {
    go run Scripts/hbswap/go/client/trade.go $1 $2 $3 $4 $5 $eth_host
}

trade 0 $eth $token_1 1.1 -2.5
#trade $user $eth $token_1 $amt_A $amt_B

#trade 0 $eth $token_1 -1.1 2

#!/usr/bin/env bash
set -e

compile() {
  ./compile.py -v -C -F 256 $1
}

rm -rf Programs/Bytecode Programs/Public-Input Programs/Schedules

compile hbswap_add_liquidity
compile hbswap_calc_price
compile hbswap_check_balance
compile hbswap_remove_liquidity
compile hbswap_trade

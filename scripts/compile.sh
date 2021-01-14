#!/usr/bin/env bash
set -e

compile() {
  ./compile.py -v -C -F 256 $1
}

rm -rf Programs/Bytecode Programs/Public-Input Programs/Schedules

compile hbswap_trade
compile hbswap_withdraw
compile hbswap_calc_price

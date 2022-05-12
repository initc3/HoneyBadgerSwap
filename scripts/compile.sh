#!/usr/bin/env bash
set -e

compile() {
  ./compile.py -v -C -F 256 $1
}

rm -rf Programs/Bytecode/* Programs/Public-Input Programs/Schedules/*

compile hbswap_secret_withdraw
compile hbswap_init_pool
compile hbswap_add_liquidity
compile hbswap_remove_liquidity
compile hbswap_trade
compile hbswap_calc_individual_price
compile hbswap_calc_batch_price
#!/usr/bin/env bash
set -e

. Scripts/hbswap/scripts/utils.sh

prepare() {
  rm -rf Player-Data

  rm -rf Persistence
  mkdir Persistence

  rm -rf Scripts/hbswap/db
  mkdir Scripts/hbswap/db

  Scripts/setup-ssl.sh $players
}

compile() {
  ./compile.py -v -C -F 256 $1
}

prepare

compile hbswap_add_liquidity
compile hbswap_calc_price
compile hbswap_check_balance
compile hbswap_remove_liquidity
compile hbswap_trade
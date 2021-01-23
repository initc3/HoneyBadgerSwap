#!/usr/bin/env bash
set -e

db_dir=${DB_PATH:-/opt/hbswap/db}

. Scripts/hbswap/scripts/utils.sh

prepare() {
  rm -rf Player-Data

  rm -rf Persistence
  mkdir Persistence

  rm -rf $db_dir
  mkdir $db_dir

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

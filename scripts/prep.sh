#!/usr/bin/env bash
set -e

. Scripts/hbswap/scripts/utils.sh

prepare() {
  rm -rf Persistence
  rm -rf Scripts/hbswap/db
  rm -rf Player-Data

  mkdir Persistence
  mkdir Scripts/hbswap/db

  Scripts/setup-ssl.sh $players
}

compile() {
  ./compile.py -v -C -F 256 $1
}

prepare

compile hbswap_trade
compile hbswap_withdraw
compile hbswap_calc_price
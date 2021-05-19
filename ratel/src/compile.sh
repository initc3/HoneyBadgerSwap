#!/usr/bin/env bash
set -e

compile_flex() {
  flex ../src/scanner/$1.l
  g++ -o $1 lex.yy.c
}

compile_flexes() {
  compile_flex split_public_private
  compile_flex org_sol
  compile_flex split_private
}

compile_sol() {
  truffle compile
}

compile_mpc() {
  cd mpc
  for d in *; do
    mkdir -p ../../../Programs/Source
    cp $d ../../../Programs/Source/$d
    ./../../../compile.py -v -C -F 256 $d
  done
  cd ..
}

run_scanner() {
  ./$1 $2 <$3
}

parse() {
  # split public and private code
  run_scanner split_public_private $1 ../src/rl/$1.rl

  # org public code
  run_scanner org_sol $1 contracts/$1.sol
  mv contracts/tmp.sol contracts/$1.sol

  # split python and MP-SPDZ code
  run_scanner split_private $1 mpc/$1.mpc
  rm mpc/$1.mpc
}

mkdir -p ratel/genfiles/contracts
mkdir -p ratel/genfiles/python
mkdir -p ratel/genfiles/mpc
cd ratel/genfiles

compile_flexes

parse test

compile_sol
compile_mpc

cd ../..

pkill -f geth || true
pkill -f python3 || true
rm -rf /opt/hbswap/db/*

bash chain-latest.sh &
sleep 3

bash setup-ssl.sh 4

python3 -m ratel.src.python.run 0 &
python3 -m ratel.src.python.run 1 &
python3 -m ratel.src.python.run 2 &
python3 -m ratel.src.python.run 3 &

python3 ratel/src/python/deploy.py
#!/usr/bin/env bash
set -e

compile_flex() {
  flex ../src/scanner/$1.l
  g++ -o $1 lex.yy.c
}

compile_flexes() {
  compile_flex split_public_private
  compile_flex split_public
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

  # reorg sol code
  run_scanner split_public $1 contracts/$1.sol
  mv contracts/tmp.sol contracts/$1.sol

  # split python and MP-SPDZ code
  run_scanner split_private $1 mpc/$1.mpc
  rm mpc/$1.mpc
}

#bash setup-ssl.sh 4
#mkdir -p ratel/genfiles/contracts
#mkdir -p ratel/genfiles/python
#mkdir -p ratel/genfiles/mpc

rm ratel/genfiles/mpc/* || true
rm Programs/Source/* || true
rm ratel/genfiles/contracts/* || true

cd ratel/genfiles

compile_flexes

parse hbswap
parse volumeMatching
parse rockPaperScissors
parse fabcar
parse review

compile_sol
compile_mpc

# bash ratel/src/compile.sh


#!/usr/bin/env bash
set -e


compile_flex() {
  flex ../src/scanner/$1.l
  g++ -o $1 lex.yy.c -ll
}

compile() {
  compile_flex split_public_private
  compile_flex org_sol
  compile_flex split_private
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
  truffle compile

  # split python and MP-SPDZ code
  run_scanner split_private $1 mpc/$1.mpc
  mv mpc/tmp.mpc mpc/$1.mpc
  cp mpc/$1.mpc ../../MP-SPDZ/Programs/Source/$1.mpc
  ./../../MP-SPDZ/compile.py $1
  rm ../../MP-SPDZ/Programs/Source/$1.mpc
}

mkdir -p ratel/genfiles/contracts
mkdir -p ratel/genfiles/python
mkdir -p ratel/genfiles/mpc
cd ratel/genfiles
compile

parse test


#!/usr/bin/env bash
set -e


#### parameters
apps=$1
finalize_on_chain=$2
is_benchmark=$3
####


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
    ./../../../compile.py -v -C -F 128 $d
  done
  cd ..
}


parse() {
  # split public and private code
  ./split_public_private $1 < ../src/rl/$1.rl

  # reorg sol code
  ./split_public $1 $$finalize_on_chain < contracts/$1.sol
  mv contracts/tmp.sol contracts/$1.sol

  # split python and MP-SPDZ code
  ./split_private $1 $finalize_on_chain $is_benchmark < mpc/$1.mpc
  rm mpc/$1.mpc
}


init() {
  rm -rf ratel/genfiles
  mkdir -p ratel/genfiles
  mkdir -p ratel/genfiles/python
  mkdir -p ../../../Programs/Source
  cp -r ratel/src/mpc ratel/genfiles/
  cp -r ratel/src/contracts ratel/genfiles/
  cp -r ratel/src/truffle-config.js ratel/genfiles
  cp -r ratel/src/node_modules ratel/genfiles
}


init

cd ratel/genfiles

#### compile ratel compiler
compile_flexes
####

#### compile application
IFS=','
read -a strarr <<< "$apps"
for app in "${strarr[@]}";
do
  parse $app
done

compile_sol
compile_mpc
####


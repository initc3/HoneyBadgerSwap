set -e

truffle_complie() {
    rm -rf build/
    truffle compile
}

extract_abi_bin() {
  jq .abi build/contracts/$1.json > genfiles/$1.abi
  jq -r .bytecode build/contracts/$1.json > genfiles/$1.bin
}

update_genfiles() {
  mkdir -p genfiles

  extract_abi_bin HbSwap
}

abigen_files() {
  INPUT_DIR=genfiles

  OUTPUT_DIR=gobingdings
  mkdir -p $OUTPUT_DIR/$2
  abigen -abi $INPUT_DIR/$1.abi -bin $INPUT_DIR/$1.bin -pkg $2 -type $1 -out $OUTPUT_DIR/$2/$2.go
}

sync_go_binding() {
  abigen_files HbSwap hbswap
}

cd Scripts/hbswap
truffle_complie
update_genfiles
sync_go_binding
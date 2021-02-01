#!/usr/bin/env bash
set -e

script=${1:-"Scripts/setup-ssl.sh"}
nparties=${2:-4}

rm -f Player-Data/*
bash $script $nparties

for ((i=0; i < $nparties; i++))
do
    mkdir -p Secrets-P$i
    mv Player-Data/P$i.key Secrets-P$i/
done

#!/usr/bin/env bash

./random-shamir.x -i 0 -N 4 -T 1 --nshares 1000 &
./random-shamir.x -i 1 -N 4 -T 1 --nshares 1000 &
./random-shamir.x -i 2 -N 4 -T 1 --nshares 1000 &
./random-shamir.x -i 3 -N 4 -T 1 --nshares 1000

python3 Scripts/hbswap/python/test/check.py
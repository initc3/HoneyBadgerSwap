#!/usr/bin/env bash

n=${1:-3}

test -e public || mkdir -p public
test -e secret || mkdir -p secret

publicdir=public
secretdir=secret

echo Setting up SSL for $n parties

for i in `seq 0 $[n-1]`; do
    openssl req -newkey rsa -nodes -x509 -out $publicdir/P$i.pem -keyout $secretdir/P$i.key -subj "/CN=P$i"
done

c_rehash public

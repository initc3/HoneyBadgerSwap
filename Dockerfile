# Go (server) dependencies
FROM golang:buster as go-deps

# ethereum
COPY --from=sbellem/go-ethereum:cfbb969da-buster \
                /go/src/golang.org/x /go/src/golang.org/x
COPY --from=sbellem/go-ethereum:cfbb969da-buster \
                /go/src/github.com/ethereum /go/src/github.com/ethereum

COPY src /go/src/github.com/initc3/HoneyBadgerSwap/src

WORKDIR /go/src/github.com/initc3/HoneyBadgerSwap/src
RUN go get -d -v ./...

# needed to deploy contracts
# TODO: verify whether poa dir is really needed, or what is needed from it, maybe
# the keystore is sufficient
COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it
COPY poa/keystore /opt/poa/keystore


# main image
FROM python:3.8

ENV PYTHONUNBUFFERED 1

# MP-SPDZ
# TODO: review dependencies as some of them may not be needed.
RUN apt-get update && apt-get install -y --no-install-recommends \
                build-essential \
                libboost-dev \
                libboost-thread-dev \
                libsodium-dev \
                libssl-dev \
                libtool \
                m4 \
                texinfo \
                yasm \
        && rm -rf /var/lib/apt/lists/*
# mpir
ENV LD_LIBRARY_PATH /usr/local/lib
RUN mkdir -p /usr/local/share/info
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/lib/libmpir*.*a /usr/local/lib/
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/lib/libmpir.so.23.0.3 /usr/local/lib/
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/lib/libmpirxx.so.8.4.3 /usr/local/lib/
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/include/mpir*.h /usr/local/include/
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/share/info/* /usr/local/share/info/
RUN set -ex \
    && cd /usr/local/lib \
    && ln -s libmpir.so.23.0.3 libmpir.so \
    && ln -s libmpir.so.23.0.3 libmpir.so.23 \
    && ln -s libmpirxx.so.8.4.3 libmpirxx.so \
    && ln -s libmpirxx.so.8.4.3 libmpirxx.so.8

ENV HBSWAP_HOME /usr/src/hbswap
WORKDIR $HBSWAP_HOME
ENV INPUTMASK_SHARES "/opt/hbswap/inputmask-shares"
ENV PREP_DIR "/opt/hbswap/preprocessing-data"
COPY --from=sbellem/mpspdz:2cf97686-randomshamirprep\
                /usr/src/MP-SPDZ/random-shamir.x /usr/src/hbswap/
COPY --from=sbellem/mpspdz:2cf97686-maliciousshamirparty \
                /usr/src/MP-SPDZ/malicious-shamir-party.x /usr/src/hbswap/
RUN mkdir -p $INPUTMASK_SHARES $PREP_DIR
COPY testkeys/public /opt/hbswap/public-keys
#############################################################################

# GO (server) dependencies
ENV PATH /usr/local/go/bin:$PATH
COPY --from=golang:1.15.8-buster /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY --from=go-deps /go/src /go/src
#############################################################################

ENV DB_PATH /opt/hbswap/db

# Python (HTTP server) dependencies for HTTP server
RUN apt-get update && apt-get install -y --no-install-recommends \
                lsof \
                libmpfr-dev \
                libmpc-dev \
        && rm -rf /var/lib/apt/lists/*

COPY src/python /usr/src/honeybadgerswap-python
WORKDIR /usr/src/honeybadgerswap-python
RUN pip install --editable .

### In development contexts, these files can be mounted along with the src code
ARG http_server_config=conf/server.toml
COPY $http_server_config /opt/hbswap/conf/server.toml
COPY scripts/mpc-node.sh /usr/src/hbswap/mpc-node.sh
COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it
COPY poa/keystore /opt/poa/keystore

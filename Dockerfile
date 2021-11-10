# Go (server) dependencies
FROM golang:buster as go-deps

# ethereum
COPY --from=sbellem/go-ethereum:cfbb969da-buster \
                /go/src/golang.org/x /go/src/golang.org/x
#COPY --from=sbellem/go-ethereum:cfbb969da-buster \
#                /go/src/github.com/ethereum /go/src/github.com/ethereum
WORKDIR /go/src/github.com/ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git

COPY src /go/src/github.com/initc3/HoneyBadgerSwap/src

#WORKDIR /go/src/github.com/initc3/HoneyBadgerSwap/src
#RUN go get -d -v ./...

# needed to deploy contracts
# TODO: verify whether poa dir is really needed, or what is needed from it, maybe
# the keystore is sufficient
#COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it
#COPY poa/keystore /opt/poa/keystore


# MPC program compilation to bytecodes
FROM python:3.8-buster as mpc-bytecodes

ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src
COPY MP-SPDZ/compile.py .
COPY MP-SPDZ/Compiler Compiler
RUN mkdir -p Programs/Source
COPY src/mpc Programs/Source
COPY scripts/compile.sh .
RUN bash compile.sh


# main image
FROM python:3.8-buster

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

ENV DB_PATH /opt/hbswap/db

# GO (server) dependencies
ENV PATH /usr/local/go/bin:$PATH
COPY --from=golang:1.13.12 /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY --from=go-deps /go/src /go/src
#RUN go build -o $HBSWAP_HOME/mpcserver /go/src/github.com/initc3/HoneyBadgerSwap/src/go/server/server.go

# Python (HTTP server) dependencies for HTTP server
RUN apt-get update && apt-get install -y --no-install-recommends \
                lsof \
                libmpfr-dev \
                libmpc-dev \
        && rm -rf /var/lib/apt/lists/*

COPY src/python /usr/src/honeybadgerswap-python
WORKDIR /usr/src/honeybadgerswap-python
RUN pip install --editable .[dev,test]

### In development contexts, these files can be mounted along with the src code
ARG http_server_config=conf/server.toml
COPY $http_server_config /opt/hbswap/conf/server.toml
COPY scripts/mpc-node.sh /usr/src/hbswap/mpc-node.sh
COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it
COPY poa/keystore /opt/poa/keystore

# MPC bytecodes and schedules -- from the compilation stage
WORKDIR $HBSWAP_HOME
RUN mkdir -p Programs
COPY --from=mpc-bytecodes /usr/src/Programs/Bytecode /usr/src/hbswap/Programs/Bytecode
COPY --from=mpc-bytecodes /usr/src/Programs/Schedules /usr/src/hbswap/Programs/Schedules

RUN apt-get update && apt-get install -y --no-install-recommends \
    nodejs npm
RUN npm install -g truffle

RUN apt-get update && apt-get install -y --no-install-recommends \
    flex \
    vim

WORKDIR /go/src/github.com/ethereum/go-ethereum
RUN make geth

RUN pip3 install web3
RUN pip3 install matplotlib

RUN mkdir -p /opt/hbswap/db
RUN mkdir -p /usr/src/hbswap/Persistence

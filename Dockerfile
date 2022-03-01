# Go (server) dependencies
FROM golang:1.17.5-bullseye as go-deps

# ethereum
WORKDIR /go/src/github.com/ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git

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

ENV HBSWAP_HOME /usr/src/hbswap
WORKDIR $HBSWAP_HOME
RUN mkdir -p /usr/src/MP-SPDZ
ENV INPUTMASK_SHARES "/opt/hbswap/inputmask-shares"
ENV PREP_DIR "/opt/hbswap/preprocessing-data"

COPY --from=sbellem/mpspdz:shamirshares-2b3b7076 \
                /usr/src/MP-SPDZ/random-shamir.x /usr/src/hbswap/
COPY --from=sbellem/mpspdz:malshamirparty-2b3b7076 \
                /usr/src/MP-SPDZ/malicious-shamir-party.x /usr/src/hbswap/
COPY --from=sbellem/mpspdz:malshamiroffline-2b3b7076 \
                /usr/src/MP-SPDZ/mal-shamir-offline.x /usr/src/hbswap/

COPY --from=sbellem/mpspdz:shamirshares-2b3b7076 \
                /usr/src/MP-SPDZ/libSPDZ.so /usr/src/MP-SPDZ/

COPY --from=sbellem/mpspdz:shamirshares-2b3b7076 \
                /usr/src/MP-SPDZ/local /usr/src/MP-SPDZ/local

RUN mkdir -p $INPUTMASK_SHARES $PREP_DIR
COPY MP-SPDZ/Scripts/setup-ssl.sh .
RUN ./setup-ssl.sh 4
#############################################################################

ENV DB_PATH /opt/hbswap/db

# GO (server) dependencies
ENV PATH /usr/local/go/bin:$PATH
COPY --from=golang:1.17.5-bullseye /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY --from=go-deps /go/src /go/src
#RUN go build -o $HBSWAP_HOME/mpcserver /go/src/github.com/initc3/HoneyBadgerSwap/src/go/server/server.go

# FIXME is this needed?
# Python (HTTP server) dependencies for HTTP server
RUN apt-get update && apt-get install -y --no-install-recommends \
                lsof \
                libmpfr-dev \
                libmpc-dev \
        && rm -rf /var/lib/apt/lists/*

# FIXME is this needed?
COPY src/python /usr/src/honeybadgerswap-python
WORKDIR /usr/src/honeybadgerswap-python
RUN pip install --editable .[dev,test]

# FIXME is this needed?
### In development contexts, these files can be mounted along with the src code
ARG http_server_config=conf/server.toml
COPY $http_server_config /opt/hbswap/conf/server.toml
COPY scripts/mpc-node.sh /usr/src/hbswap/mpc-node.sh
COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it
COPY poa/keystore /opt/poa/keystore

# FIXME speed up by building in separate stage and just copying needed files if possible
WORKDIR /go/src/github.com/ethereum/go-ethereum
RUN make geth

RUN apt-get update && apt-get install -y --no-install-recommends \
                flex \
                vim \
                nodejs \
                npm \
        && rm -rf /var/lib/apt/lists/*

RUN npm install -g truffle@5.4.29

RUN pip3 install web3==5.24.0 matplotlib

RUN mkdir -p /opt/hbswap/db
RUN mkdir -p /usr/src/hbswap/Persistence

WORKDIR $HBSWAP_HOME

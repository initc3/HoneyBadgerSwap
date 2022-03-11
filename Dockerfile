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
ENV INPUTMASK_SHARES "/opt/inputmask-shares"
ENV PREP_DIR "/opt/preprocessing-data"

# malicious-shamir-party.x
COPY --from=initc3/malicious-shamir-party.x:009d9910 \
                /usr/src/MP-SPDZ/malicious-shamir-party.x \
                /usr/local/bin/malicious-shamir-party.x
#COPY --from=initc3/malicious-shamir-party.x:009d9910 \
#                /usr/lib/x86_64-linux-gnu/libboost_system.so.1.67.0 \
#                /usr/lib/x86_64-linux-gnu/libboost_system.so.1.67.0
#                /usr/lib/x86_64-linux-gnu/libboost_system.so.1.74.0 \
#                /usr/lib/x86_64-linux-gnu/libboost_system.so.1.74.0
COPY --from=initc3/malicious-shamir-party.x:009d9910 \
                /usr/src/MP-SPDZ/libSPDZ.so /usr/src/MP-SPDZ/
COPY --from=initc3/malicious-shamir-party.x:009d9910 \
                /usr/src/MP-SPDZ/local /usr/src/MP-SPDZ/local
RUN cp /usr/local/bin/malicious-shamir-party.x /usr/src/hbswap/

# mal-shamir-offline.x
COPY --from=initc3/mal-shamir-offline.x:009d9910 \
                /usr/src/MP-SPDZ/mal-shamir-offline.x /usr/local/bin/
RUN cp /usr/local/bin/mal-shamir-offline.x /usr/src/hbswap/

# random-shamir.x
COPY --from=initc3/random-shamir.x:009d9910 \
                /usr/src/MP-SPDZ/random-shamir.x /usr/local/bin/
RUN cp /usr/local/bin/random-shamir.x /usr/src/hbswap/

# MP-SPDZ compiler
COPY --from=initc3/mpspdz:009d9910 \
                /usr/src/MP-SPDZ/compile.py /usr/src/hbswap/
COPY --from=initc3/mpspdz:009d9910 \
                /usr/src/MP-SPDZ/Compiler /usr/src/hbswap/Compiler
# ssl keys
COPY --from=initc3/mpspdz:009d9910 \
                /usr/src/MP-SPDZ/Scripts/setup-ssl.sh /usr/src/hbswap/

RUN mkdir -p $INPUTMASK_SHARES $PREP_DIR
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
                iproute2 \
        && rm -rf /var/lib/apt/lists/*

RUN npm install -g truffle@5.4.29

RUN pip3 install web3==5.24.0 matplotlib

RUN mkdir -p /opt/hbswap/db
RUN mkdir -p /usr/src/hbswap/Persistence

WORKDIR $HBSWAP_HOME

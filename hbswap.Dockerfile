FROM python:3.8 as base-mp-spdz

RUN apt-get update && apt-get install -y --no-install-recommends \
                automake \
                build-essential \
                git \
                libboost-dev \
                libboost-thread-dev \
                libsodium-dev \
                libssl-dev \
                libtool \
                m4 \
                texinfo \
                yasm \
                vim \
                gdb \
                valgrind \
        && rm -rf /var/lib/apt/lists/*

ENV MP_SPDZ_HOME /usr/src/MP-SPDZ
WORKDIR $MP_SPDZ_HOME

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

COPY Makefile .
COPY CONFIG .
COPY BMR BMR
#COPY ECDSA ECDSA
COPY Exceptions Exceptions
#COPY ExternalIO ExternalIO
#COPY FHE FHE
#COPY FHEOffline FHEOffline
COPY GC GC
COPY Machines Machines
COPY Math Math
COPY Networking Networking
COPY OT OT
COPY Processor Processor
COPY Protocols Protocols
COPY SimpleOT SimpleOT
COPY Tools Tools
COPY Utils Utils
#COPY Yao Yao

RUN make clean

# DEBUG and configuration flags
RUN echo "MY_CFLAGS += -DDEBUG_NETWORKING" >> CONFIG.mine \
        && echo "MY_CFLAGS += -DVERBOSE" >> CONFIG.mine \
        && echo "MY_CFLAGS += -DDEBUG_MAC" >> CONFIG.mine \
        && echo "MY_CFLAGS += -DDEBUG_FILE" >> CONFIG.mine \
        && echo "MOD = -DGFP_MOD_SZ=4" >> CONFIG.mine

ENV PRIME 52435875175126190479447740508185965837690552500527637822603658699938581184513
ENV N_PARTIES 4
ENV THRESHOLD 1

# Compile random-shamir
FROM base-mp-spdz as inputmask-preprocessing
ENV INPUTMASK_SHARES "/opt/hbswap/inputmask-shares"
RUN mkdir -p $INPUTMASK_SHARES \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/inputmask-shares/\"'" >> CONFIG.mine
RUN make random-shamir.x

# Compile malicious-shamir-party
FROM base-mp-spdz as malicious-shamir
ENV PREP_DIR "/opt/hbswap/preprocessing-data"
RUN mkdir -p $PREP_DIR \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/preprocessing-data/\"'" >> CONFIG.mine
RUN make malicious-shamir-party.x

########################## end of mp-spdz builds #######################################

FROM base-mp-spdz as hbswap

# GO (server) dependencies
ENV PATH /usr/local/go/bin:$PATH
COPY --from=golang:1.15.6-buster /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go get -d -v github.com/ethereum/go-ethereum

WORKDIR $GOPATH/src/github.com/ethereum/go-ethereum
RUN git checkout cfbb969da

COPY Scripts/hbswap/src /go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap

WORKDIR /go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap

RUN go get -d -v ./...

COPY Scripts /usr/src/MP-SPDZ/Scripts

# MP-SPDZ
WORKDIR $MP_SPDZ_HOME
ENV INPUTMASK_SHARES "/opt/hbswap/inputmask-shares"
ENV PREP_DIR "/opt/hbswap/preprocessing-data"
COPY --from=inputmask-preprocessing /usr/src/MP-SPDZ/random-shamir.x /usr/src/MP-SPDZ/
COPY --from=malicious-shamir /usr/src/MP-SPDZ/malicious-shamir-party.x /usr/src/MP-SPDZ/
RUN mkdir -p $INPUTMASK_SHARES
RUN mkdir -p $PREP_DIR

ENV DB_PATH /opt/hbswap/db

# Python (HTTP server) dependencies for HTTP server
RUN apt-get update && apt-get install -y --no-install-recommends \
                lsof \
                libmpfr-dev \
                libmpc-dev \
        && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/MP-SPDZ/Scripts/hbswap/src/python
RUN pip install --editable .

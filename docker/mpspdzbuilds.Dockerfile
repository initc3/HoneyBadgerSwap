FROM python:3.8-buster as compiler

ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src
COPY MP-SPDZ/compile.py .
COPY MP-SPDZ/Compiler Compiler

RUN mkdir -p Programs/Source
###############################################################################

FROM python:3.8-buster as machines

ENV PYTHONUNBUFFERED 1

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

COPY MP-SPDZ/Makefile .
COPY MP-SPDZ/CONFIG .
COPY MP-SPDZ/BMR BMR
COPY MP-SPDZ/Exceptions Exceptions
COPY MP-SPDZ/GC GC
COPY MP-SPDZ/Machines Machines
COPY MP-SPDZ/Math Math
COPY MP-SPDZ/Networking Networking
COPY MP-SPDZ/OT OT
COPY MP-SPDZ/Processor Processor
COPY MP-SPDZ/Protocols Protocols
COPY MP-SPDZ/SimpleOT SimpleOT
COPY MP-SPDZ/Tools Tools
COPY MP-SPDZ/Utils Utils

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
FROM machines as random-shamir-prep
ENV INPUTMASK_SHARES "/opt/hbswap/inputmask-shares"
RUN mkdir -p $INPUTMASK_SHARES \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/inputmask-shares/\"'" >> CONFIG.mine
RUN make random-shamir.x

# Compile malicious-shamir-party
FROM machines as malicious-shamir-party
ENV PREP_DIR "/opt/hbswap/preprocessing-data"
RUN mkdir -p $PREP_DIR \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/preprocessing-data/\"'" >> CONFIG.mine
RUN make malicious-shamir-party.x

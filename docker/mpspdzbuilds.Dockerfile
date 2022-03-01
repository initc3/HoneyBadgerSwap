FROM python:3.8 as compiler

ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src
COPY MP-SPDZ/compile.py .
COPY MP-SPDZ/Compiler Compiler

RUN mkdir -p Programs/Source
###############################################################################

FROM python:3.8 as machines

ENV PYTHONUNBUFFERED 1

RUN apt-get update && apt-get install -y --no-install-recommends \
                automake \
                build-essential \
                git \
                libboost-dev \
                libboost-thread-dev \
                libntl-dev \
                libsodium-dev \
                libssl-dev \
                libtool \
                m4 \
                texinfo \
                yasm \
        && rm -rf /var/lib/apt/lists/*

ENV MP_SPDZ_HOME /usr/src/MP-SPDZ
WORKDIR $MP_SPDZ_HOME

RUN git clone --branch \
            random-shamir-prep-upgrade https://github.com/initc3/MP-SPDZ.git \
            $MP_SPDZ_HOME

RUN echo "USE_NTL = 1" >> CONFIG.mine

RUN make clean && make mpir

# DEBUG and configuration flags
RUN echo "MY_CFLAGS += -DDEBUG_NETWORKING" >> CONFIG.mine \
        && echo "MY_CFLAGS += -DVERBOSE" >> CONFIG.mine \
        && echo "MY_CFLAGS += -DDEBUG_MAC" >> CONFIG.mine \
        && echo "MY_CFLAGS += -DDEBUG_FILE" >> CONFIG.mine \
        && echo "MOD = -DGFP_MOD_SZ=4" >> CONFIG.mine

#ENV PRIME 52435875175126190479447740508185965837690552500527637822603658699938581184513
#ENV N_PARTIES 4
#ENV THRESHOLD 1

# Compile random-shamir
FROM machines as random-shamir-shares
ENV INPUTMASK_SHARES "/opt/hbswap/inputmask-shares"
RUN mkdir -p $INPUTMASK_SHARES \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/inputmask-shares/\"'" >> CONFIG.mine
RUN make random-shamir.x
RUN ./Scripts/setup-ssl.sh 4

# Compile malicious-shamir-party
FROM machines as malicious-shamir-party
ENV PREP_DIR "/opt/hbswap/preprocessing-data"
RUN mkdir -p $PREP_DIR \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/preprocessing-data/\"'" >> CONFIG.mine
RUN make malicious-shamir-party.x
RUN ./Scripts/setup-ssl.sh 4

# Compile mal-shamir-offline
FROM machines as mal-shamir-offline
ENV PREP_DIR "/opt/hbswap/preprocessing-data"
RUN mkdir -p $PREP_DIR \
        && echo "PREP_DIR = '-DPREP_DIR=\"/opt/hbswap/preprocessing-data/\"'" >> CONFIG.mine
RUN make mal-shamir-offline.x
RUN ./Scripts/setup-ssl.sh 4

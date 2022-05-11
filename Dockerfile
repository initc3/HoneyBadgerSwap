ARG mpspdz_commit=2bb036d
FROM initc3/malicious-shamir-party.x:${mpspdz_commit} as malshamirparty
FROM initc3/mal-shamir-offline.x:${mpspdz_commit} as malshamiroffline
FROM initc3/random-shamir.x:${mpspdz_commit} as randomshamir
FROM initc3/mpspdz:${mpspdz_commit} as mpspdzbase

FROM python:3.9.12-bullseye

ENV PYTHONUNBUFFERED 1

# MP-SPDZ
# TODO: review dependencies as some of them may not be needed.
RUN apt-get update && apt-get install -y --no-install-recommends \
                build-essential \
                libboost-dev \
                libboost-thread-dev \
                libntl-dev \
                libsodium-dev \
                libssl-dev \
                libtool \
                m4 \
                texinfo \
                yasm \
                # ratel
                flex \
                nodejs \
                npm \
                # dev
                vim \
                # tc command, for testing/benchmarking
                iproute2 \
                # FIXME is this needed?
                # Python (HTTP server) dependencies
                lsof \
                libmpfr-dev \
                libmpc-dev \
        && rm -rf /var/lib/apt/lists/*

# mpir
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/include/* /usr/local/include/
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/lib/* /usr/local/lib/
COPY --from=initc3/mpir:55fe6a9 /usr/local/mpir/share/info/* /usr/local/share/info/

ENV HBSWAP_HOME /usr/src/hbswap
ENV DB_PATH /opt/hbswap/db
ENV INPUTMASK_SHARES "/opt/inputmask-shares"
ENV PREP_DIR "/opt/preprocessing-data"

RUN mkdir -p \
            /usr/src/MP-SPDZ \
            ${HBSWAP_HOME} \
            ${HBSWAP_HOME}/Persistence \
            ${DB_PATH} \
            ${INPUTMASK_SHARES} \
            ${PREP_DIR}

# malicious-shamir-party.x
COPY --from=malshamirparty \
                /usr/local/bin/malicious-shamir-party.x \
                /usr/local/bin/malicious-shamir-party.x
COPY --from=malshamirparty /usr/src/MP-SPDZ/libSPDZ.so /usr/src/MP-SPDZ/
RUN cp /usr/local/bin/malicious-shamir-party.x /usr/src/hbswap/

# mal-shamir-offline.x
COPY --from=malshamiroffline \
                /usr/local/bin/mal-shamir-offline.x /usr/local/bin/
RUN cp /usr/local/bin/mal-shamir-offline.x /usr/src/hbswap/

# random-shamir.x
COPY --from=randomshamir /usr/local/bin/random-shamir.x /usr/local/bin/
RUN cp /usr/local/bin/random-shamir.x /usr/src/hbswap/

# MP-SPDZ compiler
COPY --from=mpspdzbase /usr/src/MP-SPDZ/compile.py /usr/src/hbswap/
COPY --from=mpspdzbase /usr/src/MP-SPDZ/Compiler /usr/src/hbswap/Compiler
COPY --from=mpspdzbase /usr/src/MP-SPDZ/Programs /usr/src/hbswap/Programs
# ssl keys
COPY --from=mpspdzbase /usr/src/MP-SPDZ/Scripts/setup-ssl.sh /usr/src/hbswap/

# geth
COPY --from=initc3/geth:97745ba /usr/local/bin/geth /usr/local/bin/geth
COPY poa/keystore /opt/poa/keystore

RUN npm install -g truffle@5.4.29

RUN pip3 install \
            aiohttp \
            aiohttp_cors \
            web3==5.24.0 \
            matplotlib \
            gmpy \
            gmpy2 \
            leveldb \
            toml \
            fastapi \
            pydantic \
            uvicorn[standard] \
            # dev
            ipython \
            ipdb

WORKDIR $HBSWAP_HOME

RUN ./setup-ssl.sh 4 /opt/ssl

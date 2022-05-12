FROM debian:bullseye-slim

ARG nplayers

RUN apt-get update && apt-get install -y --no-install-recommends openssl \
        && rm -rf /var/lib/apt/lists/*

COPY MP-SPDZ/Scripts/setup-ssl.sh /usr/src/setup-ssl.sh
RUN mkdir /usr/src/Player-Data

CMD ./setup-ssl.sh $nplayers

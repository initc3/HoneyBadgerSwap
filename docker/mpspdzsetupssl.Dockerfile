FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends openssl \
        && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src
COPY testkeys/genkeys.sh .
RUN mkdir public secret

ENTRYPOINT ["bash", "genkeys.sh"]

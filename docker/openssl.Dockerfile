FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends openssl \
        && rm -rf /var/lib/apt/lists/*

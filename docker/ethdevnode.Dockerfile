FROM ethereum/client-go:latest

ENV POADIR /opt/poa
ENV DATADIR /opt/poa/data
ENV KEYSTORE /opt/poa/keystore/server_0

WORKDIR /usr/src

COPY scripts/chain-latest.sh /usr/src/chain-latest.sh
COPY poa /opt/poa

ENTRYPOINT ["sh"]
CMD ["chain-latest.sh"]

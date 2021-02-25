FROM golang:1.15.8-buster

# TODO Work with go module mode
# SEE https://github.com/golang/go/wiki/Modules
# RUN go get -d -v github.com/ethereum/go-ethereum@cfbb969da
# RUN go get -d -v github.com/ethereum/go-ethereum
ARG commit=cfbb969da

RUN go get -d -v github.com/ethereum/go-ethereum

WORKDIR $GOPATH/src/github.com/ethereum/go-ethereum
RUN git checkout $commit

FROM golang:buster as builder

# TODO Work with go module mode
# SEE https://github.com/golang/go/wiki/Modules
# RUN go get -d -v github.com/ethereum/go-ethereum@cfbb969da
RUN go get -d -v github.com/ethereum/go-ethereum

WORKDIR $GOPATH/src/github.com/ethereum/go-ethereum
RUN git checkout cfbb969da

COPY src /go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap

WORKDIR /go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap

RUN go get -d -v ./...

WORKDIR /go/src/github.com/initc3/MP-SPDZ/Scripts/hbswap/go


FROM golang:buster as deploy
COPY --from=builder /go /go
ENTRYPOINT ["go", "run"]
CMD ["deploy/deploy.go", "eth.chain"]


FROM golang:buster as client
COPY --from=builder /go /go
#ENTRYPOINT ["go", "run"]
#CMD ["client/deposit.go", "0", "10", "10", "eth.chain"]

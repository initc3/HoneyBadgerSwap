package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go_bindings/token"
	"log"
	"math/big"
)

func Approve(conn *ethclient.Client, auth *bind.TransactOpts, receiver common.Address, amt *big.Int) {
	tokenInstance, err := token.NewToken(TokenAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := tokenInstance.Approve(auth, receiver, amt)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

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

func FundToken(conn *ethclient.Client, tokenAddr common.Address, toAddr common.Address, amount *big.Int) {
	balance := GetBalanceToken(conn, toAddr, tokenAddr)
	if balance.Cmp(amount) != -1 {
		log.Printf("Funded %s to %v token\n", toAddr.Hex(), balance)
		return
	}
	amount.Sub(amount, balance)

	adminAuth := GetAccount("server_0")
	transferToken(conn, tokenAddr, adminAuth, toAddr, amount)

	balance = GetBalanceToken(conn, toAddr, tokenAddr)
	log.Printf("Funded %s to %v token\n", toAddr.Hex(), balance)
}

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

func GetBalanceToken(conn *ethclient.Client, addr common.Address, tokenAddr common.Address) (*big.Int) {
	tokenInstance, err := token.NewToken(tokenAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := tokenInstance.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func transferToken(conn *ethclient.Client, tokenAddr common.Address, fromAuth *bind.TransactOpts, toAddr common.Address, value *big.Int) {
	fromAddr := fromAuth.From.Hex()

	log.Printf("Trasfering %v token from %s to %s\n", value, fromAddr, toAddr.Hex())

	tokenInstance, err := token.NewToken(tokenAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := tokenInstance.Transfer(fromAuth, toAddr, value)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %x", receipt.Status)
	}
}
package main

import (
	"context"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/gobingdings/hbswap"
)

func Deploy(conn *ethclient.Client, auth *bind.TransactOpts) (common.Address) {
	log.Println("Deploying HbSwap contract...")

	hbswapAddr, tx, _, err := hbswap.DeployHbSwap(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy HbSwap: %v", err)
	}

	receipt, err := utils.WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatalf("Failed to WaitMined HbSwap: %v", err)
	}
	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %x", receipt.Status)
	}

	log.Println("Deployed HbSwap contract at", hbswapAddr.Hex())

	return hbswapAddr
}

func main() {
	conn := utils.GetEthClient("HTTP://127.0.0.1:8545")

	owner, _ := utils.GetAccount("account_0")

	Deploy(conn, owner)
}
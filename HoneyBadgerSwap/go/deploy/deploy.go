package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go_bindings/hbswap"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go_bindings/token"
	"log"
)

const (
	n = 4
	t = 1
)

func DeployHbSwap(conn *ethclient.Client, auth *bind.TransactOpts) (common.Address) {
	log.Println("Deploying HbSwap contract...")

	//var servers []common.Address
	//for i := 0; i < n; i++ {
	//	transactOpt, _ := utils.GetAccount(fmt.Sprintf("server_%v", i))
	//	servers = append(servers, transactOpt.From)
	//}

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

func DeployToken(conn *ethclient.Client, auth *bind.TransactOpts) (common.Address) {
	log.Println("Deploying Token contract...")

	tokenAddr, tx, _, err := token.DeployToken(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy Token: %v", err)
	}

	receipt, err := utils.WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatalf("Failed to WaitMined Token: %v", err)
	}
	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %x", receipt.Status)
	}

	log.Println("Deployed Token contract at", tokenAddr.Hex())

	return tokenAddr
}

func main() {
	conn := utils.GetEthClient("HTTP://127.0.0.1:8545")

	owner, _ := utils.GetAccount("server_0")

	DeployHbSwap(conn, owner)
	DeployToken(conn, owner)
}
// go run src/go/deploy/deploy.go

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"github.com/initc3/HoneyBadgerSwap/src/go_bindings/hbswap"
	"github.com/initc3/HoneyBadgerSwap/src/go_bindings/hbSwapToken"
	"log"
	"math/big"
)

func DeployHbSwap(conn *ethclient.Client, auth *bind.TransactOpts) common.Address {
	fmt.Println("Deploying HbSwap contract...")

	var servers []common.Address
	for i := 0; i < utils.N; i++ {
		transactOpt := utils.GetAccount(fmt.Sprintf("poa/keystore/server_%v", i))
		servers = append(servers, transactOpt.From)
	}

	hbswapAddr, tx, _, err := hbswap.DeployHbSwap(auth, conn, servers, big.NewInt(utils.T))
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

	fmt.Println("Deployed HbSwap contract at", hbswapAddr.Hex())

	return hbswapAddr
}

func DeployToken(conn *ethclient.Client, auth *bind.TransactOpts) common.Address {
	fmt.Println("Deploying Token contract...")

	tokenAddr, tx, _, err := hbSwapToken.DeployHbSwapToken(auth, conn)
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

	fmt.Println("Deployed Token contract at", tokenAddr.Hex())

	return tokenAddr
}

func main() {
	//ethHostname := os.Args[1]
	//ethUrl := utils.GetEthURL(ethHostname)
	//conn := utils.GetEthClient(ethUrl)

	conn := utils.GetEthClient(utils.TestnetWsEndpoint)
	owner := utils.GetAccount("poa/keystore/server_0")

	DeployHbSwap(conn, owner)
	//DeployToken(conn, owner)
	//DeployToken(conn, owner)
}

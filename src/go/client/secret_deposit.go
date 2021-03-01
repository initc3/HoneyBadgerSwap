package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
	"os"
)

func secretDeposit(network string, conn *ethclient.Client, auth *bind.TransactOpts, token common.Address, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.SecretDeposit(network, conn, auth, token, amt)

	utils.GetBalance(network, conn, token, auth.From)
}

func main() {
	_network := flag.String("n", "testnet", "Type 'testnet' or 'privatenet'. Default: testnet")
	flag.Parse()
	network := *_network
	fmt.Println(network)

	user := os.Args[3]
	tokenA, tokenB := common.HexToAddress(os.Args[4]), common.HexToAddress(os.Args[5])
	amtA, amtB := os.Args[6], os.Args[7]

	ethHostname := os.Args[8]
	ethUrl := utils.GetEthURL(ethHostname)
	conn := utils.GetEthClient(ethUrl)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	secretDeposit(network, conn, owner, tokenA, amtA)
	secretDeposit(network, conn, owner, tokenB, amtB)
}

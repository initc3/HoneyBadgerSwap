package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go/client/lib"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
	"os"
)

func depositETH(network string, conn *ethclient.Client, auth *bind.TransactOpts, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.FundETH(network, conn, auth.From, amt)
	auth.Value = amt
	utils.Deposit(network, conn, auth, utils.EthAddr, amt)

	utils.GetBalance(network, conn, utils.EthAddr, auth.From)
}

func depositTOK(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.FundToken(conn, tokenAddr, auth.From, amt)

	auth.Value = big.NewInt(0)
	utils.Approve(network, conn, auth, tokenAddr, utils.HbswapAddr[network], amt)
	utils.Deposit(network, conn, auth, tokenAddr, amt)

	utils.GetBalance(network, conn, tokenAddr, auth.From)
}

func main() {
	// parse cmd line arguments/flags
	var configfile string
	flag.StringVar(&configfile, "config", "/opt/hbswap/conf/client.toml", "Config file. Default: /opt/hbswap/conf/client.toml")
	flag.Parse()

	// parse config file
	config := lib.ParseClientConfig(configfile)
	network := config.EthNode.Network
	ethHostname := config.EthNode.Hostname
	fmt.Println("Eth network: ", network)
	fmt.Println("Eth hostname: ", ethHostname)

	user := utils.GetAccount(fmt.Sprintf("account_%s", os.Args[3]))
	tokenA, tokenB := common.HexToAddress(os.Args[4]), common.HexToAddress(os.Args[5])
	amtA, amtB := os.Args[6], os.Args[7]

	ethUrl := ethHostname
	if network == "privatenet" {
		ethUrl = utils.GetEthURL(ethHostname)
	}
	conn := utils.GetEthClient(ethUrl)

	if bytes.Equal(tokenA.Bytes(), utils.EthAddr.Bytes()) {
		depositETH(network, conn, user, amtA)
	} else {
		depositTOK(network, conn, user, tokenA, amtA)
	}

	if bytes.Equal(tokenB.Bytes(), utils.EthAddr.Bytes()) {
		depositETH(network, conn, user, amtB)
	} else {
		depositTOK(network, conn, user, tokenB, amtB)
	}
}

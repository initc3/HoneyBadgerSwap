package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go/client/lib"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
	"os"
)

func secretWithdraw(conn *ethclient.Client, auth *bind.TransactOpts, token common.Address, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.SecretWithdraw(conn, auth, token, amt)

	utils.GetBalance(conn, token, auth.From)
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

	user := os.Args[1]
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amtA, amtB := os.Args[4], os.Args[5]

	ethUrl := ethHostname
	if network == "privatenet" {
		ethUrl = utils.GetEthURL(ethHostname)
	}
	} else {
		ethUrl = config.EthNode.HttpEndpoint
	}
	conn := utils.GetEthClient(ethUrl)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	secretWithdraw(conn, owner, tokenA, amtA)
	secretWithdraw(conn, owner, tokenB, amtB)
}

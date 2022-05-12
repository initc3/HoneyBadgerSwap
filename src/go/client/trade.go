package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initc3/HoneyBadgerSwap/src/go/client/lib"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// parse cmd line arguments/flags
	var configfile string
	flag.StringVar(&configfile, "config", "/opt/hbswap/conf/server.toml", "Config file. Default: /opt/hbswap/conf/server.toml")
	flag.Parse()

	// parse config file
	config := lib.ParseClientConfig(configfile)
	network := config.EthNode.Network
	ethHostname := config.EthNode.Hostname
	fmt.Println("Eth network: ", network)
	fmt.Println("Eth hostname: ", ethHostname)

	user := os.Args[3]
	tokenA, tokenB := common.HexToAddress(os.Args[4]), common.HexToAddress(os.Args[5])
	amtA, amtB := os.Args[6], os.Args[7]

	ethUrl := ethHostname
	if network == "privatenet" {
		ethUrl = utils.GetEthURL(ethHostname)
	} else {
		ethUrl = config.EthNode.HttpEndpoint
	}
	conn := utils.GetEthClient(ethUrl)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	idxA, idxB := utils.TradePrep(network, conn, owner)

	cmd := exec.Command("python3", "-m", "honeybadgerswap.client.req_inputmasks", strconv.Itoa(int(idxA)), amtA, strconv.Itoa(int(idxB)), amtB)
	stdout := utils.ExecCmd(cmd)
	maskedInputs := strings.Split(stdout[:len(stdout)-1], " ")

	maskedA := utils.StrToBig(maskedInputs[0])
	maskedB := utils.StrToBig(maskedInputs[1])

	fmt.Printf("maskedInputs: %v\n", maskedInputs)
	utils.Trade(network, conn, owner, tokenA, tokenB, big.NewInt(idxA), big.NewInt(idxB), maskedA, maskedB)
}

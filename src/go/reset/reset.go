// go run src/go/reset/reset.go

package main

import (
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
)

const (
	network = "testnet"
)

func main() {
	conn := utils.GetEthClient(utils.TestnetWsEndpoint)
	owner := utils.GetAccount("poa/keystore/server_0")

	utils.ResetPrice(network, conn, owner, utils.EthAddr, utils.HbSwapTokenAddr[network])
	utils.ResetPrice(network, conn, owner, utils.EthAddr, utils.DAIAddr[network])
	utils.ResetPrice(network, conn, owner, utils.DAIAddr[network], utils.HbSwapTokenAddr[network])
}
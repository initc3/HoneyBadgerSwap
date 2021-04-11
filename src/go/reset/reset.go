// go run /go/src/github.com/initc3/HoneyBadgerSwap/src/go/reset/reset.go

package main

import (
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
)

const (
	network = "testnet"
)

func main() {
	conn := utils.GetEthClient(utils.TestnetWsEndpoint)
	auth := utils.GetAccount("server_0")
	utils.ResetPrice(network, conn, auth, utils.EthAddr, utils.HbSwapTokenAddr[network])
	utils.ResetBalance(network, conn, auth, utils.EthAddr, utils.UserAddr)
	utils.ResetBalance(network, conn, auth, utils.HbSwapTokenAddr[network], utils.UserAddr)
	//for _, tokenAddr := range utils.TokenAddrs[network] {
	//	utils.ResetBalance(network, conn, auth, tokenAddr, utils.UserAddr)
	//}
}
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"os"
)

func main() {
	user := utils.GetAccount(fmt.Sprintf("account_%s", os.Args[1]))
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amtA, amtB := os.Args[4], os.Args[5]

	conn := utils.GetEthClient(utils.TestnetHttpEndpoint)

	utils.AddLiquidity(conn, user, tokenA, tokenB, utils.StrToBig(amtA), utils.StrToBig(amtB))
}

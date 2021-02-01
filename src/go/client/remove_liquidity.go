package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"os"
)

func main() {
	user := utils.GetAccount(fmt.Sprintf("account_%s", os.Args[1]))
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amt := os.Args[4]

	conn := utils.GetEthClient(utils.HttpEndpoint)

	utils.RemoveLiquidity(conn, user, tokenA, tokenB, utils.StrToBig(amt))
}





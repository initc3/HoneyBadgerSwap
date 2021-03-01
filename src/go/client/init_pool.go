package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"os"
)

//func prepareETH(conn *ethclient.Client, auth *bind.TransactOpts, _amt string) *big.Int {
//	amt := utils.StrToBig(_amt)
//	if amt.Cmp(big.NewInt(0)) == 1 {
//		utils.FundETH(conn, auth.From, amt)
//	}
//	return amt
//}
//
//func prepareTOK(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, _amt string) {
//	amt := utils.StrToBig(_amt)
//	if amt.Cmp(big.NewInt(0)) == 0 {
//		return
//	}
//
//	utils.FundToken(conn, tokenAddr, auth.From, amt)
//
//	utils.Approve(conn, auth, tokenAddr, utils.HbswapAddr, amt)
//}

func main() {
	_network := flag.String("n", "testnet", "Type 'testnet' or 'privatenet'. Default: testnet")
	flag.Parse()
	network := *_network
	fmt.Println(network)

	user := utils.GetAccount(fmt.Sprintf("account_%s", os.Args[3]))
	tokenA, tokenB := common.HexToAddress(os.Args[4]), common.HexToAddress(os.Args[5])
	amtA, amtB := os.Args[6], os.Args[7]

	ethHostname := os.Args[8]
	ethUrl := utils.GetEthURL(ethHostname)
	conn := utils.GetEthClient(ethUrl)

	//ethA := bytes.Equal(tokenA.Bytes(), utils.EthAddr.Bytes())
	//ethB := bytes.Equal(tokenB.Bytes(), utils.EthAddr.Bytes())
	//
	//if !ethA {
	//	prepareTOK(conn, user, tokenA, amtA)
	//}
	//if !ethB {
	//	prepareTOK(conn, user, tokenB, amtB)
	//}
	//
	//value := big.NewInt(0)
	//if ethA {
	//	value = prepareETH(conn, user, amtA)
	//}
	//if ethB {
	//	value = prepareETH(conn, user, amtB)
	//}

	utils.InitPool(network, conn, user, tokenA, tokenB, utils.StrToBig(amtA), utils.StrToBig(amtB))
}

// go run src/go/client/testnet.go

package main

import (
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
)

var (
	balance *big.Int
)

func main() {
	network := "testnet"
	conn := utils.GetEthClient(utils.TestnetWsEndpoint)
	//conn := utils.GetEthClient("http://127.0.0.1:8545")

	//admin := utils.GetAccount("poa/keystore/server_0")
	//balance, _ = conn.BalanceAt(context.Background(), admin.From, nil)
	//fmt.Printf("admin eth balance %v\n", balance)
	//balance = utils.GetBalanceToken(conn, admin.From, utils.HbSwapTokenAddr[network])
	//fmt.Printf("admin hbs balance %v\n", balance)
	////
	//peer := utils.GetAccount(fmt.Sprintf("account_0"))
	//peer := utils.GetAccount(fmt.Sprintf("server_3"))
	//peer := common.HexToAddress("0xc33a4b5b609fcc294dca060347761226e78c0b7a") //Metamask
	//peer := common.HexToAddress("0x3C19cA734eeaA2b3617C76afa993A54b5C6f6448") //sylvain
	//peer := common.HexToAddress("0x786e3c83cd270414649079A758Ad92f961EDdA0A") //DAI
	//peer := common.HexToAddress("0x37a25f181a43613e2917c650bd6d2f2bb26defd2") //new server_0
	//peer := common.HexToAddress("0xabbf7aedbb03f7fb09b6c7a41623ef67b97570b2") //new server_1
	//peer := common.HexToAddress("0x00377f482929288e237c50578553641a0a2e4ce2") //new server_2
	//peer := common.HexToAddress("0xae76f3dfa7a7b1556bb49ee41663b7ac6be58da5") //new server_3

	/* transfer eth */
	//balance, _ = conn.BalanceAt(context.Background(), peer, nil)
	//fmt.Printf("peer eth balance %v\n", balance)
	//utils.FundETH("testnet", conn, peer, big.NewInt(5 * 1000000000000000000))
	//balance, _ = conn.BalanceAt(context.Background(), peer, nil)
	//fmt.Printf("peer eth balance %v\n", balance)

	/* transfer hbs */
	//balance = utils.GetBalanceToken(conn, peer, utils.HbSwapTokenAddr[network])
	//fmt.Printf("peer hbs balance %v\n", balance)
	//utils.FundToken(conn, utils.HbSwapTokenAddr[network], peer, big.NewInt(5 * 1000000000000000000))
	//balance = utils.GetBalanceToken(conn, peer, utils.HbSwapTokenAddr[network])
	//fmt.Printf("peer hbs balance %v\n", balance)
	//
	for i := 0; i < 100; i++ {
		utils.GetInputmaskCnt(network, conn)
	}
}

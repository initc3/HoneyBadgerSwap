// go run /go/src/github.com/initc3/HoneyBadgerSwap/src/go/client/testnet.go

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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

	admin := utils.GetAccount(fmt.Sprintf("server_0"))
	balance, _ = conn.BalanceAt(context.Background(), admin.From, nil)
	fmt.Printf("admin eth balance %v\n", balance)
	balance = utils.GetBalanceToken(conn, admin.From, utils.HbSwapTokenAddr[network])
	fmt.Printf("admin hbs balance %v\n", balance)

	//peer := utils.GetAccount(fmt.Sprintf("account_0"))
	//peer := utils.GetAccount(fmt.Sprintf("server_3"))
	peer := common.HexToAddress("0xc33a4b5b609fcc294dca060347761226e78c0b7a") //Metamask
	//peer := common.HexToAddress("0x3C19cA734eeaA2b3617C76afa993A54b5C6f6448")
	//peer := common.HexToAddress("0x786e3c83cd270414649079A758Ad92f961EDdA0A") //DAI

	/* transfer eth */
	//balance, _ = conn.BalanceAt(context.Background(), peer, nil)
	//fmt.Printf("peer eth balance %v\n", balance)
	//utils.FundETH("testnet", conn, peer, big.NewInt(0.1 * 1000000000000000000))
	//balance, _ = conn.BalanceAt(context.Background(), peer, nil)
	//fmt.Printf("peer eth balance %v\n", balance)

	/* transfer hbs */
	balance = utils.GetBalanceToken(conn, peer, utils.HbSwapTokenAddr[network])
	fmt.Printf("peer hbs balance %v\n", balance)
	utils.FundToken(conn, utils.HbSwapTokenAddr[network], peer, big.NewInt(10000))
	balance = utils.GetBalanceToken(conn, peer, utils.HbSwapTokenAddr[network])
	fmt.Printf("peer hbs balance %v\n", balance)
}

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
)

func main() {
	conn := utils.GetEthClient("https://kovan.infura.io/v3/6a82d2519efb4d748c02552e02e369c1")
	//conn := utils.GetEthClient("http://127.0.0.1:8545")

	//admin := utils.GetAccount(fmt.Sprintf("server_0"))

	//peer := utils.GetAccount(fmt.Sprintf("account_0"))
	//peer := utils.GetAccount(fmt.Sprintf("server_3"))
	peer := common.HexToAddress("0xc33a4b5b609fcc294dca060347761226e78c0b7a")

	var balance *big.Int

	//balance, _ = conn.BalanceAt(context.Background(), admin.From, nil)
	//fmt.Printf("balance %v\n", balance)
	//
	//utils.FundETH(conn, peer, utils.StrToBig("1000000000000000000"))
	//
	//balance, _ = conn.BalanceAt(context.Background(), peer, nil)
	//fmt.Printf("balance %v\n", balance)

	token := utils.TokenAddrs[0]
	balance = utils.GetBalanceToken(conn, peer, token)
	fmt.Printf("balance %v\n", balance)

	utils.FundToken(conn, token, peer, big.NewInt(10000))
	balance = utils.GetBalanceToken(conn, peer, token)
	fmt.Printf("balance %v\n", balance)
}

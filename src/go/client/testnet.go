package main

import (
	"context"
	"fmt"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"math/big"
)

func main() {
	conn := utils.GetEthClient("https://kovan.infura.io/v3/6a82d2519efb4d748c02552e02e369c1")

	admin := utils.GetAccount(fmt.Sprintf("server_0"))

	balance, _ := conn.BalanceAt(context.Background(), admin.From, nil)
	fmt.Printf("balance %v\n", balance)

	peer := utils.GetAccount(fmt.Sprintf("server_1"))
	utils.FundETH(conn, peer.From, big.NewInt(1))
	balance, _ = conn.BalanceAt(context.Background(), peer.From, nil)
	fmt.Printf("balance %v\n", balance)
}

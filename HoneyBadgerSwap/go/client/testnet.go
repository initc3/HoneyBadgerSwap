package main

import (
	"context"
	"fmt"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
)

func main() {
	conn := utils.GetEthClient("https://kovan.infura.io/v3/6a82d2519efb4d748c02552e02e369c1")

	user := "0"
	owner, _ := utils.GetAccount(fmt.Sprintf("account_%s", user))

	balance, _ := conn.BalanceAt(context.Background(), owner.From, nil)
	fmt.Printf("balance %v\n", balance)
}

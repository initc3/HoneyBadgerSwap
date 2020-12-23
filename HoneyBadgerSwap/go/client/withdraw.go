package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"math/big"
	"os"
)

func withdrawETH(conn *ethclient.Client, auth *bind.TransactOpts, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.Withdraw(conn, auth, utils.EthAddr, amt)

	utils.GetBalance(conn, utils.EthAddr, auth.From)
}

func withdrawTOK(conn *ethclient.Client, auth *bind.TransactOpts, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.Withdraw(conn, auth, utils.TokenAddr, amt)

	utils.GetBalance(conn, utils.TokenAddr, auth.From)
}

func main() {
	user := os.Args[1]
	amtETH, amtTOK := os.Args[2], os.Args[3]

	conn := utils.GetEthClient(utils.HttpEndpoint)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	withdrawETH(conn, owner, amtETH)
	withdrawTOK(conn, owner, amtTOK)
}
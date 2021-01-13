package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"math/big"
	"os"
)

func secretDeposit(conn *ethclient.Client, auth *bind.TransactOpts, token common.Address, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.SecretDeposit(conn, auth, token, amt)

	utils.GetBalance(conn, token,  auth.From)
}

func main() {
	user := os.Args[1]
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amtA, amtB := os.Args[4], os.Args[5]

	conn := utils.GetEthClient(utils.HttpEndpoint)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	secretDeposit(conn, owner, tokenA, amtA)
	secretDeposit(conn, owner, tokenB, amtB)
}





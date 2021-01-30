package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	utils.GetBalance(conn, utils.EthAddr, auth.From)

	utils.Withdraw(conn, auth, utils.EthAddr, amt)

	utils.GetBalance(conn, utils.EthAddr, auth.From)
}

func withdrawTOK(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.GetBalance(conn, tokenAddr, auth.From)

	utils.Withdraw(conn, auth, tokenAddr, amt)

	utils.GetBalance(conn, tokenAddr, auth.From)
}

func main() {
	user := os.Args[1]
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amtA, amtB := os.Args[4], os.Args[5]

	ethHostname := os.Args[6]
	ethUrl := utils.GetEthURL(ethHostname)
	conn := utils.GetEthClient(ethUrl)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	if bytes.Equal(tokenA.Bytes(), utils.EthAddr.Bytes()) {
		withdrawETH(conn, owner, amtA)
	} else {
		withdrawTOK(conn, owner, tokenA, amtA)
	}

	if bytes.Equal(tokenB.Bytes(), utils.EthAddr.Bytes()) {
		withdrawETH(conn, owner, amtB)
	} else {
		withdrawTOK(conn, owner, tokenB, amtB)
	}
}

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

func depositETH(conn *ethclient.Client, auth *bind.TransactOpts, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.FundETH(conn, auth.From, amt)
	auth.Value = amt
	utils.Deposit(conn, auth, utils.EthAddr, amt)

	utils.GetBalance(conn, utils.EthAddr, auth.From)
}

func depositTOK(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.FundToken(conn, tokenAddr, auth.From, amt)

	auth.Value = big.NewInt(0)
	utils.Approve(conn, auth, tokenAddr, utils.HbswapAddr, amt)
	utils.Deposit(conn, auth, tokenAddr, amt)

	utils.GetBalance(conn, tokenAddr, auth.From)
}

func main() {
	user := utils.GetAccount(fmt.Sprintf("account_%s", os.Args[1]))
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amtA, amtB := os.Args[4], os.Args[5]

	conn := utils.GetEthClient(utils.HttpEndpoint)

	if bytes.Equal(tokenA.Bytes(), utils.EthAddr.Bytes()) {
		depositETH(conn, user, amtA)
	} else {
		depositTOK(conn, user, tokenA, amtA)
	}

	if bytes.Equal(tokenB.Bytes(), utils.EthAddr.Bytes()) {
		depositETH(conn, user, amtB)
	} else {
		depositTOK(conn, user, tokenB, amtB)
	}
}
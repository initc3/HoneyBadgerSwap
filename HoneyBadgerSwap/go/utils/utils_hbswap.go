package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go_bindings/hbswap"
	"log"
	"math/big"
	"strconv"
	"time"
)

func Deposit(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.Deposit(auth, tokenAddr, amt)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

func Withdraw(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.Withdraw(auth, tokenAddr, amt)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

func SecretDeposit(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	fmt.Printf("SecretDeposit %s %v\n", tokenAddr.String(), amt)
	tx, err := hbswapInstance.SecretDeposit(auth, tokenAddr, amt)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

func SecretWithdraw(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
	fmt.Printf("secret withdraw %s\n", tokenAddr.Hex())
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	prevBalance := GetBalance(conn, tokenAddr, auth.From).Int64()

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.SecretWithdraw(auth, tokenAddr, amt)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}

	fmt.Printf("current balance %v\n", prevBalance)
	for true {
		time.Sleep(10 * time.Second)
		balance := GetBalance(conn, tokenAddr, auth.From).Int64()
		fmt.Printf("current balance %v\n", balance)
		if prevBalance + amt.Int64() == balance {
			break
		}
	}
}

func Consent(conn *ethclient.Client, auth *bind.TransactOpts, seq *big.Int) {
	fmt.Printf("Consent seq %v\n", seq)
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.Consent(auth, seq)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

func TradePrep(conn *ethclient.Client, auth *bind.TransactOpts) (int64, int64) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.TradePrep(auth)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}

	data := receipt.Logs[0].Data
	idxSell, err := strconv.ParseInt(common.Bytes2Hex(data[1 * 32 : 2 * 32]), 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	idxBuy, err := strconv.ParseInt(common.Bytes2Hex(data[2 * 32 : 3 * 32]), 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("idxSell:%v idxBuy:%v\n", idxSell, idxBuy)
	return idxSell, idxBuy
}

func Trade(conn *ethclient.Client, auth *bind.TransactOpts, tokenSell common.Address, tokenBuy common.Address, idxSell *big.Int, idxBuy *big.Int, maskedSell *big.Int, maskedBuy *big.Int) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.Trade(auth, tokenSell, tokenBuy, idxSell, idxBuy, maskedSell, maskedBuy)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

func Reset(conn *ethclient.Client, auth *bind.TransactOpts) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.Reset(auth)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: %v", receipt.Status)
	}
}

func GetBalance(conn *ethclient.Client, tokenAddr common.Address, user common.Address) *big.Int {
	fmt.Printf("Getting balance for token %v user %v\n", tokenAddr.Hex(), user.Hex())
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		fmt.Printf("here")
		log.Fatal(err)
	}

	balance, _ := hbswapInstance.Balances(nil, tokenAddr, user)
	log.Printf("%s balance: %v\n", tokenAddr.Hex(), balance)

	return balance
}

func GetInputmaskCnt(conn *ethclient.Client) *big.Int {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	cnt, _ := hbswapInstance.InputmaskCnt(nil)
	log.Printf("inputmaskCnt: %v\n", cnt)

	return cnt
}

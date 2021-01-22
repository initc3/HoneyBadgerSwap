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

func InitPool(conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amtA *big.Int, amtB *big.Int) {
	fmt.Printf("InitPool tokenA %s tokenB %s amtA %v amtB %v\n", tokenA.Hex(), tokenB.Hex(), amtA, amtB)
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.InitPool(auth, tokenA, tokenB, amtA, amtB)
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

func AddLiquidity(conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amtA *big.Int, amtB *big.Int) {
	fmt.Printf("AddLiquidity tokenA %s tokenB %s amtA %v amtB %v\n", tokenA.Hex(), tokenB.Hex(), amtA, amtB)
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.AddLiquidity(auth, tokenA, tokenB, amtA, amtB)
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

func RemoveLiquidity(conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amt *big.Int) {
	fmt.Printf("RemoveLiquidity tokenA %s tokenB %s amt %v\n", tokenA.Hex(), tokenB.Hex(), amt)
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.RemoveLiquidity(auth, tokenA, tokenB, amt)
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

func Deposit(conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
	fmt.Printf("Deposit user %s %v token %s\n", auth.From.Hex(), amt, tokenAddr.Hex())
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

	//fmt.Printf("current balance %v\n", prevBalance)
	for true {
		time.Sleep(10 * time.Second)
		balance := GetBalance(conn, tokenAddr, auth.From).Int64()
		//fmt.Printf("current balance %v\n", balance)
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

	fmt.Printf("Consent seq %v done\n", seq)
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

	fmt.Printf("idxSell:%v idxBuy:%v\n", idxSell, idxBuy)
	return idxSell, idxBuy
}

func Trade(conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, idxA *big.Int, idxB *big.Int, maskedA *big.Int, maskedB *big.Int) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.Trade(auth, tokenA, tokenB, idxA, idxB, maskedA, maskedB)
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
	tradeSeq, err := strconv.ParseInt(common.Bytes2Hex(data[0 * 32 : 1 * 32]), 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tradeSeq %v txHash %s\n", tradeSeq, tx.Hash().Hex())
}

func UpdatePrice(conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, price string) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	GetPrice(conn, tokenA, tokenB)

	fundGas(conn, auth.From)
	tx, err := hbswapInstance.UpdatePrice(auth, tokenA, tokenB, price)
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

	GetPrice(conn, tokenA, tokenB)
}

/******** read-only ********/

func GetBalance(conn *ethclient.Client, tokenAddr common.Address, user common.Address) *big.Int {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		fmt.Printf("here")
		log.Fatal(err)
	}

	balance, _ := hbswapInstance.Balances(nil, tokenAddr, user)
	fmt.Printf("On-chain balance user %s token %s: %v\n", user.Hex(), tokenAddr.Hex(), balance)

	return balance
}

func GetInputmaskCnt(conn *ethclient.Client) int64 {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	cnt, _ := hbswapInstance.InputmaskCnt(nil)
	fmt.Printf("Inputmaks shares used: %v\n", cnt)

	return cnt.Int64()
}

func GetPrice(conn *ethclient.Client, tokenA common.Address, tokenB common.Address) string {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	price, _ := hbswapInstance.Prices(nil, tokenA, tokenB)
	fmt.Printf("price: %v\n", price)

	return price
}

func GetUpdateTime(conn *ethclient.Client, tokenA common.Address, tokenB common.Address) int64 {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr, conn)
	if err != nil {
		log.Fatal(err)
	}

	t, _ := hbswapInstance.UpdateTimes(nil, tokenA, tokenB)
	fmt.Printf("updateTime: %v\n", t)

	return t.Int64()
}

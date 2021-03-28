package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go_bindings/hbswap"
	"log"
	"math/big"
)

//func Deposit(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
//	fmt.Printf("Deposit user %s %v token %s\n", auth.From.Hex(), amt, tokenAddr.Hex())
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.Deposit(auth, tokenAddr, amt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//}
//
//func SecretDeposit(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	fmt.Printf("SecretDeposit %s %v\n", tokenAddr.String(), amt)
//	tx, err := hbswapInstance.SecretDeposit(auth, tokenAddr, amt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//}
//
//func Withdraw(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.Withdraw(auth, tokenAddr, amt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//}
//
//func SecretWithdraw(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenAddr common.Address, amt *big.Int) {
//	fmt.Printf("secret withdraw %s\n", tokenAddr.Hex())
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	prevBalance := GetBalance(network, conn, tokenAddr, auth.From).Int64()
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.SecretWithdraw(auth, tokenAddr, amt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//
//	//fmt.Printf("current balance %v\n", prevBalance)
//	for true {
//		time.Sleep(10 * time.Second)
//		balance := GetBalance(network, conn, tokenAddr, auth.From).Int64()
//		//fmt.Printf("current balance %v\n", balance)
//		if prevBalance+amt.Int64() == balance {
//			break
//		}
//	}
//}
//
//func ReserveInput(network string, conn *ethclient.Client, auth *bind.TransactOpts, num *big.Int) []int64 {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.ReserverInput(auth, num)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//
//	var inputMaskIndexes []int64
//	//TODO: fix below
//	data := receipt.Logs[0].Data
//	idxSell, err := strconv.ParseInt(common.Bytes2Hex(data[1*32:2*32]), 16, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	idxBuy, err := strconv.ParseInt(common.Bytes2Hex(data[2*32:3*32]), 16, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("idxSell:%v idxBuy:%v\n", idxSell, idxBuy)
//	return idxSell, idxBuy
//}

//func InitPool(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amtA *big.Int, amtB *big.Int) {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.InitPool(auth, tokenA, tokenB, amtA, amtB)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//}
//
//func AddLiquidity(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, idxA *big.Int, idxB *big.Int, maskedAmtA *big.Int, maskedAmtB *big.Int) {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.AddLiquidity(auth, tokenA, tokenB, idxA, idxB, maskedAmtA, maskedAmtB)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//}
//
//func RemoveLiquidity(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, idx *big.Int, amt *big.Int) {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.RemoveLiquidity(auth, tokenA, tokenB, idx, amt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//}

func Consent(network string, conn *ethclient.Client, auth *bind.TransactOpts, seq *big.Int) {
	fmt.Printf("Consent seq %v\n", seq)
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		log.Fatal(err)
	}

	fundGas(network, conn, auth.From)
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

//func Trade(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, idxA *big.Int, idxB *big.Int, maskedAmtA *big.Int, maskedAmtB *big.Int) {
//	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fundGas(network, conn, auth.From)
//	tx, err := hbswapInstance.Trade(auth, tokenA, tokenB, idxA, idxB, maskedAmtA, maskedAmtB)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	receipt, err := WaitMined(context.Background(), conn, tx, 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if receipt.Status == 0 {
//		log.Fatalf("Transaction status: %v", receipt.Status)
//	}
//
//	data := receipt.Logs[0].Data
//	tradeSeq, err := strconv.ParseInt(common.Bytes2Hex(data[0*32:1*32]), 16, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("tradeSeq %v txHash %s\n", tradeSeq, tx.Hash().Hex())
//}

func UpdatePrice(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address, checkpointSeq *big.Int, price string) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		log.Fatal(err)
	}

	GetPrice(network, conn, tokenA, tokenB)

	fundGas(network, conn, auth.From)
	tx, err := hbswapInstance.UpdatePrice(auth, tokenA, tokenB, checkpointSeq, price)
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

	GetPrice(network, conn, tokenA, tokenB)
}

func ResetPrice(network string, conn *ethclient.Client, auth *bind.TransactOpts, tokenA common.Address, tokenB common.Address) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		log.Fatal(err)
	}

	price := GetPrice(network, conn, tokenA, tokenB)
	if price == "" {
		return
	}

	var servers []common.Address
	for i := 0; i < N; i++ {
		transactOpt := GetAccount(fmt.Sprintf("server_%v", i))
		servers = append(servers, transactOpt.From)
	}

	fmt.Println("ResetPrice")
	tx, err := hbswapInstance.ResetPrice(auth, tokenA, tokenB, servers)
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

func ResetBalance(network string, conn *ethclient.Client, auth *bind.TransactOpts, token common.Address, user common.Address) {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		log.Fatal(err)
	}

	balance := GetBalance(network, conn, token, user)
	if balance.Cmp(big.NewInt(0)) == 0 {
		return
	}

	fmt.Printf("ResetBalance token %s user %s\n", token.Hex(), user.Hex())
	tx, err := hbswapInstance.ResetBalance(auth, token, user)
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

	GetBalance(network, conn, token, user)
}

/******** read-only ********/

func GetBalance(network string, conn *ethclient.Client, token common.Address, user common.Address) *big.Int {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		fmt.Printf("here")
		log.Fatal(err)
	}

	balance, _ := hbswapInstance.PublicBalance(nil, token, user)
	fmt.Printf("GetBalance token %s user %s balance %v\n", token.Hex(), user.Hex(), balance)

	return balance
}

func GetInputmaskCnt(network string, conn *ethclient.Client) int64 {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		log.Fatal(err)
	}

	cnt, _ := hbswapInstance.InputmaskCnt(nil)
	fmt.Printf("Inputmaks shares used: %v\n", cnt)

	return cnt.Int64()
}

func GetPrice(network string, conn *ethclient.Client, tokenA common.Address, tokenB common.Address) string {
	hbswapInstance, err := hbswap.NewHbSwap(HbswapAddr[network], conn)
	if err != nil {
		log.Fatal(err)
	}

	price, _ := hbswapInstance.Prices(nil, tokenA, tokenB)
	fmt.Printf("price: %v\n", price)

	return price
}
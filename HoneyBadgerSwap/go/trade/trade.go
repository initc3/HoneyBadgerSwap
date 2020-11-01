package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/gobingdings/hbswap"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	hbswapAddr = "0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2"
)

func TradePrep(conn *ethclient.Client, auth *bind.TransactOpts) (int64, int64) {
	hbswapInstance, err := hbswap.NewHbSwap(common.HexToAddress(hbswapAddr), conn)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := hbswapInstance.TradePrep(auth)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := utils.WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: ", receipt.Status)
	}

	data := receipt.Logs[0].Data
	idxETH, err := strconv.ParseInt(common.Bytes2Hex(data[1 * 32 : 2 * 32]), 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	idxTOK, err := strconv.ParseInt(common.Bytes2Hex(data[2 * 32 : 3 * 32]), 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	return idxETH, idxTOK
}

func Trade(conn *ethclient.Client, auth *bind.TransactOpts, idxETH *big.Int, idxTOK *big.Int, maskedETH *big.Int, maskedTOK *big.Int) {
	hbswapInstance, err := hbswap.NewHbSwap(common.HexToAddress(hbswapAddr), conn)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := hbswapInstance.Trade(auth, idxETH, idxTOK, maskedETH, maskedTOK)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := utils.WaitMined(context.Background(), conn, tx, 0)
	if err != nil {
		log.Fatal(err)
	}

	if receipt.Status == 0 {
		log.Fatalf("Transaction status: ", receipt.Status)
	}
}

func main() {
	amtETH, amtTOK := os.Args[1], os.Args[2]

	conn := utils.GetEthClient("HTTP://127.0.0.1:8545")

	owner, _ := utils.GetAccount("account_0")

	idxETH, idxTOK := TradePrep(conn, owner)

	cmd := exec.Command("python3", "Scripts/hbswap/python/req_inputmasks.py", strconv.Itoa(int(idxETH)), strconv.Itoa(int(idxTOK)), amtETH, amtTOK)
	stdout := utils.ExecCmd(cmd)
	maskedInputs := strings.Split(stdout[:len(stdout) - 1], " ")

	maskedETH, _ := utils.StrToBig(maskedInputs[0])
	maskedTOK, _ := utils.StrToBig(maskedInputs[1])

	fmt.Printf("maskedInputs: %v\n", maskedInputs)
	Trade(conn, owner, big.NewInt(idxETH), big.NewInt(idxTOK), maskedETH, maskedTOK)
}





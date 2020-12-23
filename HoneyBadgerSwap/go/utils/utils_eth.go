package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/********** external functions **********/

func FundETH(conn *ethclient.Client, toAddr common.Address, amount *big.Int) {
	adminAuth := GetAccount("server_0")
	transferETH(conn, chainID, adminAuth, toAddr, amount)

	balance := GetBalanceETH(conn, toAddr)
	log.Printf("Funded account %s to %v ETH\n", toAddr.Hex(), balance)
}

func GetBalanceETH(conn *ethclient.Client, addr common.Address) (*big.Int) {
	balance, err := conn.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Println(err)
	}

	return balance
}

/********** internal functions **********/
func fundGas(conn *ethclient.Client, toAddr common.Address) {
	balance := GetBalanceETH(conn, toAddr)
	amount := minBalance
	if balance.Cmp(amount) != -1 {
		log.Printf("Funded %s to %v ETH on mainchain\n", toAddr.Hex(), balance)
		return
	}
	amount.Sub(amount, balance)

	adminAuth := GetAccount("server_0")
	transferETH(conn, chainID, adminAuth, toAddr, amount)

	balance = GetBalanceETH(conn, toAddr)
	log.Printf("Funded account %s to %v ETH\n", toAddr.Hex(), balance)
}

func transferETH(conn *ethclient.Client, chainId string, fromAuth *bind.TransactOpts, toAddr common.Address, amount *big.Int) {
	ctx := context.Background()

	fromAddr := fromAuth.From

	nonce, err := conn.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, toAddr, amount, gasLimit, gasPrice, nil)
	signedTx, err := fromAuth.Signer(types.NewEIP155Signer(stringToBigInt(chainId)), fromAddr, tx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sending %v wei from %s to %s\n", amount, fromAddr.Hex(), toAddr.Hex())
	err = conn.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Println(err)
		transferETH(conn, chainId, fromAuth, toAddr, amount)
	}

	log.Printf("Waiting for transaction to be mined...\n")
	receipt, err := WaitMined(ctx, conn, signedTx, 0)
	if err != nil {
		log.Println(err)
		transferETH(conn, chainId, fromAuth, toAddr, amount)
	}
	if receipt.Status == 0 {
		log.Printf("Transaction status: %x", receipt.Status)
		transferETH(conn, chainId, fromAuth, toAddr, amount)
	}
}
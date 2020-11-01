package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	GOPATH					= os.Getenv("GOPATH")
)

func ExecCmd(cmd *exec.Cmd) string {
	fmt.Println(cmd)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("err:\n%s\n", stderr.String())
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("out:\n%s\n", stdout.String())
	return stdout.String()
}

func StrToBig(st string) (*big.Int, bool) {
	return new(big.Int).SetString(st, 10)
}

func GetEthClient(ethInstance string) (*ethclient.Client) {
	conn, err := ethclient.Dial(ethInstance)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func GetAccount(account string) (*bind.TransactOpts, *ecdsa.PrivateKey) {
	dir := GOPATH + "/src/github.com/initc3/MP-SPDZ/Scripts/hbswap/poa/keystore/" + account + "/"

	list, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var name string
	for _, info := range list {
		name = info.Name()
		if err != nil {
			log.Fatal(err)
		}
	}

	bytes, err := ioutil.ReadFile(dir + name)
	if err != nil {
		log.Fatal(err)
	}

	password := ""
	auth, err := bind.NewTransactor(strings.NewReader(string(bytes)), password)
	if err != nil {
		log.Fatal(err)
	}

	keys, _ := keystore.DecryptKey(bytes, password)
	privateKey := keys.PrivateKey

	auth.GasLimit = 8000000

	return auth, privateKey
}

func WaitMined(ctx context.Context, ec *ethclient.Client,
	tx *types.Transaction, blockDelay uint64) (*types.Receipt, error) {
	// an error possibly returned when a transaction is pending
	const missingFieldErr = "missing required field 'transactionHash' for Log"

	if ec == nil {
		return nil, errors.New("nil ethclient")
	}
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	// wait tx to be mined
	txHashBytes := common.HexToHash(tx.Hash().Hex())
	for {
		receipt, rerr := ec.TransactionReceipt(ctx, txHashBytes)
		if rerr == nil {
			//log.Infof("Transaction mined. Waiting for %d block confirmations", blockDelay)
			if blockDelay == 0 {
				return receipt, rerr
			}
			break
		} else if rerr == ethereum.NotFound || rerr.Error() == missingFieldErr {
			// Wait for the next round
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-queryTicker.C:
			}
		} else {
			return receipt, rerr
		}
	}
	// wait for enough block confirmations
	ddl := big.NewInt(0)
	latestBlockHeader, err := ec.HeaderByNumber(ctx, nil)
	if err == nil {
		ddl.Add(new(big.Int).SetUint64(blockDelay), latestBlockHeader.Number)
	}
	for {
		latestBlockHeader, err := ec.HeaderByNumber(ctx, nil)
		if err == nil && ddl.Cmp(latestBlockHeader.Number) < 0 {
			receipt, rerr := ec.TransactionReceipt(ctx, txHashBytes)
			if rerr == nil {
				log.Println("tx confirmed!")
				return receipt, rerr
			} else if rerr == ethereum.NotFound || rerr.Error() == missingFieldErr {
				return nil, errors.New("tx is dropped due to chain re-org")
			} else {
				return receipt, rerr
			}
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}


package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	N = 4
	T = 1
)

var (
	GOPATH     = os.Getenv("GOPATH")
	KEYSTORE   = os.Getenv("POA_KEYSTORE")
	minBalance = big.NewInt(300000000000000000)

	chainID = map[string]string{
		"testnet": "42",
		"privatenet": "123",
	}

	EthAddr    = common.HexToAddress("0x0000000000000000000000000000000000000000")
	HbswapAddr = map[string]common.Address{
		"testnet":    common.HexToAddress("0x9ed1a58ff0479e36a4ead46647741f72ec9c15fe"),
		"privatenet": common.HexToAddress("0xf74eb25ab1785d24306ca6b3cbff0d0b0817c5e2"),
	}
	HbSwapTokenAddr = map[string]common.Address{
		"testnet":    common.HexToAddress("0x78160ee9e55fd81626f98d059c84d21d8b71bfda"),
	}
	DAIAddr = map[string]common.Address{
		"testnet":    common.HexToAddress("0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa"),
	}
	//TokenAddrs = map[string][]common.Address{
	//	"testnet": {
	//		common.HexToAddress("0x63e7f20503256ddcfec64872aadb785d5a290cbb"),
	//		common.HexToAddress("0x403b0f962566ffb960d0de98875dc09603aa67e9"),
	//	},
	//	"privatenet": {
	//		common.HexToAddress("0x6b5c9637e0207c72Ee1a275b6C3b686ba8D87385"),
	//		common.HexToAddress("0x8C89e5D2bCc0e4C26E3295d48d052E11bd03C06A"),
	//	},
	//}
	//TODO: delete it after testing
	UserAddr = common.HexToAddress("0xc33a4b5b609fcc294dca060347761226e78c0b7a")

	HttpPort = 8545
	WsPort  = 8546

	TestnetHttpEndpoint = "https://kovan.infura.io/v3/6a82d2519efb4d748c02552e02e369c1"
	TestnetWsEndpoint   = "wss://kovan.infura.io/ws/v3/6a82d2519efb4d748c02552e02e369c1"

	AddLiquidity = "0xec7d4752dd44bf7fc59045c9d80163de2a1b9dbd9032d11cb1156f7f867c6411"
	InitPool = "0xfaaebcb30b1b421f4f2ca7f2620e5add6a64532c087ee0646fd665a33d36fdf5"
	RemoveLiquidity = "0xa8dbaaebbb025c88e9e34c84635cd8238043556e9af43fb161508c898a8e1ef9"
	SecretDeposit = "0x07c06144435b7d2bdccf9ee7e5a7022c63382ac7c3a0e14ed08b5969dedf0ecf"
	SecretWithdraw = "0x4ef3cc4825a92c3b6922acc8a45152cc96ef48463e8ed500dacd5df9e29a67f3"
	Trade = "0x2b4d91cd20cc8800407e3614b8466a6f0729ac3b1fa43d4e2b059ff5593cbae6"
)

func ExecCmd(cmd *exec.Cmd) string {
	fmt.Printf("Cmd:\n====================\n%v\n====================\n", cmd)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
	if err != nil {
		fmt.Printf("err:\n%s\n", stderr.String())
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("Output:\n====================\n%s====================\n", stdout.String())
	return stdout.String()
}

func StrToBig(st string) *big.Int {
	v, _ := new(big.Int).SetString(st, 10)
	return v
}

func GetEthClient(ethInstance string) *ethclient.Client {
	conn, err := ethclient.Dial(ethInstance)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func GetAccount(account string) *bind.TransactOpts {
	dir := filepath.Join(KEYSTORE, account)

	list, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Error reading directory %s", dir)
		log.Fatal(err)
	}

	var name string
	for _, info := range list {
		name = info.Name()
		if err != nil {
			log.Fatal(err)
		}
	}

	//bytes, err := ioutil.ReadFile(dir + name)
	bytes, err := ioutil.ReadFile(filepath.Join(dir, name))
	if err != nil {
		log.Fatal(err)
	}

	password := ""
	auth, err := bind.NewTransactor(strings.NewReader(string(bytes)), password)
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 8000000

	return auth
}

func WaitMined(ctx context.Context, ec *ethclient.Client,
	tx *types.Transaction, blockDelay uint64) (*types.Receipt, error) {
	const missingFieldErr = "missing required field 'transactionHash' for Log"

	if ec == nil {
		return nil, errors.New("nil ethclient")
	}
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	txHashBytes := common.HexToHash(tx.Hash().Hex())
	for {
		receipt, rerr := ec.TransactionReceipt(ctx, txHashBytes)
		if rerr == nil {
			if blockDelay == 0 {
				return receipt, rerr
			}
			break
		} else if rerr == ethereum.NotFound || rerr.Error() == missingFieldErr {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-queryTicker.C:
			}
		} else {
			return receipt, rerr
		}
	}
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
				//fmt.Println("tx confirmed!")
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

func stringToBigInt(v string) *big.Int {
	value := big.NewInt(0)
	value.SetString(v, 10)
	return value
}

func GetIPAddr(hostname string) net.IP {
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Sprintf("Error looking up hostname %s", hostname)
		log.Fatal(err)
	}
	return addrs[0]
}

func GetURL(hostname string, port int, scheme string) string {
	addr := GetIPAddr(hostname)
	host := fmt.Sprintf("%s:%d", addr, port)
	u := &url.URL{
		Scheme: scheme,
		Host:   host,
	}
	return u.String()
}

func GetEthURL(hostname string) string {
	return GetURL(hostname, HttpPort, "http")
}

func GetEthWsURL(hostname string) string {
	return GetURL(hostname, WsPort, "ws")
}
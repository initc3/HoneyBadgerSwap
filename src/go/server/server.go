package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go/server/lib"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"math/big"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	prog      = "./malicious-shamir-party.x"
	players   = "4"
	threshold = "1"
	mpcPort   = "5000"
	blsPrime  = "52435875175126190479447740508185965837690552500527637822603658699938581184513"
	nshares   = 1000
	batchSize       = 2
	returnPriceInterval = 10

	confirmation = 2
	blockTime = 5
)

var (
	network        string
	serverID       string
	conn           *ethclient.Client
	server         *bind.TransactOpts
	mutexTask      = &sync.Mutex{}
	leaderHostname string
)

func genInputmask() {
	tot := utils.GetInputmaskCnt(network, conn)
	for true {
		cnt := utils.GetInputmaskCnt(network, conn)

		if cnt+100 > tot {
			go func() {
				fmt.Printf("Generating new inputmasks...\n")

				cmd := exec.Command("./random-shamir.x", "-i", serverID, "-N", players, "-T", threshold, "--nshares", strconv.Itoa(nshares), "--host", leaderHostname)
				utils.ExecCmd(cmd)

				cmd = exec.Command("python3", "-m", "honeybadgerswap.server.proc_inputmask", serverID, strconv.Itoa(int(tot)))
				utils.ExecCmd(cmd)

				tot += nshares
				fmt.Printf("Total inputmask number: %v\n", tot)
			}()
		}

		time.Sleep(30 * time.Second)
	}
}

func watch() {
	ctx := context.Background()

	blkNum, _ := conn.BlockNumber(ctx)
	for true{
		curBlockNum, _ := conn.BlockNumber(ctx)
		fmt.Println("curBlockNum", curBlockNum)
		if curBlockNum - blkNum > confirmation {
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(blkNum)),
				ToBlock: big.NewInt(int64(curBlockNum - confirmation)),
				Addresses: []common.Address{utils.HbswapAddr[network]},
			}
			logs, _ := conn.FilterLogs(ctx, query)
			for _, log := range logs {
				switch log.Topics[0].Hex() {
				case utils.SecretDeposit:
					oce := utils.ParseSecretDeposit(network, conn, log)

					mutexTask.Lock()

					fmt.Printf("**** SecretDeposit ****\n")

					token := strings.ToLower(oce.Token.Hex())
					user := strings.ToLower(oce.User.Hex())
					amt := oce.Amt.String() // fix

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.secret_deposit", serverID, token, user, amt)
					utils.ExecCmd(cmd)

					mutexTask.Unlock()

				case utils.SecretWithdraw:
					oce := utils.ParseSecretWithdraw(network, conn, log)

					mutexTask.Lock()

					fmt.Printf("**** SecretWithdraw ****\n")

					seq := oce.Seq
					token := strings.ToLower(oce.Token.Hex())
					user := strings.ToLower(oce.User.Hex())
					amt := oce.Amt.String() // fix

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.secret_withdraw_set_data", serverID, user, token, amt)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_secret_withdraw")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.secret_withdraw_org_data", serverID, token, user, amt)
					stdout := utils.ExecCmd(cmd)
					enough, _ := strconv.Atoi(stdout[:1])
					if enough == 1 {
						utils.Consent(network, conn, server, seq)
					}

					mutexTask.Unlock()

				case utils.InitPool:
					oce := utils.ParseInitPool(network, conn, log)

					mutexTask.Lock()

					fmt.Printf("**** InitPool ****\n")

					user := strings.ToLower(oce.User.Hex())
					tokenA := strings.ToLower(oce.TokenA.Hex())
					tokenB := strings.ToLower(oce.TokenB.Hex())
					amtA := oce.AmtA.String() // fix
					amtB := oce.AmtB.String() // fix

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.init_pool_set_data", serverID, user, tokenA, tokenB, amtA, amtB)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_init_pool")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.init_pool_org_data", serverID, tokenA, tokenB, user, amtA, amtB)
					stdout := utils.ExecCmd(cmd)
					outputs := strings.Split(stdout[:len(stdout) - 1], "\n")
					validOrder, _ := strconv.Atoi(outputs[0])
					if validOrder == 1 {
						price := outputs[1]
						utils.UpdatePrice(network, conn, server, common.HexToAddress(tokenA), common.HexToAddress(tokenB), big.NewInt(0), price)
					}

					mutexTask.Unlock()

				case utils.AddLiquidity:
					oce := utils.ParseAddLiquidity(network, conn, log)

					mutexTask.Lock()

					fmt.Printf("**** AddLiquidity ****\n")

					user := strings.ToLower(oce.User.Hex())
					tokenA := strings.ToLower(oce.TokenA.Hex())
					tokenB := strings.ToLower(oce.TokenB.Hex())
					idxA := oce.IdxA.String()
					idxB := oce.IdxB.String()
					maskedAmtA := oce.MaskedAmtA.String()
					maskedAmtB := oce.MaskedAmtB.String()

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.add_liquidity_set_data", serverID, user, tokenA, tokenB, idxA, maskedAmtA, idxB, maskedAmtB)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_add_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.add_liquidity_org_data", serverID, user, tokenA, tokenB, user)
					utils.ExecCmd(cmd)

					mutexTask.Unlock()

				case utils.RemoveLiquidity:
					oce := utils.ParseRemoveLiquidity(network, conn, log)

					mutexTask.Lock()

					fmt.Printf("**** RemoveLiquidity ****\n")

					user := strings.ToLower(oce.User.Hex())
					tokenA := strings.ToLower(oce.TokenA.Hex())
					tokenB := strings.ToLower(oce.TokenB.Hex())
					idx := oce.Idx.String()
					maskedAmt := oce.Idx.String()

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.remove_liquidity_set_data", serverID, user, tokenA, tokenB, idx, maskedAmt)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_remove_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.remove_liquidity_org_data", serverID, user, tokenA, tokenB)
					stdout := utils.ExecCmd(cmd)
					zeroTotalLT, _ := strconv.Atoi(stdout[:1])
					if zeroTotalLT == 1 {
						utils.ResetPrice(network, conn, server, common.HexToAddress(tokenA), common.HexToAddress(tokenB))
					}

					mutexTask.Unlock()

				case utils.Trade:
					oce := utils.ParseTrade(network, conn, log)

					mutexTask.Lock()

					fmt.Printf("**** Trade ****\n")

					tradeSeq := oce.TradeSeq.String()
					user := strings.ToLower(oce.User.Hex())
					tokenA := strings.ToLower(oce.TokenA.Hex())
					tokenB := strings.ToLower(oce.TokenB.Hex())
					idxA := oce.IdxA.String()
					idxB := oce.IdxB.String()
					maskedAmtA := oce.MaskedAmtA.String()
					maskedAmtB := oce.MaskedAmtB.String()

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.trade_set_data", serverID, user, tokenA, tokenB, idxA, maskedAmtA, idxB, maskedAmtB)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_trade")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.trade_org_data", serverID, user, tokenA, tokenB)
					stdout := utils.ExecCmd(cmd)
					outputs := strings.Split(stdout, "\n")
					orderSucceed, _ := strconv.Atoi(outputs[0][:1])

					time.Sleep(returnPriceInterval * time.Second)

					balanceA := outputs[1]
					balanceB := outputs[2]
					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.trade_update_balance", serverID, user, tokenA, tokenB, balanceA, balanceB)
					utils.ExecCmd(cmd)

					if orderSucceed == 1 {
						changeB := outputs[3]
						changeA := outputs[4]

						cmd := exec.Command("python3", "-m", "honeybadgerswap.server.calc_individual_price_set_data", serverID, changeB, changeA, tokenA, tokenB)
						utils.ExecCmd(cmd)

						cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_calc_individual_price")
						utils.ExecCmd(cmd)

						cmd = exec.Command("python3", "-m", "honeybadgerswap.server.calc_individual_price_org_data", serverID, tokenA, tokenB, tradeSeq)
						utils.ExecCmd(cmd)

					} else {
						cmd = exec.Command("python3", "-m", "honeybadgerswap.server.set_price_zero", serverID, tradeSeq)
						utils.ExecCmd(cmd)
					}

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.calc_batch_price_set_data", serverID, tokenA, tokenB)
					stdout = utils.ExecCmd(cmd)
					cnt, _ := strconv.ParseFloat(strings.Split(stdout[:len(stdout) - 1], " ")[1], 32)
					fmt.Println("cnt", cnt)

					if cnt >= batchSize {
						cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_calc_batch_price")
						stdout := utils.ExecCmd(cmd)

						cmd = exec.Command("python3", "-m", "honeybadgerswap.server.calc_batch_price_org_data", serverID, tokenA, tokenB)
						stdout = utils.ExecCmd(cmd)
						batchPrice := stdout[:len(stdout) - 1]
						fmt.Println(batchPrice)
						seq, _ := strconv.Atoi(tradeSeq)
						utils.UpdatePrice(network, conn, server, common.HexToAddress(tokenA), common.HexToAddress(tokenB), big.NewInt(int64(seq)), batchPrice)
					}

					mutexTask.Unlock()
				}
			}
			blkNum = curBlockNum - confirmation + 1
		}
		time.Sleep(blockTime * time.Second)
	}
}

func main() {
	// parse cmd line arguments/flags
	var configfile string
	flag.StringVar(&configfile, "config", "/opt/hbswap/conf/server.toml", "Config file. Default: /opt/hbswap/conf/server.toml")
	flag.StringVar(&serverID, "id", "0", "Server id. Default: 0")
	flag.Parse()

	// parse config file
	config := lib.ParseServerConfig(configfile)

	network = config.EthNode.Network
	ethHostname := config.EthNode.Hostname
	leaderHostname = config.LeaderHostname
	fmt.Println("Eth network: ", network)
	fmt.Println("Eth hostname: ", ethHostname)
	fmt.Println("Leader hostname: ", leaderHostname)

	fmt.Printf("Starting mpc server %v\n", serverID)
	server = utils.GetAccount(fmt.Sprintf("server_%s", serverID))

	var wsUrl string
	if network == "privatenet" {
		wsUrl = utils.GetEthWsURL(ethHostname)
	} else {
		wsUrl = utils.TestnetWsEndpoint
	}
	conn = utils.GetEthClient(wsUrl)

	//TODO: deleting this after testing
	if serverID == "0" && network == "testnet"{
		utils.ResetPrice(network, conn, server, utils.EthAddr, utils.TokenAddrs[network][0])
		utils.ResetBalance(network, conn, server, utils.EthAddr, utils.UserAddr)
		for _, tokenAddr := range utils.TokenAddrs[network] {
			utils.ResetBalance(network, conn, server, tokenAddr, utils.UserAddr)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go genInputmask()
	go watch()
	wg.Wait()
}

package main

import (
	"container/heap"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/HoneyBadgerSwap/src/go/server/lib"
	"github.com/initc3/HoneyBadgerSwap/src/go/utils"
	"github.com/initc3/HoneyBadgerSwap/src/go_bindings/hbswap"
	"log"
	"math"
	"math/big"
	"os"
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
	prepDir   = "/opt/hbswap/preprocessing-data"
	batch     = 1
)

var (
	network        string
	serverID       string
	conn           *ethclient.Client
	server         *bind.TransactOpts
	pq             utils.PriorityQueue
	mutexPQ        = &sync.Mutex{}
	eventSet       map[utils.EventID]bool
	leaderHostname string
)

func checkBalance(token string, user string, amt string) int {
	cmd := exec.Command("python3", "-m", "honeybadgerswap.server.check_balance_set_data", serverID, token, user, amt)
	utils.ExecCmd(cmd)

	cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_check_balance")
	stdout := utils.ExecCmd(cmd)

	cmd = exec.Command("python3", "-m", "honeybadgerswap.server.check_balance_org_data", serverID)
	stdout = utils.ExecCmd(cmd)
	enoughBalance, _ := strconv.Atoi(stdout[:1])
	fmt.Printf("enoughBalance %v\n", enoughBalance)

	return enoughBalance
}

func updateBalance(token string, user string, amt string, flag string) {
	cmd := exec.Command("python3", "-m", "honeybadgerswap.server.update_balance", serverID, token, user, amt, flag)
	utils.ExecCmd(cmd)
}

func genInputmask() {
	tot := utils.GetInputmaskCnt(network, conn)
	for true {
		cnt := utils.GetInputmaskCnt(network, conn)

		if cnt+100 > tot {
			fmt.Printf("Generating new inputmasks...\n")

			cmd := exec.Command("./random-shamir.x", "-i", serverID, "-N", players, "-T", threshold, "--nshares", strconv.Itoa(nshares), "--host", leaderHostname)
			utils.ExecCmd(cmd)

			cmd = exec.Command("python3", "-m", "honeybadgerswap.server.proc_inputmask", serverID, strconv.Itoa(int(tot)))
			utils.ExecCmd(cmd)

			tot += nshares
			fmt.Printf("Total inputmask number: %v\n", tot)
		}

		time.Sleep(30 * time.Second)
	}
}

func watch() {
	hbswapInstance, err := hbswap.NewHbSwap(utils.HbswapAddr[network], conn)

	initPoolChannel := make(chan *hbswap.HbSwapInitPool)
	initPoolSub, err := hbswapInstance.WatchInitPool(nil, initPoolChannel)
	if err != nil {
		log.Fatal("watch InitPool err:", err)
	}

	AddLiquidityChannel := make(chan *hbswap.HbSwapAddLiquidity)
	AddLiquiditySub, err := hbswapInstance.WatchAddLiquidity(nil, AddLiquidityChannel)
	if err != nil {
		log.Fatal("watch AddLiquidity err:", err)
	}

	RemoveLiquidityChannel := make(chan *hbswap.HbSwapRemoveLiquidity)
	RemoveLiquiditySub, err := hbswapInstance.WatchRemoveLiquidity(nil, RemoveLiquidityChannel)
	if err != nil {
		log.Fatal("watch RemoveLiquidity err:", err)
	}

	tradeChannel := make(chan *hbswap.HbSwapTrade)
	tradeSub, err := hbswapInstance.WatchTrade(nil, tradeChannel)
	if err != nil {
		log.Fatal("watch Trade err:", err)
	}

	secretDepositPrepChannel := make(chan *hbswap.HbSwapSecretDeposit)
	secretDepositPrepSub, err := hbswapInstance.WatchSecretDeposit(nil, secretDepositPrepChannel)
	if err != nil {
		log.Fatal("watch LocalDepositPrep err:", err)
	}

	secretWithdrawChannel := make(chan *hbswap.HbSwapSecretWithdraw)
	secretWithdrawSub, err := hbswapInstance.WatchSecretWithdraw(nil, secretWithdrawChannel)
	if err != nil {
		log.Fatal("watch secretWithdraw err:", err)
	}

	for {
		select {
		case err := <-initPoolSub.Err():
			log.Fatal(err)
		case oce := <-initPoolChannel:
			go func() {
				fmt.Println("Push InitPool")
				task := utils.Task{
					EventID: utils.EventID{
						BlockNumber: oce.Raw.BlockNumber,
						TxIndex:     oce.Raw.TxIndex,
						LogIndex:    oce.Raw.Index,
					},
					EventName: "InitPool",
					Parameters: []string{
						strings.ToLower(oce.User.String()),
						strings.ToLower(oce.TokenA.String()),
						strings.ToLower(oce.TokenB.String()),
						oce.AmtA.String(),
						oce.AmtB.String(),
					},
				}
				mutexPQ.Lock()
				heap.Push(&pq, &task)
				mutexPQ.Unlock()
			}()

		case err := <-AddLiquiditySub.Err():
			log.Fatal(err)
		case oce := <-AddLiquidityChannel:
			go func() {
				fmt.Println("Push AddLiquidity")
				task := utils.Task{
					EventID: utils.EventID{
						BlockNumber: oce.Raw.BlockNumber,
						TxIndex:     oce.Raw.TxIndex,
						LogIndex:    oce.Raw.Index,
					},
					EventName: "AddLiquidity",
					Parameters: []string{
						strings.ToLower(oce.User.String()),
						strings.ToLower(oce.TokenA.String()),
						strings.ToLower(oce.TokenB.String()),
						oce.AmtA.String(),
						oce.AmtB.String(),
					},
				}
				mutexPQ.Lock()
				heap.Push(&pq, &task)
				mutexPQ.Unlock()
			}()

		case err := <-RemoveLiquiditySub.Err():
			log.Fatal(err)
		case oce := <-RemoveLiquidityChannel:
			go func() {
				fmt.Println("Push RemoveLiquidity")
				task := utils.Task{
					EventID: utils.EventID{
						BlockNumber: oce.Raw.BlockNumber,
						TxIndex:     oce.Raw.TxIndex,
						LogIndex:    oce.Raw.Index,
					},
					EventName: "RemoveLiquidity",
					Parameters: []string{
						strings.ToLower(oce.User.String()),
						strings.ToLower(oce.TokenA.String()),
						strings.ToLower(oce.TokenB.String()),
						oce.Amt.String(),
					},
				}
				mutexPQ.Lock()
				heap.Push(&pq, &task)
				mutexPQ.Unlock()
			}()

		case err := <-tradeSub.Err():
			log.Fatal(err)
		case oce := <-tradeChannel:
			go func() {
				fmt.Println("Push Trade")
				task := utils.Task{
					EventID: utils.EventID{
						BlockNumber: oce.Raw.BlockNumber,
						TxIndex:     oce.Raw.TxIndex,
						LogIndex:    oce.Raw.Index,
					},
					EventName: "Trade",
					Parameters: []string{
						oce.TradeSeq.String(),
						strings.ToLower(oce.User.String()),
						strings.ToLower(oce.TokenA.String()),
						strings.ToLower(oce.TokenB.String()),
						oce.IdxA.String(),
						oce.IdxB.String(),
						oce.MaskedA.String(),
						oce.MaskedB.String(),
					},
				}
				mutexPQ.Lock()
				heap.Push(&pq, &task)
				mutexPQ.Unlock()
			}()

		case err := <-secretDepositPrepSub.Err():
			log.Fatal(err)
		case oce := <-secretDepositPrepChannel:
			go func() {
				fmt.Println("Push SecretDeposit")
				task := utils.Task{
					EventID: utils.EventID{
						BlockNumber: oce.Raw.BlockNumber,
						TxIndex:     oce.Raw.TxIndex,
						LogIndex:    oce.Raw.Index,
					},
					EventName: "SecretDeposit",
					Parameters: []string{
						strings.ToLower(oce.Token.String()),
						strings.ToLower(oce.User.String()),
						oce.Amt.String(),
					},
				}
				mutexPQ.Lock()
				heap.Push(&pq, &task)
				mutexPQ.Unlock()

			}()

		case err := <-secretWithdrawSub.Err():
			log.Fatal(err)
		case oce := <-secretWithdrawChannel:
			go func() {
				fmt.Println("Push SecretWithdraw")
				task := utils.Task{
					EventID: utils.EventID{
						BlockNumber: oce.Raw.BlockNumber,
						TxIndex:     oce.Raw.TxIndex,
						LogIndex:    oce.Raw.Index,
					},
					EventName: "SecretWithdraw",
					Parameters: []string{
						oce.Seq.String(),
						strings.ToLower(oce.Token.String()),
						strings.ToLower(oce.User.String()),
						oce.Amt.String(),
					},
				}
				mutexPQ.Lock()
				heap.Push(&pq, &task)
				mutexPQ.Unlock()
			}()

		}
	}
}

func processTasks() {
	for true {
		for pq.Len() > 0 {
			mutexPQ.Lock()
			task := heap.Pop(&pq).(*utils.Task)
			mutexPQ.Unlock()

			if _, ok := eventSet[task.EventID]; ok {
				continue
			}
			eventSet[task.EventID] = true

			switch task.EventName {
			case "InitPool":
				fmt.Printf("**** InitPool ****\n")

				user := task.Parameters[0]
				tokenA := task.Parameters[1]
				tokenB := task.Parameters[2]
				amtA := task.Parameters[3]
				amtB := task.Parameters[4]

				if checkBalance(tokenA, user, amtA) == 1 && checkBalance(tokenB, user, amtB) == 1 {
					_amtA, _ := strconv.Atoi(amtA)
					_amtB, _ := strconv.Atoi(amtB)
					amtLiquidity := fmt.Sprintf("%f", math.Sqrt(float64(_amtA*_amtB)))
					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.init_pool", serverID, tokenA, tokenB, amtA, amtB, amtLiquidity)
					utils.ExecCmd(cmd)

					updateBalance(tokenA, user, fmt.Sprintf("-%s", amtA), "1")
					updateBalance(tokenB, user, fmt.Sprintf("-%s", amtB), "1")
					updateBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, amtLiquidity, "1")

					_tokenA := common.HexToAddress(tokenA)
					_tokenB := common.HexToAddress(tokenB)
					price := fmt.Sprintf("%f", float64(_amtB)/float64(_amtA))

					utils.UpdatePrice(network, conn, server, _tokenA, _tokenB, big.NewInt(0), price)
				}

			case "AddLiquidity":
				fmt.Printf("**** AddLiquidity ****\n")

				user := task.Parameters[0]
				tokenA := task.Parameters[1]
				tokenB := task.Parameters[2]
				amtA := task.Parameters[3]
				amtB := task.Parameters[4]

				if checkBalance(tokenA, user, amtA) == 1 && checkBalance(tokenB, user, amtB) == 1 {
					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.add_liquidity_set_data", serverID, user, tokenA, tokenB, amtA, amtB)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_add_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.add_liquidity_org_data", serverID, tokenA, tokenB)
					stdout := utils.ExecCmd(cmd)
					amts := strings.Split(strings.Split(stdout, "\n")[0], " ")
					amtA = amts[0]
					amtB = amts[1]
					amtLiquidity := amts[2]
					fmt.Printf("amt_A %s amt_B %s amt %s\n", amtA, amtB, amtLiquidity)

					updateBalance(tokenA, user, fmt.Sprintf("-%s", amtA), "0")
					updateBalance(tokenB, user, fmt.Sprintf("-%s", amtB), "0")
					updateBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, amtLiquidity, "0")
				}

			case "RemoveLiquidity":
				fmt.Printf("**** RemoveLiquidity ****\n")

				user := task.Parameters[0]
				tokenA := task.Parameters[1]
				tokenB := task.Parameters[2]
				amtLiquidity := task.Parameters[3]

				if checkBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, amtLiquidity) == 1 {
					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.remove_liquidity_set_data", serverID, user, tokenA, tokenB, amtLiquidity)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_remove_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.remove_liquidity_org_data", serverID, tokenA, tokenB, amtLiquidity)
					stdout := utils.ExecCmd(cmd)
					amts := strings.Split(strings.Split(stdout, "\n")[0], " ")
					amtA := amts[0]
					amtB := amts[1]
					fmt.Printf("amt_A %s amt_B %s\n", amtA, amtB)

					updateBalance(tokenA, user, amtA, "0")
					updateBalance(tokenB, user, amtB, "0")
					updateBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, fmt.Sprintf("-%s", amtLiquidity), "1")
				}

			case "SecretDeposit":
				fmt.Printf("**** SecretDeposit ****\n")

				token := task.Parameters[0]
				user := task.Parameters[1]
				amt := task.Parameters[2]

				updateBalance(token, user, amt, "1")

			case "Trade":
				fmt.Printf("**** Trade ****\n")

				tradeSeq := task.Parameters[0]
				user := task.Parameters[1]
				tokenA := task.Parameters[2]
				tokenB := task.Parameters[3]
				idxA := task.Parameters[4]
				idxB := task.Parameters[5]
				maskedA := task.Parameters[6]
				maskedB := task.Parameters[7]

				cmd := exec.Command("python3", "-m", "honeybadgerswap.server.trade_set_data", serverID, user, tokenA, tokenB, idxA, idxB, maskedA, maskedB)
				utils.ExecCmd(cmd)
				os.RemoveAll(fmt.Sprintf(prepDir))
				os.Mkdir(fmt.Sprintf(prepDir), 0777)

				cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_trade")
				utils.ExecCmd(cmd)

				cmd = exec.Command("python3", "-m", "honeybadgerswap.server.trade_org_data", serverID, tokenA, tokenB, tradeSeq)
				stdout := utils.ExecCmd(cmd)
				changes := strings.Split(strings.Split(stdout, "\n")[0], " ")
				changeA := changes[0]
				changeB := changes[1]
				fmt.Printf("changeA %s changeB %s\n", changeA, changeB)

				updateBalance(tokenA, user, changeA, "0")
				updateBalance(tokenB, user, changeB, "0")

				seq, _ := strconv.Atoi(tradeSeq)
				if seq%batch == 0 {
					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.calc_price_set_data", serverID, tokenA, tokenB)
					utils.ExecCmd(cmd)
					os.RemoveAll(fmt.Sprintf(prepDir))
					os.Mkdir(fmt.Sprintf(prepDir), 0777)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_calc_price")
					stdout := utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.calc_price_org_data", serverID)
					stdout = utils.ExecCmd(cmd)
					price := strings.Split(stdout, "\n")[0]
					fmt.Printf("avg_price %s\n", price)
					_tokenA := common.HexToAddress(tokenA)
					_tokenB := common.HexToAddress(tokenB)
					utils.UpdatePrice(network, conn, server, _tokenA, _tokenB, big.NewInt(int64(seq)), price)
				}

			case "SecretWithdraw":
				fmt.Printf("**** SecretWithdraw ****\n")

				seq := task.Parameters[0]
				token := task.Parameters[1]
				user := task.Parameters[2]
				amt := task.Parameters[3]

				if checkBalance(token, user, amt) == 1 {
					utils.Consent(network, conn, server, utils.StrToBig(seq))
					updateBalance(token, user, fmt.Sprintf("-%s", amt), "1")
				}

			}
		}
		time.Sleep(time.Second)
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
	eventSet = map[utils.EventID]bool{}

	var wsUrl string
	if network == "privatenet" {
		wsUrl = utils.GetEthWsURL(ethHostname)
	} else {
		wsUrl = utils.TestnetWsEndpoint
		//wsUrl = config.EthNode.WsEndpoint
	}
	conn = utils.GetEthClient(wsUrl)

	var wg sync.WaitGroup
	wg.Add(1)
	go genInputmask()
	go watch()
	go processTasks()
	wg.Wait()
}

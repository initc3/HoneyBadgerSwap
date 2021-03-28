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
	//prepDir   = "/opt/hbswap/preprocessing-data"
	batchSize       = 2
	returnPriceInterval = 10
)

var (
	network        string
	serverID       string
	conn           *ethclient.Client
	server         *bind.TransactOpts
	pq             utils.PriorityQueue
	mutexPQ        = &sync.Mutex{}
	mutexTask      = &sync.Mutex{}
	eventSet       = map[utils.EventID]bool{}
	leaderHostname string
	tokenPairs     = map[string]bool{}
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
						oce.IdxA.String(),
						oce.IdxB.String(),
						oce.MaskedAmtA.String(),
						oce.MaskedAmtB.String(),
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
						oce.Idx.String(),
						oce.MaskedAmt.String(),
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
						oce.MaskedAmtA.String(),
						oce.MaskedAmtB.String(),
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
				go func() {
					mutexTask.Lock()

					fmt.Printf("**** InitPool ****\n")

					user := task.Parameters[0]
					tokenA := task.Parameters[1]
					tokenB := task.Parameters[2]
					amtA := task.Parameters[3] // fix
					amtB := task.Parameters[4] // fix

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
				}()

			case "AddLiquidity":
				go func() {
					mutexTask.Lock()

					fmt.Printf("**** AddLiquidity ****\n")

					user := task.Parameters[0]
					tokenA := task.Parameters[1]
					tokenB := task.Parameters[2]
					idxA := task.Parameters[3]
					idxB := task.Parameters[4]
					maskedAmtA := task.Parameters[5]
					maskedAmtB := task.Parameters[6]

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.add_liquidity_set_data", serverID, user, tokenA, tokenB, idxA, maskedAmtA, idxB, maskedAmtB)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_add_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.add_liquidity_org_data", serverID, user, tokenA, tokenB, user)
					utils.ExecCmd(cmd)

					mutexTask.Unlock()
				}()

			case "RemoveLiquidity":
				go func() {
					mutexTask.Lock()

					fmt.Printf("**** RemoveLiquidity ****\n")

					user := task.Parameters[0]
					tokenA := task.Parameters[1]
					tokenB := task.Parameters[2]
					idx := task.Parameters[3]
					maskedAmt := task.Parameters[4]

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
				}()

			case "SecretDeposit":
				go func() {
					mutexTask.Lock()

					fmt.Printf("**** SecretDeposit ****\n")

					token := task.Parameters[0]
					user := task.Parameters[1]
					amt := task.Parameters[2] // fix

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.secret_deposit", serverID, token, user, amt)
					utils.ExecCmd(cmd)

					mutexTask.Unlock()
				}()

			case "Trade":
				go func() {
					mutexTask.Lock()

					fmt.Printf("**** Trade ****\n")

					tradeSeq := task.Parameters[0]
					user := task.Parameters[1]
					tokenA := task.Parameters[2]
					tokenB := task.Parameters[3]
					idxA := task.Parameters[4]
					idxB := task.Parameters[5]
					maskedAmtA := task.Parameters[6]
					maskedAmtB := task.Parameters[7]

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
				}()

			case "SecretWithdraw":
				go func() {
					mutexTask.Lock()

					fmt.Printf("**** SecretWithdraw ****\n")

					seq := task.Parameters[0]
					token := task.Parameters[1]
					user := task.Parameters[2]
					amt := task.Parameters[3]

					cmd := exec.Command("python3", "-m", "honeybadgerswap.server.secret_withdraw_set_data", serverID, user, token, amt)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leaderHostname, "hbswap_secret_withdraw")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "-m", "honeybadgerswap.server.secret_withdraw_org_data", serverID, token, user, amt)
					stdout := utils.ExecCmd(cmd)
					enough, _ := strconv.Atoi(stdout[:1])
					if enough == 1 {
						utils.Consent(network, conn, server, utils.StrToBig(seq))
					}

					mutexTask.Unlock()
				}()

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

	var wsUrl string
	if network == "privatenet" {
		wsUrl = utils.GetEthWsURL(ethHostname)
	} else {
		wsUrl = utils.TestnetWsEndpoint
		//wsUrl = config.EthNode.WsEndpoint
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
	go processTasks()
	wg.Wait()
}

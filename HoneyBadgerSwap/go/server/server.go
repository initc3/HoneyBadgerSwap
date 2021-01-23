package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go_bindings/hbswap"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	prog          = "./malicious-shamir-party.x"
	players       = "4"
	threshold     = "1"
	mpcPort       = "5000"
	blsPrime      = "52435875175126190479447740508185965837690552500527637822603658699938581184513"
	nshares       = 1000
	checkInterval = 10
	prep_dir      = "/opt/hbswap/preprocessing-data"
)

var (
	serverID string
	conn     *ethclient.Client
	server   *bind.TransactOpts
	prevTime = int64(0)
)

func checkBalance(token string, user string, amt string, leader_hostname string) int {
	cmd := exec.Command("python3", "Scripts/hbswap/python/server/check_balance_set_data.py", serverID, token, user, amt)
	utils.ExecCmd(cmd)

	cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leader_hostname, "hbswap_check_balance")
	stdout := utils.ExecCmd(cmd)

	cmd = exec.Command("python3", "Scripts/hbswap/python/server/check_balance_org_data.py", serverID)
	stdout = utils.ExecCmd(cmd)
	enoughBalance, _ := strconv.Atoi(stdout[:1])
	fmt.Printf("enoughBalance %v\n", enoughBalance)

	return enoughBalance
}

func updateBalance(token string, user string, amt string, flag string) {
	cmd := exec.Command("python3", "Scripts/hbswap/python/server/update_balance.py", serverID, token, user, amt, flag)
	utils.ExecCmd(cmd)
}

func genInputmask(leader_hostname string) {
	tot := utils.GetInputmaskCnt(conn)
	for true {
		cnt := utils.GetInputmaskCnt(conn)

		if cnt+100 > tot {
			fmt.Printf("Generating new inputmasks...\n")

			cmd := exec.Command("./random-shamir.x", "-i", serverID, "-N", players, "-T", threshold, "--nshares", strconv.Itoa(nshares), "--host", leader_hostname)
			utils.ExecCmd(cmd)

			cmd = exec.Command("python3", "Scripts/hbswap/python/server/proc_inputmask.py", serverID, strconv.Itoa(int(tot)))
			utils.ExecCmd(cmd)

			tot += nshares
			fmt.Printf("Total inputmask number: %v\n", tot)
		}

		time.Sleep(30 * time.Second)
	}
}

func watch(leader_hostname string) {
	hbswapInstance, err := hbswap.NewHbSwap(utils.HbswapAddr, conn)

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
				fmt.Printf("**** InitPool ****\n")

				user := oce.User.String()
				tokenA := oce.TokenA.String()
				tokenB := oce.TokenB.String()
				amtA := oce.AmtA.String()
				amtB := oce.AmtB.String()

				if checkBalance(tokenA, user, amtA, leader_hostname) == 1 && checkBalance(tokenB, user, amtB, leader_hostname) == 1 {
					amtLiquidity := fmt.Sprintf("%f", math.Sqrt(float64(oce.AmtA.Int64()*oce.AmtB.Int64())))
					cmd := exec.Command("python3", "Scripts/hbswap/python/server/init_pool.py", serverID, tokenA, tokenB, amtA, amtB, amtLiquidity)
					utils.ExecCmd(cmd)

					updateBalance(tokenA, user, fmt.Sprintf("-%s", amtA), "1")
					updateBalance(tokenB, user, fmt.Sprintf("-%s", amtB), "1")
					updateBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, amtLiquidity, "1")
				}
			}()

		case err := <-AddLiquiditySub.Err():
			log.Fatal(err)
		case oce := <-AddLiquidityChannel:
			go func() {
				fmt.Printf("**** AddLiquidity ****\n")

				user := oce.User.String()
				tokenA := oce.TokenA.String()
				tokenB := oce.TokenB.String()
				amtA := oce.AmtA.String()
				amtB := oce.AmtB.String()

				if checkBalance(tokenA, user, amtA, leader_hostname) == 1 && checkBalance(tokenB, user, amtB, leader_hostname) == 1 {
					cmd := exec.Command("python3", "Scripts/hbswap/python/server/add_liquidity_set_data.py", serverID, user, tokenA, tokenB, amtA, amtB)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leader_hostname, "hbswap_add_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "Scripts/hbswap/python/server/add_liquidity_org_data.py", serverID, tokenA, tokenB)
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
			}()

		case err := <-RemoveLiquiditySub.Err():
			log.Fatal(err)
		case oce := <-RemoveLiquidityChannel:
			go func() {
				fmt.Printf("**** RemoveLiquidity ****\n")

				user := oce.User.String()
				tokenA := oce.TokenA.String()
				tokenB := oce.TokenB.String()
				amtLiquidity := oce.Amt.String()

				if checkBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, amtLiquidity, leader_hostname) == 1 {
					cmd := exec.Command("python3", "Scripts/hbswap/python/server/remove_liquidity_set_data.py", serverID, user, tokenA, tokenB, amtLiquidity)
					utils.ExecCmd(cmd)

					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leader_hostname, "hbswap_remove_liquidity")
					utils.ExecCmd(cmd)

					cmd = exec.Command("python3", "Scripts/hbswap/python/server/remove_liquidity_org_data.py", serverID, tokenA, tokenB, amtLiquidity)
					stdout := utils.ExecCmd(cmd)
					amts := strings.Split(strings.Split(stdout, "\n")[0], " ")
					amtA := amts[0]
					amtB := amts[1]
					fmt.Printf("amt_A %s amt_B %s\n", amtA, amtB)

					updateBalance(tokenA, user, amtA, "0")
					updateBalance(tokenB, user, amtB, "0")
					updateBalance(fmt.Sprintf("%s+%s", tokenA, tokenB), user, fmt.Sprintf("-%s", amtLiquidity), "1")
				}
			}()

		case err := <-tradeSub.Err():
			log.Fatal(err)
		case oce := <-tradeChannel:
			go func() {
				fmt.Printf("**** Trade ****\n")

				//if serverID != "0" {
				//	time.Sleep(1 * time.Second)
				//}

				user := oce.User.Hex()
				tokenA := oce.TokenA.String()
				tokenB := oce.TokenB.String()

				cmd := exec.Command("python3", "Scripts/hbswap/python/server/trade_set_data.py", serverID, user, tokenA, tokenB, oce.IdxA.String(), oce.IdxB.String(), oce.MaskedA.String(), oce.MaskedB.String())
				utils.ExecCmd(cmd)
				os.RemoveAll(fmt.Sprintf(prep_dir))
				os.Mkdir(fmt.Sprintf(prep_dir), 0777)
				cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leader_hostname, "hbswap_trade")
				utils.ExecCmd(cmd)

				cmd = exec.Command("python3", "Scripts/hbswap/python/server/trade_org_data.py", serverID, tokenA, tokenB, oce.TradeSeq.String())
				stdout := utils.ExecCmd(cmd)
				changes := strings.Split(strings.Split(stdout, "\n")[0], " ")
				changeA := changes[0]
				changeB := changes[1]
				fmt.Printf("changeA %s changeB %s\n", changeA, changeB)

				updateBalance(tokenA, user, changeA, "0")
				updateBalance(tokenB, user, changeB, "0")

				if time.Now().Unix()-prevTime > checkInterval {
					cmd = exec.Command("python3", "Scripts/hbswap/python/server/calc_price_set_data.py", serverID, tokenA, tokenB)
					utils.ExecCmd(cmd)

					os.RemoveAll(fmt.Sprintf(prep_dir))
					os.Mkdir(fmt.Sprintf(prep_dir), 0777)
					cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leader_hostname, "hbswap_calc_price")
					stdout := utils.ExecCmd(cmd)
					price := strings.Split(stdout, "\n")[1]
					fmt.Printf("avg_price %s\n", price)

					if serverID == "0" {
						utils.UpdatePrice(conn, server, oce.TokenA, oce.TokenB, price)
					} else {
						prevBlockNum := utils.GetUpdateTime(conn, oce.TokenA, oce.TokenB)
						for true {
							time.Sleep(time.Second)
							curBlockNum := utils.GetUpdateTime(conn, oce.TokenA, oce.TokenB)
							if curBlockNum > prevBlockNum {
								break
							}
						}
					}

					prevTime = time.Now().Unix()
				}
			}()

		case err := <-secretDepositPrepSub.Err():
			log.Fatal(err)
		case oce := <-secretDepositPrepChannel:
			go func() {
				fmt.Printf("**** SecretDeposit ****\n")

				updateBalance(oce.Token.Hex(), oce.User.Hex(), oce.Amt.String(), "1")
			}()

		case err := <-secretWithdrawSub.Err():
			log.Fatal(err)
		case oce := <-secretWithdrawChannel:
			go func() {
				fmt.Printf("**** SecretWithdraw ****\n")

				if checkBalance(oce.Token.String(), oce.User.String(), oce.Amt.String(), leader_hostname) == 1 {
					utils.Consent(conn, server, oce.Seq)
					updateBalance(oce.Token.Hex(), oce.User.Hex(), fmt.Sprintf("-%s", oce.Amt.String()), "1")
				}
			}()

		}
	}
}

func main() {
	serverID = os.Args[1]
	fmt.Printf("Starting mpc server %v\n", serverID)

	ethHostname := os.Args[2]
	wsUrl := utils.GetEthWsURL(ethHostname)
	conn = utils.GetEthClient(wsUrl)

	leader_hostname := os.Args[3]

	server = utils.GetAccount(fmt.Sprintf("server_%s", serverID))

	var wg sync.WaitGroup
	wg.Add(1)
	go genInputmask(leader_hostname)
	go watch(leader_hostname)
	wg.Wait()
}

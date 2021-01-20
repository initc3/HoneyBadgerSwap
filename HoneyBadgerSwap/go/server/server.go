package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go_bindings/hbswap"
	"log"
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
	sz        = 32
	nshares   = 1000
	interval  = 10
	prep_dir  = "/opt/hbswap/preprocessing-data"
)

var (
	serverID string
	mut      sync.Mutex
	conn     *ethclient.Client
	server   *bind.TransactOpts
	prevTime int64
)

//func dbPut(key string, value []byte) {
//	mut.Lock()
//	db, _ := leveldb.OpenFile(fmt.Sprintf("Scripts/hbswap/db/server%s", serverID), nil)
//	err := db.Put([]byte(key), value, nil)
//	if err != nil {
//		fmt.Println("Error writing to database")
//	}
//	db.Close()
//	mut.Unlock()
//}

//func dbGet(key string) string {
//	mut.Lock()
//	db, _ := leveldb.OpenFile(fmt.Sprintf("Scripts/hbswap/db/server%s", serverID), nil)
//	data, err := db.Get([]byte(key), nil)
//	if err != nil {
//		fmt.Println("Error getting from database")
//	}
//	db.Close()
//	mut.Unlock()
//	return string(data)
//}

func genInputmask(leader_hostname string) {
	tot := int(utils.GetInputmaskCnt(conn).Int64())
	for true {
		cnt := utils.GetInputmaskCnt(conn)

		if int(cnt.Int64())+100 > tot {
			fmt.Printf("Generating new inputmasks...\n")

			cmd := exec.Command("./random-shamir.x", "-i", serverID, "-N", players, "-T", threshold, "--nshares", strconv.Itoa(nshares), "--host", leader_hostname)
			utils.ExecCmd(cmd)

			cmd = exec.Command("python3", "Scripts/hbswap/python/server/proc_inputmask.py", serverID, strconv.Itoa(tot))
			utils.ExecCmd(cmd)

			tot += nshares
			fmt.Printf("Total inputmask number: %v\n", tot)
		}

		time.Sleep(30 * time.Second)
	}
}

func watch(leader_hostname string) {
	hbswapInstance, err := hbswap.NewHbSwap(utils.HbswapAddr, conn)

	//tradePrepChannel := make(chan *hbswap.HbSwapTradePrep)
	//tradePrepSub, err := hbswapInstance.WatchTradePrep(nil, tradePrepChannel)
	//if err != nil {
	//	log.Fatal("watch TradePrep err:", err)
	//}

	initPoolChannel := make(chan *hbswap.HbSwapInitPool)
	initPoolSub, err := hbswapInstance.WatchInitPool(nil, initPoolChannel)
	if err != nil {
		log.Fatal("watch InitPool err:", err)
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
				fmt.Printf("****New liquidity pool...\n")

				cmd := exec.Command("python3", "Scripts/hbswap/python/server/init_pool.py", serverID, oce.TokenA.String(), oce.TokenB.String(), oce.AmtA.String(), oce.AmtB.String())
				utils.ExecCmd(cmd)
			}()

		//case err := <- tradePrepSub.Err():
		//	log.Fatal(err)
		//case oce := <-tradePrepChannel:
		//	fmt.Printf("Preparing inputmasks with for %v and %v\n", oce.IdxA, oce.IdxB)
		//
		//	_ = os.Remove(fmt.Sprintf("Persistence/Transactions-P%v.data", serverID))
		//	cmd := exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "hbswap_trade_prep")
		//	utils.ExecCmd(cmd)
		//
		//	f, err := os.Open(fmt.Sprintf("Persistence/Transactions-P%v.data", serverID))
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	share1 := make([]byte, sz)
		//	_, err = f.Read(share1)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	share2 := make([]byte, sz)
		//	_, err = f.Read(share2)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Printf("Inputmask-%v: %x\n", oce.IdxA, share1)
		//	fmt.Printf("Inputmask-%v: %x\n", oce.IdxB, share2)
		//
		//	dbPut(oce.IdxA.String(), share1)
		//	dbPut(oce.IdxB.String(), share2)

		case err := <-tradeSub.Err():
			log.Fatal(err)
		case oce := <-tradeChannel:
			go func() {
				fmt.Printf("****Trade\n")

				if serverID != "0" {
					time.Sleep(1 * time.Second)
				}

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
				println(strings.Split(stdout, "\n")[0])
				changes := strings.Split(strings.Split(stdout, "\n")[0], " ")
				fmt.Printf("change_A %s change_B %s\n", changes[0], changes[1])

				cmd = exec.Command("python3", "Scripts/hbswap/python/server/update_balance.py", serverID, tokenA, user, changes[0], "0")
				utils.ExecCmd(cmd)

				cmd = exec.Command("python3", "Scripts/hbswap/python/server/update_balance.py", serverID, tokenB, user, changes[1], "0")
				utils.ExecCmd(cmd)

				if time.Now().Unix()-prevTime > interval {
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
						prevTime := utils.GetUpdateTime(conn, oce.TokenA, oce.TokenB)
						for true {
							time.Sleep(1 * time.Second)
							curTime := utils.GetUpdateTime(conn, oce.TokenA, oce.TokenB)
							if curTime > prevTime {
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
				fmt.Printf("****SecretDeposit\n")

				cmd := exec.Command("python3", "Scripts/hbswap/python/server/update_balance.py", serverID, oce.Token.Hex(), oce.User.Hex(), oce.Amt.String(), "1")
				utils.ExecCmd(cmd)
			}()

		case err := <-secretWithdrawSub.Err():
			log.Fatal(err)
		case oce := <-secretWithdrawChannel:
			go func() {
				fmt.Printf("****SecretWithdraw\n")

				//if serverID != "0" {
				//	time.Sleep(1 * time.Second)
				//}

				cmd := exec.Command("python3", "Scripts/hbswap/python/server/withdraw_set_data.py", serverID, oce.Token.String(), oce.User.String(), oce.Amt.String())
				utils.ExecCmd(cmd)

				os.RemoveAll(fmt.Sprintf(prep_dir))
				os.Mkdir(fmt.Sprintf(prep_dir), 0777)
				cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "-P", blsPrime, "--hostname", leader_hostname, "hbswap_withdraw")
				stdout := utils.ExecCmd(cmd)

				cmd = exec.Command("python3", "Scripts/hbswap/python/server/withdraw_check.py", serverID)
				stdout = utils.ExecCmd(cmd)
				agree, _ := strconv.Atoi(stdout[:1])
				fmt.Printf("agree %v\n", agree)

				if agree == 1 {
					utils.Consent(conn, server, oce.Seq)
					cmd := exec.Command("python3", "Scripts/hbswap/python/server/update_balance.py", serverID, oce.Token.Hex(), oce.User.Hex(), fmt.Sprintf("-%s", oce.Amt.String()), "1")
					utils.ExecCmd(cmd)
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

	prevTime = 0

	var wg sync.WaitGroup
	wg.Add(1)
	go watch(leader_hostname)
	go genInputmask(leader_hostname)
	wg.Wait()
}

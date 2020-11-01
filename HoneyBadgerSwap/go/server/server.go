package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/gobingdings/hbswap"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"os/exec"
	"sync"
)

const (
	hbswapAddr = "0xF74Eb25Ab1785D24306CA6b3CBFf0D0b0817C5E2"
	prog = "./malicious-shamir-party.x"
	players = "3"
	threshold = "1"
	mpcPort = "5000"
	sz = 32
)

var (
	serverID	string
	mut 		sync.Mutex

)

func dbPut(key string, value []byte) {
	mut.Lock()
	db, _ := leveldb.OpenFile(fmt.Sprintf("Scripts/hbswap/db/server%s", serverID), nil)
	err := db.Put([]byte(key), value, nil)
	if err != nil {
		fmt.Println("Error writing to database")
	}
	db.Close()
	mut.Unlock()
}

func Watch(conn *ethclient.Client) {
	hbswapInstance, err := hbswap.NewHbSwap(common.HexToAddress(hbswapAddr), conn)

	inputmaskChannel := make(chan *hbswap.HbSwapInputmask)
	inputmaskSub, err := hbswapInstance.WatchInputmask(nil, inputmaskChannel)
	if err != nil {
		log.Fatal("watch Inputmask err:", err)
	}

	tradeChannel := make(chan *hbswap.HbSwapTrade)
	tradeSub, err := hbswapInstance.WatchTrade(nil, tradeChannel)
	if err != nil {
		log.Fatal("watch Trade err:", err)
	}

	for {
		select {
		case err := <- inputmaskSub.Err():
			log.Fatal(err)
		case oce := <-inputmaskChannel:
			fmt.Printf("Preparing inputmask with indexes %v and %v\n", oce.IdxETH, oce.IdxTOK)

			_ = os.Remove(fmt.Sprintf("Persistence/Transactions-P%v.data", serverID))
			cmd := exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "hbswap_inputmask")
			utils.ExecCmd(cmd)

			f, err := os.Open(fmt.Sprintf("Persistence/Transactions-P%v.data", serverID))
			if err != nil {
				log.Fatal(err)
			}
			share1 := make([]byte, sz)
			_, err = f.Read(share1)
			if err != nil {
				log.Fatal(err)
			}
			share2 := make([]byte, sz)
			_, err = f.Read(share2)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Inputmask-%v: %x\n", oce.IdxETH, share1)
			fmt.Printf("Inputmask-%v: %x\n", oce.IdxTOK, share2)

			dbPut(oce.IdxETH.String(), share1)
			dbPut(oce.IdxTOK.String(), share2)

		case err := <- tradeSub.Err():
			log.Fatal(err)
		case oce := <-tradeChannel:
			fmt.Printf("Starting to trade...\n")
			cmd := exec.Command("python3", "Scripts/hbswap/python/set_data.py", serverID, oce.IdxETH.String(), oce.IdxTOK.String(), oce.MaskedETH.String(), oce.MaskedTOK.String())
			utils.ExecCmd(cmd)

			cmd = exec.Command(prog, "-N", players, "-T", threshold, "-p", serverID, "-pn", mpcPort, "hbswap_trade")
			utils.ExecCmd(cmd)

			cmd = exec.Command("python3", "Scripts/hbswap/python/org_data.py", serverID)
			utils.ExecCmd(cmd)
		}
	}
}

func main() {
	serverID = os.Args[1]

	conn := utils.GetEthClient("ws://127.0.0.1:8546")

	var wg sync.WaitGroup
	wg.Add(1)
	Watch(conn)
	wg.Wait()
}

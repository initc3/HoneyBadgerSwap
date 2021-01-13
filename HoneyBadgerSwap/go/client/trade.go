package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	user := os.Args[1]
	tokenA, tokenB := common.HexToAddress(os.Args[2]), common.HexToAddress(os.Args[3])
	amtA, amtB := os.Args[4], os.Args[5]

	conn := utils.GetEthClient(utils.HttpEndpoint)

	owner := utils.GetAccount(fmt.Sprintf("account_%s", user))

	idxA, idxB := utils.TradePrep(conn, owner)

	cmd := exec.Command("python3", "Scripts/hbswap/python/client/req_inputmasks.py", strconv.Itoa(int(idxA)), amtA, strconv.Itoa(int(idxB)), amtB)
	stdout := utils.ExecCmd(cmd)
	maskedInputs := strings.Split(stdout[:len(stdout) - 1], " ")

	maskedA := utils.StrToBig(maskedInputs[0])
	maskedB := utils.StrToBig(maskedInputs[1])


	fmt.Printf("maskedInputs: %v\n", maskedInputs)
	utils.Trade(conn, owner, tokenA, tokenB, big.NewInt(idxA), big.NewInt(idxB), maskedA, maskedB)
}





package main

import "github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"

func main() {
	conn := utils.GetEthClient(utils.HttpEndpoint)

	owner := utils.GetAccount("server_0")

	utils.Reset(conn, owner)
}

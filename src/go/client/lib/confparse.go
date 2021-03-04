package lib

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
)

type EthNodeConfig struct {
	Network  string
	Hostname string
	HttpPort int
	WsPort   int
}

type ServerConfig struct {
	Id       int
	Host     string
	HttpPort int
}

type ClientConfig struct {
	N       int
	T       int
	EthNode EthNodeConfig
	Servers []ServerConfig
}

func ParseClientConfig(configfile string) ClientConfig {
	fmt.Println("Parsing config file: ", configfile)
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}
	config := ClientConfig{}
	toml.Unmarshal(data, &config)
	return config
}

it add package lib

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
)

type EthNodeConfig struct {
	Network      string
	Hostname     string
	HttpEndpoint string
	WsEndpoint   string
	HttpPort     int
	WsPort       int
}

type _ServerConfig struct {
	Id       int
	Host     string
	HttpPort int
	HttpHost string
}

type ServerConfig struct {
	N              int
	T              int
	LeaderHostname string
	EthNode        EthNodeConfig
	Servers        []_ServerConfig
}

func ParseServerConfig(configfile string) ServerConfig {
	log.Println("Parsing config file: ", configfile)
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}
	config := ServerConfig{}
	toml.Unmarshal(data, &config)
	return config
}

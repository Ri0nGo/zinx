package config

import (
	"encoding/json"
	"os"
)

type ZinxConfig struct {
	PORT          int    `json:"port"`
	MaxPacketSize uint32 `json:"max_packet_size"`
	MaxConn       int    `json:"max_conn"`
	Name          string `json:"name"`
	IP            string `json:"ip"`
	Version       string `json:"version"`
}

var Conf *ZinxConfig

func (zc *ZinxConfig) Reload() {
	bytes, err := os.ReadFile("../../config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &Conf)
	if err != nil {
		panic(err)
	}
}

func init() {
	Conf = &ZinxConfig{
		Name:          "ZinxServerApp",
		IP:            "127.0.0.1",
		PORT:          8000,
		MaxPacketSize: 512,
		MaxConn:       1000,
		Version:       "1.0",
	}
	Conf.Reload()
}

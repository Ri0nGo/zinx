package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ZinxConfig struct {
	PORT              int    `json:"port"`
	MaxConn           int    `json:"max_conn"`
	WorkPoolSize      uint32 `json:"work_pool_size"`
	MaxPacketSize     uint32 `json:"max_packet_size"`
	MaxWorkTaskNumber uint32 `json:"max_work_task_number"`
	Name              string `json:"name"`
	IP                string `json:"ip"`
	Version           string `json:"version"`
}

var Conf *ZinxConfig

func (zc *ZinxConfig) Reload() {
	bytes, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(os.Getwd())
		panic(err)
	}

	err = json.Unmarshal(bytes, &Conf)
	if err != nil {
		panic(err)
	}
}

func init() {
	Conf = &ZinxConfig{
		Name:              "ZinxServerApp",
		IP:                "127.0.0.1",
		PORT:              8000,
		MaxPacketSize:     512,
		MaxConn:           1000,
		Version:           "1.0",
		MaxWorkTaskNumber: 1024,
		WorkPoolSize:      50,
	}
	Conf.Reload()
}

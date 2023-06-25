package main

import (
	"zinx/znet"
)

func RunServer(svcName string, ip string, port int) {
	server := znet.NewServer(svcName, ip, port)
	server.Serve()
}

func main() {
	ip := "127.0.0.1"
	port := 8888
	svcName := "Zinx-TCP-Server01"
	RunServer(svcName, ip, port)
}

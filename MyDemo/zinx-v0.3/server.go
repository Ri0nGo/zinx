package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// PingRouter
// @Description: 基础BaseRouter，相当于实现了BaseRouter的方法，这样就可以进行绑定Router
type PingRouter struct {
	znet.BaseRouter
}

func (b *PingRouter) PreRouter(request ziface.IRequest) {
	fmt.Println("Callback Pre Handler")

	_, err := request.GetConn().GetConnection().Write([]byte("Pre Router Ping...\n"))
	if err != nil {
		fmt.Println("pre write buf error:", err)
	}
}

func (b *PingRouter) Handler(request ziface.IRequest) {
	fmt.Println("Callback Handler ...")

	_, err := request.GetConn().GetConnection().Write([]byte("Handler Router Ping...\n"))
	if err != nil {
		fmt.Println("handler write buf error:", err)
	}
}

func (b *PingRouter) AfterRouter(request ziface.IRequest) {
	fmt.Println("Callback After Handler ...")

	_, err := request.GetConn().GetConnection().Write([]byte("After Router Ping...\n"))
	if err != nil {
		fmt.Println("After write buf error:", err)
	}
}

func RunServer(svcName string, ip string, port int) {
	server := znet.NewServer(svcName, ip, port)
	server.AddRouter(&PingRouter{})
	server.Serve()
}

func main() {
	ip := "127.0.0.1"
	port := 8888
	svcName := "Zinx-TCP-Server01"
	RunServer(svcName, ip, port)
}

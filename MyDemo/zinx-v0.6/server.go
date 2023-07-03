package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call back PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsg().GetMsgId(), ", data=", string(request.GetMsg().GetData()))

	//回写数据
	err := request.GetConn().SendMsg(request.GetMsg().GetMsgId(), []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type TestRouter struct {
	znet.BaseRouter
}

func (pr *TestRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call back TestRouter handler")
	fmt.Println("recv from client : msgId=", request.GetMsg().GetMsgId(), ", data=", string(request.GetMsg().GetData()))

	//回写数据
	err := request.GetConn().SendMsg(request.GetMsg().GetMsgId(), []byte("test..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//创建一个server句柄
	s := znet.NewServer()

	//配置路由
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &TestRouter{})

	//开启服务
	s.Serve()
}

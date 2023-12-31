package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

/*
模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		//发封包message消息
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMessage(0, []byte("Zinx V0.5 Client Test Message")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		fmt.Println(msgHead.GetDataLen())
		if msgHead.GetDataLen() > 0 {
			dataBuf := make([]byte, msgHead.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, dataBuf)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			msgHead.SetData(dataBuf)
			fmt.Println("==> Recv Msg: ID=", msgHead.GetMsgId(), ", len=", msgHead.GetDataLen(), ", data=", string(msgHead.GetData()))
		}

		time.Sleep(1 * time.Second)
	}
}

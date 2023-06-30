package main

import (
	"fmt"
	"io"
	"net"
	"testing"
	"zinx/znet"
)

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("start net server error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("client conn error: ", err)
			continue
		}

		go func(conn net.Conn) {
			// 创建一个包对象
			dp := znet.NewDataPack()

			for {
				headData := make([]byte, dp.GetHeadLen())

				// 读取消息头
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read conn msg error: ", err)
					break
				}
				// 拆解消息包（消息ID 和 消息数据长度）
				msgPack, err := dp.UnPack(headData)
				if err != nil {
					fmt.Println("unpack error: ", err)
					return
				}

				// 若消息数据长度大于0，则开始读取数据
				if msgPack.GetDataLen() > 0 {
					msgData := make([]byte, msgPack.GetDataLen())
					_, err := io.ReadFull(conn, msgData)
					if err != nil {
						fmt.Println("read msg data error: ", err)
					}
					msgPack.SetData(msgData)
					fmt.Printf("===> Server Receive Msg: msg_id: %d, len: %d, data: %s \n",
						msgPack.GetMsgId(), msgPack.GetDataLen(), msgPack.GetData())
				}
			}
		}(conn)
	}

}

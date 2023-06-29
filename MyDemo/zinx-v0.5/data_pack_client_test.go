package zinx_v0_5

import (
	"fmt"
	"net"
	"testing"
	"zinx/znet"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("conn service error: ", err)
		return
	}

	dp := znet.NewDataPack()

	// 封装消息
	msg1 := znet.Message{}
	msg1.SetMsgId(1)
	msg1.SetDataLen(6)
	msg1.SetData([]byte{'R', 'i', 'o', 'n', 'G', 'o'})
	sendBytesData1, err := dp.Pack(&msg1)
	if err != nil {
		fmt.Println("send bytes data1 error: ", err)
		return
	}

	msg2 := znet.Message{}
	msg2.SetMsgId(2)
	msg2.SetDataLen(6)
	msg2.SetData([]byte{'G', 'o', 'o', 'G', 'l', 'e'})
	sendBytesData2, err := dp.Pack(&msg2)
	if err != nil {
		fmt.Println("send bytes data2 error: ", err)
		return
	}

	// 合并两个数据包
	sendBytesData1 = append(sendBytesData1, sendBytesData2...)

	conn.Write(sendBytesData1)
	select {}
}

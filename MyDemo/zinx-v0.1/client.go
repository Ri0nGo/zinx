package main

import (
	"fmt"
	"net"
	"time"
)

func RunClient(ip string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Printf("conn %s:%d failed", err)
	}

	for {
		_, err := conn.Write([]byte("hello Zinx"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)

	}
}

func main() {
	ip := "127.0.0.1"
	port := 8888
	RunClient(ip, port)

}

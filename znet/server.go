package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Port      int
	IP        string
	IPVersion string
	Name      string
}

func NewServer(name string, ip string, port int) ziface.IServer {
	address := net.ParseIP(ip)
	if address == nil {
		panic("ip is incorrect format")
	}
	if port <= 0 && port > 65535 {
		panic("port is incorrect format")
	}
	s := &Server{
		Port:      port,
		IP:        ip,
		IPVersion: "tcp",
		Name:      name,
	}
	return s
}

func (s *Server) Start() {
	//TODO implement me
	go func() {
		// 1. create tcp socket
		addr, err := net.ResolveIPAddr(s.IPVersion, s.IP)
		if err != nil {
			fmt.Println("resolve tcp ip error: ", err)
			return
		}

		// 2. listen socket
		listen, err := net.Listen(s.IPVersion, s.IP)
		if err != nil {
			fmt.Println("listen tcp ip error: ", err)
			return
		}

		fmt.Println("Start Zinx Server: ", s.Name, " success, Now Listening...")
		// 3. accept client conn
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("accept conn error: ", err)
				continue
			}

			// handle conn buss
			go func() {
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read msg error", err)
						continue
					}

					fmt.Println("receive client msg: ", string(buf[:n]))

					if _, err := conn.Write(buf[:n]); err != nil {
						fmt.Println("write msg error", err)
						continue
					}

				}
			}()
		}

	}()

}

func (s *Server) Stop() {
	//TODO implement me

}

func (s *Server) Serve() {
	//TODO implement me
	s.Start()
}

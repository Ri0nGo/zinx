package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/config"
	"zinx/ziface"
)

type Server struct {
	Port      int
	IP        string
	IPVersion string
	Name      string
	Version   string
	Router    ziface.IRouter
}

func NewServer() ziface.IServer {
	address := net.ParseIP(config.Conf.IP)
	if address == nil {
		panic("ip is incorrect format")
	}
	if config.Conf.PORT <= 0 && config.Conf.PORT > 65535 {
		panic("port is incorrect format")
	}
	s := &Server{
		Port:      config.Conf.PORT,
		IP:        config.Conf.IP,
		IPVersion: "tcp",
		Name:      config.Conf.Name,
		Version:   config.Conf.Version,
		Router:    nil,
	}
	return s
}

func (s *Server) Start() {
	//TODO implement me
	fmt.Printf("Server: %s, Address: %s:%d start success....\nt ", s.Name, s.IP, s.Port)

	go func() {
		// 1. create tcp socket
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp ip error: ", err)
			return
		}

		// 2. listen socket
		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp ip error: ", err)
			return
		}

		fmt.Println("Start Zinx Server: ", s.Name, " success, Now Listening...")

		var cid uint64
		cid = 0

		// 3. accept client conn
		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("accept conn error: ", err)
				continue
			}

			// handle conn buss
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
		}

	}()

}

func (s *Server) Stop() {
	//TODO implement me

}

func (s *Server) Serve() {
	//TODO implement me
	s.Start()

	select {}
}

func (s *Server) AddRouter(request ziface.IRouter) {
	s.Router = request
	fmt.Println("Add Router Success...")
}

func CallbackToClient(conn *net.TCPConn, data []byte, count int) error {
	//回显业务
	fmt.Println("[Conn Handle] CallBackToClient ... ")
	if _, err := conn.Write(data[:count]); err != nil {
		fmt.Println("write back buf err ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

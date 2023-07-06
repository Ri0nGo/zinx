package znet

import (
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
	//Router    ziface.IRouter
	MsgHandler ziface.IMsgHandler
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
		Port:       config.Conf.PORT,
		IP:         config.Conf.IP,
		IPVersion:  "tcp",
		Name:       config.Conf.Name,
		Version:    config.Conf.Version,
		MsgHandler: NewMsgHandler(),
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
			dealConn := NewConnection(conn, cid, s.MsgHandler)
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

func (s *Server) AddRouter(msgId uint32, request ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, request)
	fmt.Println("Add Router Success...")
}

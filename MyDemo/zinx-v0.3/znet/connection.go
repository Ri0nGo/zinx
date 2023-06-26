package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	connID      uint64
	conn        *net.TCPConn
	isClosed    bool
	exitBufChan chan bool
	localAddr   string
	remoteAddr  string
	name        string
	Router      ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint64, router ziface.IRouter) *Connection {
	return &Connection{
		connID:      connID,
		conn:        conn,
		isClosed:    false,
		exitBufChan: make(chan bool, 1),
		localAddr:   conn.LocalAddr().String(),
		remoteAddr:  conn.RemoteAddr().String(),
		Router:      router,
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine Is Running....")
	defer fmt.Println(c.RemoteAddrString(), " conn reader exit! ")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.conn.Read(buf)
		if err != nil {
			fmt.Printf("conn id: %d, read msg error: %v \n", c.connID, err)
			c.exitBufChan <- true
			return
		}
		request := Request{
			conn: c,
			data: buf,
		}
		// 执行用户定义的方法，这里感觉使用Handler名称来替代Router更好
		go func(request ziface.IRequest) {
			c.Router.PreRouter(request)
			c.Router.Handler(request)
			c.Router.AfterRouter(request)

		}(&request)

	}
}

func (c *Connection) Start() {
	fmt.Printf("conn id = %d is starting...\n", c.connID)

	go c.StartReader()

	for {
		select {
		case <-c.exitBufChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	//TODO implement me
	fmt.Printf("conn id: %d is stop \n", c.connID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true
	c.conn.Close()

	// notify buf channel the channel will closed
	c.exitBufChan <- true
	close(c.exitBufChan)
}

func (c *Connection) SendMsg(data []byte) error {
	_, err := c.conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetConnID() uint64 {
	return c.connID
}

func (c *Connection) GetConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetName() string {
	return c.name
}

func (c *Connection) RemoteAddrString() string {
	return c.remoteAddr
}

func (c *Connection) LocalAddrString() string {
	return c.localAddr
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

package znet

import (
	"fmt"
	"io"
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
	MsgHandler  ziface.IMsgHandler
}

func NewConnection(conn *net.TCPConn, connID uint64, handler ziface.IMsgHandler) *Connection {
	return &Connection{
		connID:      connID,
		conn:        conn,
		isClosed:    false,
		exitBufChan: make(chan bool, 1),
		localAddr:   conn.LocalAddr().String(),
		remoteAddr:  conn.RemoteAddr().String(),
		MsgHandler:  handler,
	}
}

// StartReader 读数据Goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine Is Running....")
	defer fmt.Println(c.RemoteAddrString(), " conn reader exit! ")
	defer c.Stop()

	for {
		// 创建消息包对象
		dp := NewDataPack()

		// 获取客户端发送的消息头长度
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.conn, headData)
		if err != nil {
			fmt.Printf("conn id: %d, read msg error: %v \n", c.connID, err)
			c.exitBufChan <- true
			continue
		}

		// 拆包，获取msgid 和 data len
		messagePkg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Printf("conn id: %d, unpack msg data error: %v \n", c.connID, err)
			c.exitBufChan <- true
			continue
		}

		// 读取真实数据
		if messagePkg.GetDataLen() > 0 {
			dataBuf := make([]byte, messagePkg.GetDataLen())
			_, err = io.ReadFull(c.conn, dataBuf)
			if err != nil {
				fmt.Printf("conn id: %d,  read msg data error: %v \n", c.connID, err)
				c.exitBufChan <- true
				continue
			}
			messagePkg.SetData(dataBuf)
		}

		request := Request{
			conn: c,
			data: messagePkg,
		}

		err = c.MsgHandler.DoMsgHandler(&request)
		if err != nil {
			fmt.Println("exec msg handler error: ", err)
		}
	}
}

// Start 开始执行
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

// Stop 停止执行
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

// SendMsg 发送消息
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	dp := NewDataPack()
	msgPackBytes, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Printf("msg id: %d, pack msg data error:", msgId)
	}
	_, err = c.conn.Write(msgPackBytes)
	if err != nil {
		return err
	}
	return nil
}

// GetConnID 获取连接id
func (c *Connection) GetConnID() uint64 {
	return c.connID
}

// GetConnection 获取连接对象
func (c *Connection) GetConnection() *net.TCPConn {
	return c.conn
}

// GetName 获取连接名称
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

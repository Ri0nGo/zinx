package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	SendMsg(msgId uint32, data []byte) error
	GetConnID() uint64
	GetConnection() *net.TCPConn
	GetName() string
	RemoteAddr() net.Addr
	LocalAddr() net.Addr
	RemoteAddrString() string
	LocalAddrString() string
}

type HandFunc func(*net.TCPConn, []byte, int) error

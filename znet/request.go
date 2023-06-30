package znet

import (
	"zinx/ziface"
)

type Request struct {
	conn ziface.IConnection
	data ziface.IMessage
}

func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}

func (r *Request) GetMsg() ziface.IMessage {
	return r.data
}

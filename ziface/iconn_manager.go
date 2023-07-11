package ziface

import "zinx/mydemo/zinx-v0.8/ziface"

type IConnManager interface {
	Add(conn ziface.IConnection)
	Remove(conn ziface.IConnection)
	Len() int
	Get(connId uint64) (ziface.IConnection, error)
	Clear()
}

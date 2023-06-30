package ziface

type IRequest interface {
	GetConn() IConnection
	GetMsg() IMessage
}

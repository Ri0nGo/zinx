package znet

import (
	"zinx/ziface"
)

type BaseRouter struct {
}

func (b *BaseRouter) PreRouter(request ziface.IRequest) {
}

func (b *BaseRouter) Handler(request ziface.IRequest) {
}

func (b *BaseRouter) AfterRouter(request ziface.IRequest) {
}

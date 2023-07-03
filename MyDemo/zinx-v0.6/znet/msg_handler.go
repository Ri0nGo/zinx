package znet

import (
	"fmt"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// AddRouter 添加router对象
func (m *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	// check msg id is exists
	if _, ok := m.Apis[msgId]; ok {
		fmt.Println(msgId, " msg id already exist")
		return
	}
	m.Apis[msgId] = router
	fmt.Println("Add router, msg id = ", msgId)
}

// DoMsgHandler 处理handler
func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) error {
	var (
		router ziface.IRouter
		ok     bool
	)
	if router, ok = m.Apis[request.GetMsg().GetMsgId()]; !ok {
		fmt.Println()
		return fmt.Errorf("msg id: %d not found", request.GetMsg().GetMsgId())
	}
	router.PreRouter(request)
	router.Handler(request)
	router.AfterRouter(request)
	return nil
}

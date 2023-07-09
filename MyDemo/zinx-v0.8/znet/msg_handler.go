package znet

import (
	"fmt"
	"zinx/config"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis           map[uint32]ziface.IRouter
	TaskQueue      []chan ziface.IRequest
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, config.Conf.WorkPoolSize),
		WorkerPoolSize: config.Conf.WorkPoolSize,
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

func (m *MsgHandler) SendMsgToQueue(request ziface.IRequest) {
	workerId := request.GetConn().GetConnID() % uint64(config.Conf.WorkPoolSize)
	fmt.Printf("Conn Id: %d, Msg Id: %d \n", request.GetConn().GetConnID(), request.GetMsg().GetMsgId())
	m.TaskQueue[workerId] <- request
}

// StartWorker 开启一个worker 去处理链接中handle
func (m *MsgHandler) StartWorker(workId int, taskQueue chan ziface.IRequest) {
	fmt.Printf("start work id: %d \n", workId)
	for {
		select {
		case task := <-taskQueue:
			m.DoMsgHandler(task)
		}
	}

}

// StartWorkerPool 启动协程池
func (m *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		// 创建一个任务池队列
		m.TaskQueue[i] = make(chan ziface.IRequest, config.Conf.MaxWorkTaskNumber)
		// 将work id 和任务池队列传递给 worker 监听
		go m.StartWorker(i, m.TaskQueue[i])
	}
}

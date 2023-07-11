package znet

import (
	"fmt"
	"sync"
	ziface2 "zinx/mydemo/zinx-v0.8/ziface"
	"zinx/ziface"
)

type ConnManager struct {
	connections map[uint64]ziface.IConnection
	lock        sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint64]ziface.IConnection),
	}
}

func (cm *ConnManager) Add(conn ziface2.IConnection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	cm.connections[conn.GetConnID()] = conn
}

func (cm *ConnManager) Remove(conn ziface2.IConnection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	delete(cm.connections, conn.GetConnID())
}

func (cm *ConnManager) Len() int {
	cm.lock.RLock()
	defer cm.lock.Unlock()

	return len(cm.connections)
}

func (cm *ConnManager) Get(connId uint64) (ziface2.IConnection, error) {
	cm.lock.RLock()
	defer cm.lock.Unlock()

	if connection, ok := cm.connections[connId]; ok {
		return connection, nil
	}
	return nil, fmt.Errorf("%d not found in connections", connId)
}

func (cm *ConnManager) Clear() {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	for connId, connection := range cm.connections {
		connection.Stop()
		delete(cm.connections, connId)
	}
	fmt.Println("clear all connection")
}

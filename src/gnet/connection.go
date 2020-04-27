package gnet

import (
	"log"
	"net"
	"sync"
)

type ConnectionManager struct {
	connMap sync.Map
}

func CreateManager() ConnectionManager {
	return ConnectionManager{
		connMap: sync.Map{},
	}
}

func (cm *ConnectionManager) Store(addr string, conn net.Conn) {
	cm.connMap.Store(addr, conn)
}

func (cm *ConnectionManager) ApplyInRange(function func(key, value interface{}) bool) {
	cm.connMap.Range(function)
}

func (cm *ConnectionManager) Delete(addr string) {
	cm.connMap.Delete(addr)
}

func (cm *ConnectionManager) Shutdown() {
	cm.connMap.Range(func(key, value interface{}) bool {
		if err := value.(net.Conn).Close(); nil != err {
			log.Println(err)
		}
		cm.connMap.Delete(key)
		return true
	})
}

package guser

import (
	"log"
	"net"
	"sync"
)

type ConnectionManager struct {
	connMap sync.Map
}

var cm ConnectionManager
var cmInit sync.Once

func GetConnectionManager() *ConnectionManager {
	cmInit.Do(func() {
		cm = ConnectionManager{
			connMap: sync.Map{},
		}
	})
	return &cm
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

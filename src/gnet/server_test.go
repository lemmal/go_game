package gnet

import (
	"encoding/json"
	"fmt"
	"go_game/src/gevent"
	"log"
	"net"
	"testing"
)

var conn net.Conn

func TestConnect(t *testing.T) {
	event := initEvent()
	msg, err := json.Marshal(event)
	if nil != err {
		log.Fatal(err)
	}
	protocol := CreateProtocol(int32(4+len(msg)), 1, msg)
	connect("tcp", "127.0.0.1 : 2046")
	call(protocol)
	close()
}

func initEvent() gevent.Event {
	param := make(map[string]interface{})
	param["score"] = 10
	return gevent.CreateEvent(1, 2, 1, param)
}

func connect(network string, address string) {
	c, err := net.Dial(network, address) //服务器的ip地址和端口
	if err != nil {
		fmt.Println(" err = ", err)
		return
	}
	conn = c
}
func call(protocol Protocol) {
	write, err := conn.Write(protocol.ToBytes())
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(write)
}

func close() {
	err := conn.Close()
	if nil != err {
		log.Fatal(err)
	}
}

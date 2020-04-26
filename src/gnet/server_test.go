package gnet

import (
	"fmt"
	"log"
	"net"
	"testing"
)

var conn net.Conn

func TestConnect(t *testing.T) {
	msg := []byte("tell me something")
	protocol := CreateProtocol(int32(4+len(msg)), 1, msg)
	connect("tcp", "127.0.0.1 : 2046")
	call(protocol)
	close()
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

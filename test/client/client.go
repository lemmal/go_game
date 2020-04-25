package client

import (
	"fmt"
	"log"
	"net"
)

var conn net.Conn

func Connect(network string, address string) {
	c, err := net.Dial(network, address) //服务器的ip地址和端口
	if err != nil {
		fmt.Println(" err = ", err)
		return
	}
	conn = c
}
func Call(content string) {
	write, err := conn.Write([]byte(content))
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(write)
}

func Close() {
	err := conn.Close()
	if nil != err {
		log.Fatal(err)
	}
}

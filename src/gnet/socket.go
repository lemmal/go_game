package gnet

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

var connMap sync.Map

func Bind() {
	listen, err := net.Listen("tcp", "127.0.0.1 : 2046")
	if nil != err {
		log.Fatal(err)
	}
	accept(listen)
}

func accept(listen net.Listener) {
	for {
		conn, err := listen.Accept()
		if nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			continue
		}
		connMap.Store(conn.RemoteAddr().String(), conn)
	}
}

func LoopConn() {
	for {
		selectConn()
	}
}

func selectConn() {
	connMap.Range(func(key, conn interface{}) bool {
		var buf = make([]byte, 1024)
		length, err := conn.(net.Conn).Read(buf)
		if nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			connMap.Delete(conn.(net.Conn).RemoteAddr().String())
			return true
		}
		//TODO
		protocol := BuildProtocolFromBytes(buf[0:length])
		fmt.Println(protocol)
		return true
	})
}

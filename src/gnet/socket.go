package gnet

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

var conns = make(map[string]net.Conn)
var mutex sync.Mutex

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
		mutex.Lock()
		conns[conn.RemoteAddr().String()] = conn
		mutex.Unlock()
	}
}

func LoopConn() {
	for {
		selectConn()
	}
}

func selectConn() {
	mutex.Lock()
	defer mutex.Unlock()
	for _, conn := range conns {
		var buf = make([]byte, 1024)
		length, err := conn.Read(buf)
		if nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			delete(conns, conn.RemoteAddr().String())
			continue
		}
		//TODO
		protocol := BuildProtocolFromBytes(buf[0:length])
		fmt.Println(protocol)
	}
}

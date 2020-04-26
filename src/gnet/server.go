package gnet

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Server struct {
	connMap    sync.Map  //管理客户端连接
	network    string    //网络层协议
	host       string    //服务端host
	port       int       //服务端端口
	shutdownCh chan bool //关闭信号通道
}

func CreateServer(network string, host string, port int) IServer {
	return &Server{
		connMap:    sync.Map{},
		network:    network,
		host:       host,
		port:       port,
		shutdownCh: make(chan bool),
	}
}

func (server *Server) Start() {
	var builder strings.Builder
	builder.WriteString(server.host)
	builder.WriteString(" : ")
	builder.WriteString(strconv.FormatInt(int64(server.port), 10))
	listen, err := net.Listen(server.network, builder.String())
	if nil != err {
		log.Fatal(err)
	}
	log.Printf("=====  server start! network:{%s}, host:{%s}, port:{%d}  =====\n", server.network, server.host, server.port)
	go server.accept(listen)
	go server.loopConn()
}

func (server *Server) GetShutdownChan() chan bool {
	return server.shutdownCh
}

func (server *Server) Shutdown() {
	server.connMap.Range(func(key, value interface{}) bool {
		server.connMap.Delete(key)
		return true
	})
	log.Printf("=====  server destroy! host:{%s}, port:{%d}  =====\n", server.host, server.port)
}

func (server *Server) accept(listen net.Listener) {
	for {
		conn, err := listen.Accept()
		if nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			continue
		}
		server.connMap.Store(conn.RemoteAddr().String(), conn)
	}
}

func (server *Server) loopConn() {
	for {
		server.selectConn()
	}
}

func (server *Server) selectConn() {
	server.connMap.Range(func(key, conn interface{}) bool {
		var buf = make([]byte, 1024)
		length, err := conn.(net.Conn).Read(buf)
		if nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			server.connMap.Delete(conn.(net.Conn).RemoteAddr().String())
			return true
		}
		//TODO
		protocol := BuildProtocolFromBytes(buf[0:length])
		fmt.Println(protocol)
		return true
	})
}

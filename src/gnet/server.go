package gnet

import (
	"go_game/src/gevent"
	"go_game/src/gnet/iface"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Server struct {
	connMap    sync.Map           //管理客户端连接
	network    string             //网络层协议
	host       string             //服务端host
	port       int32              //服务端端口
	shutdownCh chan bool          //关闭信号通道
	eventLoops []gevent.EventLoop //worker协程池
}

func CreateServer(network string, host string, port int32) iface.IServer {
	server := &Server{
		connMap:    sync.Map{},
		network:    network,
		host:       host,
		port:       port,
		shutdownCh: make(chan bool),
		eventLoops: make([]gevent.EventLoop, 8),
	}
	for index, _ := range server.eventLoops {
		server.eventLoops[index] = gevent.CreateLoop(100)
	}
	return server
}

func (server *Server) Start() {
	var builder strings.Builder
	builder.WriteString(server.host)
	builder.WriteString(" : ")
	builder.WriteString(strconv.FormatInt(int64(server.port), 10))
	listener, err := net.Listen(server.network, builder.String())
	if nil != err {
		log.Fatal(err)
	}
	log.Printf("=====  server start! network:{%s}, host:{%s}, port:{%d}  =====\n", server.network, server.host, server.port)
	go server.accept(listener)
	go server.loopConn()
	for _, loop := range server.eventLoops {
		l := loop
		go l.Start()
	}
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
		protocol := BuildProtocolFromBytes(buf[0:length])
		//TODO 验证protocol
		event := gevent.CreateEventFromBytes(protocol.msg)
		index := event.GetUserId() % int32(len(server.eventLoops))
		server.eventLoops[index].Push(event)
		return true
	})
}

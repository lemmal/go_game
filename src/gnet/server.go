package gnet

import (
	"go_game/src/gevent"
	"go_game/src/gnet/iface"
	"go_game/src/guser"
	"go_game/src/util"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	network    string             //网络层协议
	host       string             //服务端host
	port       int32              //服务端端口
	shutdownCh chan bool          //关闭信号通道
	eventLoops []gevent.EventLoop //worker协程池
}

func CreateServer(network string, host string, port int32) iface.IServer {
	server := &Server{
		network:    network,
		host:       host,
		port:       port,
		shutdownCh: make(chan bool),
		eventLoops: make([]gevent.EventLoop, 8),
	}
	for index := range server.eventLoops {
		server.eventLoops[index] = gevent.CreateLoop(100)
	}
	return server
}

func (server *Server) Start() {
	var builder strings.Builder
	builder.WriteString(server.host)
	builder.WriteString(":")
	builder.WriteString(strconv.FormatInt(int64(server.port), 10))
	listener, err := net.Listen(server.network, builder.String())
	if nil != err {
		log.Fatal(err)
	}
	log.Printf("=====  server start! network:{%s}, host:{%s}, port:{%d}  =====\n", server.network, server.host, server.port)
	go server.accept(listener)
	for _, loop := range server.eventLoops {
		l := loop
		go l.Start()
	}
}

func (server *Server) GetShutdownChan() chan bool {
	return server.shutdownCh
}

func (server *Server) Shutdown() {
	guser.GetConnectionManager().Shutdown()
	guser.GetUserManager().Shutdown()
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
		guser.GetConnectionManager().Store(conn.RemoteAddr().String(), conn)
		go server.selectConn(conn)
	}
}

func (server *Server) selectConn(conn net.Conn) {
	for {
		var head = make([]byte, 4)
		if _, err := io.ReadFull(conn, head); nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			guser.GetConnectionManager().Delete(conn.RemoteAddr().String())
			if user, exist := guser.GetUserManager().Conn2user[conn]; exist {
				guser.GetUserManager().Lost(user.UserId)
			}
			return
		}
		length := util.Bytes2Int(head)
		buf := make([]byte, length)
		if _, err := io.ReadFull(conn, buf); nil != err {
			if err != io.EOF {
				log.Println(err)
			}
			guser.GetConnectionManager().Delete(conn.RemoteAddr().String())
			if user, exist := guser.GetUserManager().Conn2user[conn]; exist {
				guser.GetUserManager().Lost(user.UserId)
			}
			return
		}
		protocol := BuildProtocolFromBytes(length, buf)
		//TODO User->Connection管理
		event := gevent.CreateEventFromBytes(protocol.msg)
		if !server.validMsg(&protocol, &event) {
			guser.GetConnectionManager().Delete(conn.RemoteAddr().String())
			msgId := guser.GetUserManager().NextMsgId(event.GetUserId())
			guser.GetUserManager().Lost(event.GetUserId())
			log.Printf("msgId not match. current : %d, want : %d", protocol.msgId, msgId)
			return
		}
		index := event.GetUserId() % int32(len(server.eventLoops))
		loginEvent := gevent.LoginEvent{
			Event: &event,
			Conn:  conn,
		}
		server.eventLoops[index].Push(&loginEvent)
		guser.GetUserManager().IncrNextMsgId(event.GetUserId())
	}
}

func (server *Server) validMsg(protocol *Protocol, event *gevent.Event) bool {
	if protocol.msgId <= 0 {
		return true
	}
	return protocol.msgId == guser.GetUserManager().NextMsgId((*event).GetUserId())
}

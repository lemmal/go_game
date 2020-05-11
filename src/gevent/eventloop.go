package gevent

import (
	"fmt"
	"go_game/src/gevent/iface"
	"go_game/src/guser"
	"log"
)

type EventLoop struct {
	events chan iface.IEvent //事件队列
}

func CreateLoop(maxLen int32) EventLoop {
	return EventLoop{
		events: make(chan iface.IEvent, maxLen),
	}
}

func (loop *EventLoop) Push(event iface.IEvent) {
	loop.events <- event
}

func (loop *EventLoop) Start() {
	for event := range loop.events {
		doDispatch(event)
	}
}

func doDispatch(event iface.IEvent) {
	//TODO 非常简陋，有空修改
	switch event.GetId() {
	case SYS_ID:
		switch event.GetCommand() {
		case CMD_LOGIN:
			onLogin(event.(*LoginEvent))
			break
		}
	}
	fmt.Printf("%d, %d, %d, %v\n", event.GetUserId(), event.GetId(), event.GetCommand(), event.GetParam())
}

func onLogin(event *LoginEvent) {
	if user, exist := guser.GetUserManager().Id2User[event.UserId]; exist {
		guser.GetConnectionManager().Delete(user.Conn.RemoteAddr().String())
		guser.GetUserManager().Lost(event.GetUserId())
		log.Printf("duplicate user : %v", event.GetUserId())
	}
	guser.GetUserManager().Login(event.GetUserId(), event.Conn)
}

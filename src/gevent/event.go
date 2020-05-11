package gevent

import (
	"encoding/json"
	"log"
	"net"
)

const (
	SYS_ID int32 = 1 //系统事件
	EXT_ID int32 = 2 //扩展事件

	CMD_LOGIN int32 = 1 //登录
	CMD_LOST  int32 = 2 //登出
)

type Event struct {
	UserId  int32                  `json:"userId"`  //用户id
	Id      int32                  `json:"id"`      //事件id
	Command int32                  `json:"command"` //事件命令
	Param   map[string]interface{} `json:"param"`   //事件参数
}

type LoginEvent struct {
	*Event
	Conn net.Conn
}

func CreateEventFromBytes(buf []byte) Event {
	m := make(map[string]interface{})
	if err := json.Unmarshal(buf, &m); nil != err {
		log.Println(err)
	}
	userId := m["userId"]
	id := m["id"]
	command := m["command"]
	param := m["param"]
	return Event{
		UserId:  int32(int(userId.(float64))),
		Id:      int32(int(id.(float64))),
		Command: int32(int(command.(float64))),
		Param:   param.(map[string]interface{}),
	}
}

func CreateEvent(userId int32, id int32, command int32, param map[string]interface{}) Event {
	return Event{
		UserId:  userId,
		Id:      id,
		Command: command,
		Param:   param,
	}
}

func (e *Event) GetUserId() int32 {
	return e.UserId
}

func (e *Event) GetId() int32 {
	return e.Id
}

func (e *Event) GetCommand() int32 {
	return e.Command
}

func (e *Event) GetParam() interface{} {
	return e.Param
}

package gevent

import (
	"encoding/json"
	"log"
)

type Event struct {
	UserId  int32                  `json:"userId"`  //用户id
	Id      int32                  `json:"id"`      //事件id
	Command int32                  `json:"command"` //事件命令
	Param   map[string]interface{} `json:"param"`   //事件参数
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
	return CreateEvent(int32(int(userId.(float64))), int32(int(id.(float64))), int32(int(command.(float64))), param.(map[string]interface{}))
}

func CreateEvent(userId int32, id int32, command int32, param map[string]interface{}) Event {
	return Event{
		UserId:  userId,
		Id:      id,
		Command: command,
		Param:   param,
	}
}

func (event *Event) GetUserId() int32 {
	return event.UserId
}

func (event *Event) GetId() int32 {
	return event.Id
}

func (event *Event) GetCommand() int32 {
	return event.Command
}

func (event *Event) GetParam() interface{} {
	return event.Param
}

package gevent

type Event struct {
	userId  int
	id      int
	command string
	param   map[string]interface{}
}

func CreateEvent(userId int, id int, command string, param map[string]interface{}) Event {
	return Event{
		userId:  userId,
		id:      id,
		command: command,
		param:   param,
	}
}

func (event *Event) GetUserId() int {
	return event.userId
}

func (event *Event) GetId() int {
	return event.id
}

func (event *Event) GetCommand() string {
	return event.command
}

func (event *Event) GetParam() interface{} {
	return event.param
}

package gevent

type Event struct {
	id      int
	command string
	param   map[string]interface{}
}

func CreateEvent(id int, command string, param map[string]interface{}) Event {
	return Event{
		id:      id,
		command: command,
		param:   param,
	}
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

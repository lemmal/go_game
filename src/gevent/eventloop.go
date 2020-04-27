package gevent

import (
	"fmt"
)

type EventLoop struct {
	events chan Event //事件队列
}

func CreateLoop(maxLen int32) EventLoop {
	return EventLoop{
		events: make(chan Event, maxLen),
	}
}

func (loop *EventLoop) Push(event Event) {
	loop.events <- event
}

func (loop *EventLoop) Start() {
	for event := range loop.events {
		doDispatch(event)
	}
}

func doDispatch(event Event) {
	//TODO
	fmt.Println(event)
}

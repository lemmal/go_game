package gevent

import (
	"fmt"
)

type EventLoop struct {
	events chan Event //事件队列
}

func CreateLoop(maxLen int) EventLoop {
	return EventLoop{
		events: make(chan Event, maxLen),
	}
}

func (loop *EventLoop) Push(event Event) {
	loop.events <- event
	fmt.Println(len(loop.events))
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

package gevent

import (
	"container/list"
	"fmt"
)

type EventLoop struct {
	events list.List //事件队列
	maxLen int       //最大长度
	worker chan bool //任务协程
}

func CreateLoop(maxLen int) EventLoop {
	return EventLoop{
		events: list.List{},
		maxLen: maxLen,
		worker: make(chan bool),
	}
}

func (loop *EventLoop) Push(event Event) {
	loop.events.PushFront(event)
	working := <-loop.worker
	if !working {
		loop.worker <- true
	}
}

func (loop *EventLoop) Start() {
	loop.worker <- false
	for {
		select {
		case working := <-loop.worker:
			for working {
				back := loop.events.Back()
				if nil == back {
					working = false
					loop.worker <- false
					break
				}
				loop.events.Remove(back)
				event := back.Value.(Event)
				doDispatch(event)
			}
		}
	}
}

func doDispatch(event Event) {
	//TODO
	fmt.Println(event)
}

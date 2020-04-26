package iface

type IEventLoop interface {
	Push(event IEvent) //事件进入队列
	Start()            //事件分发处理

}

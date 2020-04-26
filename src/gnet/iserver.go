package gnet

type IServer interface {
	Start()                     //启动服务
	GetShutdownChan() chan bool //标记服务关闭
	Shutdown()                  //关闭服务
}

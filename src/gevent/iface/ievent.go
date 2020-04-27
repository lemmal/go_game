package iface

type IEvent interface {
	GetUserId() int32      //用户id
	GetId() int32          //事件id
	GetCommand() string    //事件命令
	GetParam() interface{} //事件参数
}

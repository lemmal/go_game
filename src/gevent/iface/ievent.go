package iface

type IEvent interface {
	GetId() int
	GetCommand() string
	GetParam() interface{}
}

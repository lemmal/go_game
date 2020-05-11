package guser

import (
	"log"
	"net"
	"sync"
)

type UserManager struct {
	Id2User   map[int32]User    //id-用户管理
	Conn2user map[net.Conn]User //连接-用户管理
	locker    sync.Mutex        //同步锁
}

var um UserManager
var umInit sync.Once

func GetUserManager() *UserManager {
	umInit.Do(func() {
		um = UserManager{
			Id2User:   make(map[int32]User),
			Conn2user: make(map[net.Conn]User),
			locker:    sync.Mutex{},
		}
	})
	return &um
}

func (um *UserManager) NextMsgId(userId int32) int32 {
	if user, ok := um.Id2User[userId]; !ok {
		return -1
	} else {
		return user.msgId + 1
	}
}

func (um *UserManager) IncrNextMsgId(userId int32) {
	user := um.Id2User[userId]
	user.incrNextMsgId()
}

func (um *UserManager) Login(userId int32, conn net.Conn) {
	user := User{
		msgId: 0,
		Conn:  conn,
	}
	um.locker.Lock()
	defer um.locker.Unlock()
	um.Id2User[userId] = user
	um.Conn2user[conn] = user
}

func (um *UserManager) Lost(userId int32) {
	um.locker.Lock()
	defer um.locker.Unlock()
	user := um.Id2User[userId]
	delete(um.Id2User, userId)
	delete(um.Conn2user, user.Conn)
	if err := user.Conn.Close(); nil != err {
		log.Println(err)
	}
}

func (um *UserManager) Shutdown() {
	um.locker.Lock()
	defer um.locker.Unlock()
	for id, user := range um.Id2User {
		delete(um.Id2User, id)
		delete(um.Conn2user, user.Conn)
	}
}

type User struct {
	UserId int32
	msgId  int32
	Conn   net.Conn
}

func (u *User) incrNextMsgId() {
	u.msgId += 1
}

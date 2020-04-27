package gnet

import (
	"go_game/src/util"
)

type Protocol struct {
	len   int32  //msg长度
	msgId int32  //msgId
	msg   []byte //msg内容
}

func CreateProtocol(len, msgId int32, msg []byte) Protocol {
	return Protocol{
		len:   len,
		msgId: msgId,
		msg:   msg,
	}
}

func BuildProtocolFromBytes(buf []byte) Protocol {
	length := util.Bytes2Int(buf[0:4])
	return Protocol{
		len:   length,
		msgId: util.Bytes2Int(buf[4:8]),
		msg:   buf[8 : length+4],
	}
}

func (p *Protocol) ToBytes() (buf []byte) {
	buf = make([]byte, 0)
	buf = append(append(append(buf, util.Int2Bytes(p.len)...), util.Int2Bytes(p.msgId)...), p.msg...)
	return buf
}

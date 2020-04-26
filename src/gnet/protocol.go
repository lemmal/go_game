package gnet

import (
	"bytes"
	"encoding/binary"
	"log"
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
	length := bytes2Int(buf[0:4])
	return Protocol{
		len:   length,
		msgId: bytes2Int(buf[4:8]),
		msg:   buf[8:length],
	}
}

func bytes2Int(buf []byte) int32 {
	var result int32
	nb := bytes.NewBuffer(buf)
	if err := binary.Read(nb, binary.BigEndian, &result); nil != err {
		log.Println(err)
	}
	return result
}

func int2Bytes(i int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, i); nil != err {
		log.Println(err)
	}
	return bytesBuffer.Bytes()
}

func (p *Protocol) ToBytes() (buf []byte) {
	buf = make([]byte, 0)
	buf = append(append(append(buf, int2Bytes(p.len)...), int2Bytes(p.msgId)...), p.msg...)
	return buf
}

package util

import (
	"bytes"
	"encoding/binary"
	"log"
)

func Bytes2Int(buf []byte) int32 {
	var result int32
	nb := bytes.NewBuffer(buf)
	if err := binary.Read(nb, binary.BigEndian, &result); nil != err {
		log.Println(err)
	}
	return result
}

func Int2Bytes(i int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, i); nil != err {
		log.Println(err)
	}
	return bytesBuffer.Bytes()
}

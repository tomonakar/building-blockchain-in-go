package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts int64 to a byte array
func IntToHex(num int64) []byte {
	// bytes.Bufferの初期化 @see https://stackoverflow.com/questions/48335214/how-do-you-initialize-an-empty-bytes-buffer-of-size-n-in-go
	// https://stackoverflow.com/questions/48335214/how-do-you-initialize-an-empty-bytes-buffer-of-size-n-in-go
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

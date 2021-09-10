package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := encoder("Hello, goim!")
	decoder(data)
}

/**
goim 协议结构
4bytes PacketLen 包长度
2bytes HeaderLen 头长度
2bytes Version 协议版本号
4bytes Operation 协议指令
4bytes Sequence 序列号
PacketLen-HeaderLen Body 实际业务数据
 */
func encoder(body string) []byte {
	headerLen := 16
	packageLen := len(body) + headerLen
	ret := make([]byte, packageLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(packageLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))

	version := 102
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	operation := 7
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation))
	sequence := 1
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))

	byteBody := []byte(body)

	copy(ret[16:], byteBody)

	return ret
}

func decoder(data []byte) {
	if len(data) < 16 {
		fmt.Println("data len < 16. err package")
		return
	}

	packageLen := binary.BigEndian.Uint32(data[:4])
	fmt.Println("package len:", packageLen)

	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Println("header len:", headerLen)

	version := binary.BigEndian.Uint16(data[6:8])
	fmt.Println("version:", version)

	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Println("operation:", operation)

	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Println("sequence:", sequence)

	body := string(data[16:])
	fmt.Println("body:", body)
}

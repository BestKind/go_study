package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9500")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	fmt.Println("listen to 9500")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handler(conn)
		}
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		msg, err := unpack(reader)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read pack err: ", err)
			break
		}
		fmt.Println("receive: ", msg)
		conn.Write([]byte(msg))
	}
}

func unpack(reader *bufio.Reader) (string, error) {
	lenByte, err := reader.Peek(2)
	if err != nil {
		return "", err
	}
	length := binary.BigEndian.Uint16(lenByte)
	fmt.Println("length : ", length)

	if uint16(reader.Buffered()) < length {
		return "", errors.New(fmt.Sprintf("Data Err! %v", reader.Buffered()))
	}
	pack := make([]byte, int(length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[2:]), nil
}

package main

import (
	"encoding/binary"
	"net"
)

func main() {
	conn, err:= net.Dial("tcp", "127.0.0.1:9500")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	data, _ := pack("1 => start \n ff")
	_, err = conn.Write(data)
	data, _ = pack("2 => start 11111 \n dd")
	_, err = conn.Write(data)
	data, _ = pack("3 => start 222222222 end \n sss")
	_, err = conn.Write(data)
}

func pack(message string) ([]byte, error) {
	headerLen := 2
	packageLength := len(message) + headerLen
	pkg := make([]byte, packageLength)
	binary.BigEndian.PutUint16(pkg[:2], uint16(packageLength))

	body := []byte(message)
	copy(pkg[2:], body)
	return pkg, nil
}

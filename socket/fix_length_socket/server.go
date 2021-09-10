package main

import (
	"bufio"
	"fmt"
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

	for {
		var data [10]byte
		n, err := bufio.NewReader(conn).Read(data[:])
		if err != nil {
			fmt.Println("receive err: ", err)
			break
		}
		fmt.Println("receive : ", string(data[:n]))
		res := data[:n]
		res = append(res, '\n')
		conn.Write(res)
	}
}

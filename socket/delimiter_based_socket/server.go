package main

import (
	"bufio"
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

	for {
		reader := bufio.NewReader(conn)
		for {
			slice, err := reader.ReadSlice('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Println("read slice err: ", err)
				break
			}
			fmt.Println("receive: ", string(slice))
			conn.Write(slice)
		}
	}
}

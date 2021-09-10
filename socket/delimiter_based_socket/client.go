package main

import (
	"net"
)

func main() {
	conn, err:= net.Dial("tcp", "127.0.0.1:9500")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("1 => start \n ff"))
	conn.Write([]byte("2 => start 11111 \n dd"))
	conn.Write([]byte("3 => start 222222222 end \n sss"))
}

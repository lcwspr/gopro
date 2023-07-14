package main

import (
	"fmt"
	"gateway/001/unpack"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Printf("conn failed err: %v\n", err.Error())
		return
	}
	defer conn.Close()

	unpack.Encode(conn, "Hello World 0!!!")

}

package main

import (
	"fmt"
	"gateway/001/unpack"
	"net"
)

func main() {

	// simple tcp server
	// 1. listen ip + port
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Printf("listen fail: %v", err)
		return
	}

	// 2.accept client req

	fmt.Printf("start server")

	// 3. create goroutine for each req
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		// create goroutine for each connect
		go process(conn)
	}

}

func process(conn net.Conn) {

	// {TODO}: 这里不填写会有什么问题?
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect fail, err is %v", err)
			break
		}

		str := string(bt)
		fmt.Printf("recv from client, data is %+v\n", str)

	}
}

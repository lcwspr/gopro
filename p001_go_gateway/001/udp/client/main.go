package main

import (
	"fmt"
	"net"
)

func main() {

	// 1 connect to server
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})

	if err != nil {
		fmt.Printf("connect error, err is %v\n", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 100; i++ {

		// 2 send data
		_, err = conn.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("send data failed, err: %v\n", err)
			return
		}

		// 3 recv data
		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("receive data failed, err: %v\n", err)
			return
		}

		fmt.Printf("receive from addr: %v data: %v\n", remoteAddr, string(result[:n]))

	}

}

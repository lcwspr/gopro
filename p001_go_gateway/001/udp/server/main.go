package main

import (
	"fmt"
	"net"
)

func main() {

	// step1
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		fmt.Printf("listen failed, err : %v\n", err)
		return
	}

	// 2 循环读取消息内容
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from add: %v, err : %v\n", addr, err)
			break
		}

		go func() {

			// 3 return data
			fmt.Printf("addr: %v data: %v len: %v\n", addr, string(data[:n]), n)
			_, err = listen.WriteToUDP([]byte("recv success"), addr)
			if err != nil {
				fmt.Printf("write failed, err: %v\n", err)
			}
		}()
	}

}

package main

import (
	"fmt"
	"net"
)

// client_server.go
func main() {
	// 客户端主动连接server
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 缓冲读取
	buff := make([]byte, 1024)
	for {
		fmt.Print("请输入发送的内容：")
		fmt.Scan(&buff)
		//fmt.Println("发送的内容:%s\n", string(buff))
		conn.Write(buff)
		// 接收
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := buff[:n]
		fmt.Printf(string(msg))
	}
}

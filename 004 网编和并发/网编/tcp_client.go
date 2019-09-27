package main

import (
	"fmt"
	"net"
)

func main(){
	conn, e := net.Dial("tcp", "127.0.0.1:8888")
	if e != nil{
		fmt.Println("连接失败：", e)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		fmt.Print("请输入：")
		fmt.Scan(&buf)
		conn.Write(buf)

		n, e := conn.Read(buf)
		if e != nil{
			fmt.Println(e)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))

	}

}

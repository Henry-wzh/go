package main

import (
"fmt"
"net"
	"reflect"
	"strings"
)

func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8888")
	if e != nil {
		fmt.Println("创建失败：", e)
	}
	defer listener.Close()
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("连接失败：", e)
			continue
		}
		go ClientConn(conn)
	}

}

func ClientConn(conn net.Conn){
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, reflect.TypeOf(ipAddr))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println(err)
			return
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		res := strings.ToUpper(string(msg))
		conn.Write([]byte(res))
	}



}
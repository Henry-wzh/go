package main

import (
	"fmt"
	"net"
	"strings"
)
// tcp_server.go
func main(){
listener, err := net.Listen("tcp", "127.0.0.1:8888")
if err != nil{
fmt.Println(err)
}
// 关闭资源
defer listener.Close()
for {
// 阻塞等待客户端连接
conn, err := listener.Accept()
if err != nil{
fmt.Println(err)
continue
}
// 创建一个子协程，处理用户请求
go ClientConn(conn)
}
}

// 创建一个子协程，处理用户请求
func ClientConn(conn net.Conn){
// 打印客户端地址
ipAddr := conn.RemoteAddr().String()
fmt.Println(ipAddr, "连接成功")
// 缓冲读取
buffer := make([]byte, 1024)
for {
// n 接收用户数据的长度
n,err := conn.Read(buffer)
if err != nil{
fmt.Println(err)
return
}
// 截取用户有效数据
msg := buffer[:n]
fmt.Printf("接收到%s, %s\n", ipAddr, string(msg))
// 若对方发送exit，退出连接
if "exit" == string(msg){
fmt.Printf("%s退出连接\n", ipAddr)
return
}
// 把接收的数据转成大写，返回客户端
res := strings.ToUpper(string(msg))
conn.Write([]byte(res))
}
}
// HTTP服务端
package main

import (
"fmt"
"net/http"
)

func main() {
	// 第一个参数：路由匹配字符串
	// 第二个参数：响应事件
	http.HandleFunc("/go", myHandler)
	// 第二个参数一般不用，传个空
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// 定义响应函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// GET  POST
	fmt.Println(r.Method)
	// url  /go
	fmt.Println(r.URL.Path)
	fmt.Println(r.Header)
	fmt.Println("body=", r.Body)
	w.Write([]byte("golang"))

}
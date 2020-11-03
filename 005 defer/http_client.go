// HTTP客户端
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)

	// 关闭响应资源
	defer resp.Body.Close()
	// 缓冲区接数据
	buf := make([]byte, 1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		// EOF:end-of-file, 到达文件的底部
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		if n == 0 {
			fmt.Println("读取内容结束！")
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println(string(tmp))
}
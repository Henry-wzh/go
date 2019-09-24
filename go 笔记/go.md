# 环境安装

## 1. 安装

-   下载并安装 goland ide 编辑器
-   下载 golang 文件并安装

## 2. 测试

### 1. go

1.  每个 go 源代码文件开头必须是，package 声明，代表数据哪个包
2.  声明 main，代表可以编译
3.  有且只有一个 main 函数，主函数
4.  导入的包必须使用，不使用会报错
5.  左括号不能单起一行，否则会报错

### 2. goland编辑器

1.  goland 会自动导入使用的包

```go
// test.go
package main
import "fmt"

func main() {
    fmt.Println("hello world!")
}
```

```go
// 命令行
go env
go build test.go
./test
// 或者
go run test.go
```




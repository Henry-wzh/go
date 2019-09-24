# 0. 环境准备

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

-   goland 会自动导入使用的包

```go
// test.go
package main
import "fmt"

func main() {
    fmt.Println("hello world!")
}
```

### 2. go编译器测试

```go
go env
```

### 3. goland编辑器

```go
// 命令行
go env
go build test.go
./test
// 或者
go run test.go
```

## web框架

1.  beego
2.  gin

# 1. 基础语法

## 1. 变量

### 1. go关键字

-   公共25个关键字

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

### 2. 预定义的名字

-   用于内建的常量、类型和函数

#### 1. 内建常量(4)

-   true、false、iota、nil

#### 2. 内建类型(20)

-   int、int8、int16、int32、int64（5）
-   uint、uint8、uint16、uint32、uint64 、uintptr（6）
-   float32、float64、complex64、complex128 （4）
-    bool、byte、rune、string、error（5）

#### 3. 内建函数(13)

-   make、len、cap、new、append、copy、close、delete
-   complex、real、imag
-   panic、recover

### 3. 示例

-  变量只要定义就必须使用，否则会报错

```go
package main
import ("fmt")

func main(){
    // 直接定义
    var a int
    var b int
    // 一次定义多个
    var c, d int
    // 括号里
    var(
    	e int
        f int
    )
     fmt.Println(a, b, c, d, e, f)
}
```

-   变量的初始化

```go
package main
import ("fmt" "reflect")
// 变量
func main(){
    // 方式一
    var v1 int = 10
    // 方式二：编译器自动推到类型
    var v2 = 5
    fmt.Println(reflect.TypeOf(v1))
    fmt.Println(reflect.TypeOf(v2))
    // 方式三： :=  声明并赋值
    v3 := 1
    fmt.Println(reflect.TypeOf(v3))   
    // 方式四
    var v4, v5, v6
    v4, v5, v6 = 1, 2,3
    fmt.Println(v4, v5, v6)
}
```

-   匿名变量：解决变量定义不使用报错的问题

```go
package main
import ("fmt")
// (int, string) 定义函数的返回值，返回 2 个值
func test()(int, string){
    return 666, 'xixi'
}

func main(){
    // 调用函数并接收返回值。定义变量接收
    // _ 表示匿名函数，防止编译不通过
    i, _ := test()
    fmt.Println(i)
}
```

## 2. 常量和枚举

### 1. 常量

#### 1. 直接定义

-   const PI = 3.14

```go
package main
import "fmt"
// 常量
func main(){
    const pi = 3.1415926
    fmt.Println(pi)
}
```

#### 2. 通过函数

```go
package main
import "fmt"

const pi = 3.1415926
func const(){
    const(
    	pi = 3.14
    )
    // 就近原则
    fmt.Println(pi)
}

func main(){
	const()
}
```

#### 3. iota常量生成器

-   常量递增
-   还可一参与计算

```go
// ista = 0 的常量
func enums3(){
    const(
    	python = iota
        java
        golang
    )
    fmt.Printzhln(python, java, golang)
}

func enums4(){
    const(
        // 位运算，b 左移
        b = 1 << (10 * iota)
        kb
        mb
    )
    fmt.Println(b, kb, mb)
}
```

#### 4. 常量补充

-   变量、函数可以先使用后定义

```go
package main
import (
	"fmt"
	"reflect"
)

func main(){
	enums()
	fmt.Println(name)
}

func enums() {
	const(
		a = iota
		b
		c = 9
		d = iota
		e
		f
	)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(reflect.TypeOf(f))
}

var name string = "echo"
```

### 2. 枚举

```go
package main
import "fmt"
// 定义枚举类型
func enums(){
    const(
    	python = 0
        java = 1
        golang = 2
    )
    fmt.Println(python, java, golang)
}
// 定义枚举类型 2
func enums2(){
    const(
    	python = 0
        // 默认和上面的值一样
        java
        golang = 1
        php
    )
    fmt.Println(python, java, golang, php)
}

func main(){
    enums()
    enums2()
    enums3()
}
```

## 3. 基本数据类型

### 1. 整型

-   int32、int64
-   uint32、uint64
-   int：一个字长

```go
package main
import ("fmt", "reflect")
// 整型
func main(){
    // int 和 uint，一个字长
    var v1 in32
    v1 = 123
    v2 := 64
    mt.Println(reflect.TypeOf(v1))
    fmt.Println(reflect.TypeOf(v2))
}
```

### 2.  浮点型

-   float32 和 float64
-   int(a)：强制类型转换s

```go
func main(){
    var v3 float32
    v3 = 12
    // 编译器自动推断成 一个字长
    v4 := 12.5
    mt.Println(reflect.TypeOf(v3))
    fmt.Println(reflect.TypeOf(v4))
   	
    var c float32 = 1
    fmt.Println(unsafe.Sizeof(c))
    
    // 不同类型的不能想加，编译错误
    var a = 1
    fmt.Println(a + v4)
    // 可以类型强制转换
    fmt.Println(a + int(v4))
}
```

### 3. bool

-   true 和 false
-   **注意**：定义的 bool 类型不能赋值给其他类型的值

```go
package main

func main(){
    var v1 bool
    v1 = true
    // 可以使用表达式
    v2 := (1==2)
    fmt.Println(v1, v2)
    fmt.Println(reflect.TypeOf(v2))
    // 取反
    v3 := !v2
    fmt.Println(v3)
    // 逻辑元算符
    if v3 == true && v2 == false{
        fmt.Println("right")
    }
}
```

### 4. byte 和 string

-   byte：字符，单引号
-   string：字符串，双引号

```go
package main

// byte 类型
func main(){
    // 字符和字符串
    var ch byte
    // 字符用单引号，字符串使用双引号
    ch = 'a'
    fmt.Println(ch)
    fmt.Printf("ch = %c\n", ch)
}
```

```go
// Prinf() 输出格式化
%d		// 十进制整型
%b		// 二进制
%c		// 字符输出
%s		// 字符串
%v		// 默认格式
%%		// 输出 %
```

```go
package main
import "fmt"

func main(){
    var str string
    str = 'abc'
    fmt.Printf("%c\n", str[0])
    // uint8 与 byte 一样
    fmt.Printf(reflect.TypeOf(str[0]))
    // 反引号，不进行转义
    str2 := `hello
    abc \n \r ...`
    fmt.Println(str2)
}
```

### 5. 类型别名

-   go语言以包为单位
-   main.自定义类型名

```go
func main(){
    // 定义 myint 类型，实际就是 int
    type myint int
    var i myint = 100
    fmt.Println(i)
    // main.myint，拥有原来类型的所有性质，并且可以自定义方法
    fmt.Println(reflect.TypeOf(i))
}
```

### 6. 类型转换

-   不支持自动类型转换

```go
func main(){
    var ch byte = 'a'
    // 强制转换
    var i int = int(ch)
    fmt.Println(i)
}
```

## 4. fmt 包

-   输入、输出
-   println：默认有换行
-   print 和 printf：没有换行

```go
package main
import "fmt"

func main(){
    a := 15
    // 输出
    fmt.Printf("a=%d\n", a)
    // 输入
    var v int
    fmt.Printf("请输入一个整型：")
    fmt.Scan(&v)
    fmt.Println(v)
}
```



## 5. 流程控制

### 1. 选择

#### 1. if...else

```go
func main(){
    const filename = 'abc.txt'
    // 选择方式一
    content, e = ioutil.ReadFile(filename)
    // 相当于 None
    if e!= nil{
        fmt.Pringln(e)
    }else{
        fmt.Printf("%s\n", content)
    }
    
    // 选择方式二
    if content, e = ioutil.ReadFile(filename): e != nil{
        fmt.Pringln(e)
    }else{
        fmt.Printf("%s\n", content)
    }
}
```

#### 2. switch

-   默认有 break

```go
package main
import "fmt"

func main(){
    res := grade(-1)
    fmt.Println(res)
}

func grade(score int)(string){
    // 定义一个 string 用于返回
    res := ""
	switch {
	case score < 0 || score > 100:
		res = "输入有误"
        // 若想继续执行，添加 fallthrough
	case score < 60:
        res = "不及格"
	case score < 80:
        res = "良好"
	default:
        res = "优秀"
    return res
	}
}
```

### 2. 循环

```go
package main
import "fmt"

func main(){
	sum(100)
	sum2(100)
}

func sum(num int){
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum2(num int){
	i := 0
	sum := 0
	for i <= num{
		sum += i
		i ++
	}
	fmt.Println(sum)
}
```

### 3. 跳转














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
-   Printf()：**输出格式化**

| 格式化 | 功能                                       |
| ------ | ------------------------------------------ |
| **%v** | 按值的本来值输出                           |
| %+v    | 在 %v 基础上，对结构体字段名和值进行展开   |
| %#v    | 输出 Go 语言语法格式的值                   |
| %T     | 输出 Go 语言语法格式的类型和值             |
| **%%** | 输出 % 本体                                |
| **%b** | 整型以二进制方式显示                       |
| %o     | 整型以八进制方式显示                       |
| **%d** | 整型以十进制方式显示                       |
| **%x** | 整型以十六进制方式显示                     |
| %X     | 整型以十六进制、字母大写方式显示           |
| %U     | Unicode 字符                               |
| %f     | 浮点数                                     |
| **%p** | 指针，十六进制方式显示，并加上前导的**0x** |

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
	if content, e := ioutil.ReadFile(filename); e != nil{
		fmt.Println(e)
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

-   for循环

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

-   range

```go
func main(){
    s := 'abc'
    for i, c := range s{
        fmt.printf("%d, %c\n", i, c)
    }
}
```

### 3. 跳转

-   break和continue
-   goto

```go
func main(){
    for i := 0; i<5; i++{
        fmt.println(i)
        goto OUT
    }
    
    OUT:
    fmt.println("yes")
}
```

# 2. 函数

## 1. 自定义函数

-   函数和字段，名字首字母小写代表的是private(**只有包内才能调用**)，大写为public
-   `args ...int`：**表示任意个 int 型数据**

```go
// 也可以这样定义
func example(v1, v2 int){}

// 不定参
func example(args ...int){
    for _, n := range args{
        fmt.Println(n)
    }
}

// 返回多个值
func example()(int, string){}
// 返回一个值
func example() int {}

// 返回值，定义变量名
func example()(a int, str string){
    a = 123
    str = 'abc'
    return
}
```

-   1-100的和

```go
package main

import "fmt"

func main() {
	res := rsum(100)
	fmt.Println(res)

}

// 循环版
func sum(num int) (res int) {
	res = 0
	for i := 1; i <= num; i++ {
		res += i
	}
	return
}

// 递归版
func rsum(num int) (res int) {
	if num == 1 {
		return 1
	}
	return num + rsum(num-1)
}
```

## 2. defer 关键字

### 1. 功能

1.  defer⽤于**延迟一个函数或者方法的执行**
2.  defer语句经常被用于**处理成对的操作**，如打开、关闭、连接、断开连接、加锁、释放锁
3.  通过defer机制，不论函数逻辑多复杂，都能**保证在任何执行路径下，资源被释放**
4.  释放资源的defer应该**直接跟在请求资源的语句后**，以免忘记释放资源

### 2. 单个defer

```go
func main(){
    // 延迟执行，一般用于资源关闭
    defer fmt.Println("defer")
    fmt.Println("main")
}
```

### 2. 多个defer

-   逆序执行，即使有异常也会继续执行

```go
func main(){
    defer fmt.Println("a")
    defer fmt.Println("b")
    defer test(0)
    defer fmt.Println("c")
}
// c b a

func test(num int){
    fmt.Println(100/num)
}
```

# 3. 复合数据类型

-   按照Go语言规范，任何类型在未初始化时都对应一个零值：布尔类型是false，整型是0，字符串是""，而**指针，函数，interface，slice，channel和map的零值都是nil**。

-   总共 **5** 种数据类型

|  类型   |  名称  | 默认值  |
| :-----: | :----: | :-----: |
| pointer |  指针  | **nil** |
|  array  |  数组  |    0    |
|  slice  |  切片  |   nil   |
|   map   |  字典  |   nil   |
| struct  | 结构体 |    -    |

## 1. 指针

1.  &a：代表取a的地址
2.  指针和变量是配合使用的
3.  指针也有类型，只能指向定义类型的数据，如果是对象定义类名

### 1. 定义和初始化

```go
func main(){
    a := 10
    var p *int = &a
    fmt.Printf("%x\n", p)
    fmt.Println(*p)
    // 修改数据
    *p = 20
    fmt.Println(*p, num)
    
	b := "abc"
	var ip *string = &b
	fmt.Printf("%X\n", ip)
	fmt.Println(*ip)
}
```

### 2. 空指针

-   var ptr *int

```go
// 是否为空的判断
if ptr == nil{
    fmt.Println("ptr指针为空")
}
```

### 3. 值传递和引用传递

-   **值传递**：相当于是变量的副本，不会修改原来的变量
-   **引用传递**：传递的是变量的地址，相当于是指针
-   go语言，函数传参都是值传递，**会单独开启一块内存空间**

```c
// c 语言版本
void pass_by_val(int a){
    a++;
}
void pass_by_ref(int *a){
    (*a)++;
}

int main(){
    int a = 3;
    pass_by_val(a);
    printf("%d\n", a);
    pass_by_ref(&a);
    printf("%d\n", a);
}
```

```go
// go 语言版本的交换数据
func swap(a, b *int){
    *a, *b = *b, *a
}

func main(){
    a, b := 3, 4
    swap(&a, &b)
    fmt.Println(a, b)
}
```

### 2. new() 和 make()

#### 1. 区别

1.  Go有两个数据结构创建函数：new和make。
2.  基本的区别是`new(T)`返回一个`*T`，返回的这个指针可以被**隐式地消除引用**。
3.  `make(T, args)`返回一个普通的T。通常情况下，T内部有一些隐式的指针。
4.  一句话，**new返回一个指向已清零内存的指针，而make返回一个复杂的结构。**

#### 2. 几点说明

1.  **new()**用来**分配内存**，但与其他语言中的同名函数不同，它**不会初始化内存**，**只会将内存置零**
2.  **make(t Type, size ... IntergerType)**会返回一个指针，该指针指向新分配的，类型为 t 的零置，适用于创建结构体
4.  **make()的目的不同于new()**，它**只能创建 slice、map、channel**，并返回类型为T(**非指针**)的已初始化(非零值)的值
5.  new是一般是**创建对象**的，也可以创建变量。

```go
func main(){
    // new
    p := new([] int)
    fmt.Println(p)
    // make,   类型， 长度，容量
    v := make([] int, 10, 50)
    fmt.Println(v)
}
```

## 2. 数组

### 1. 声明数组

#### 1. 一维数组

```go
// 数据定义
func main(){
   	// 直接定义
	var arr1 [10] int  
    // :=
    arr2 := [3]int{1,2,3}
    // 省略大小，编译器自动推断
    arr3 := [...]int{1,2,3,4,5}
    fmt.Println(arr1, arr2, arr3)
}
```

#### 2. 二维数组

-   定义

```go
func main(){
    var grid [4][5] int
    fmt.Println(grid)
}
```

### 2. 数组遍历

```go
// 数据定义
func main(){
    arr := [...]int(1,2,3,4,5)
    // 方式一
    for i:=0; i<len(arr3); i++{
        fmt.Println(arr[i])
    }
    // 方式二
    for _, v := range arr{
        fmt.Println(v)
    }
}
```

-   数组传递是传值还是引用？

```go
// 数组的传递是值传递
package main
import "fmt"

func printArr(arr [5]int) {
	arr[0] = 18
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	arr := [...] int{1, 2, 3, 4, 5}
	printArr(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
```

```go
package main

import "fmt"

func printArr(arr *[5]int) {
	arr[0] = 18
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	arr := [...] int{1, 2, 3, 4, 5}
	printArr(&arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
```

### 3. 数组切片

-   切片操作和 python 中的 list 类似，但底层不同，go的切片是 view 操作，类似数据库的 view 操作
-   切片结果是 **slice 类型**：[]类型（如：[]int)，此时是对 **array** 的引用
-   切片索引**不能**超过数组的**最大索引值**

```go
func main(){
    arr := [...]int{1,2,3,4,5,6,7,8,9}
    fmt.Println(arr[2:6])
}
```

## 3. slice

### 0. slice结构

1.  一个slice是一个数组某个部分的引用。
2.  在内存中，它是一个包含3个域的结构体：**指向slice中第一个元素的指针**，slice的**长度**，以及slice的**容量**。
    -   长度是下标操作的上界，如x[i]中i必须小于长度。
    -   容量是分割操作的上界，如x[i:j]中j不能大于容量。
3.  数组的slice并不会实际复制一份数据，它只是创建一个新的数据结构，包含了另外的一个指针，一个长度和一个容量数据。 
4.  如同分割一个字符串，**分割数组也不涉及复制操作**：它只是新建了一个结构来放置一个不同的**指针，长度和容量。**

### 1. 定义

1.  数组的长度固定，go 提供了数组切片(slice)
2.  取值时，索引值超出范围会报错

```go
func main(){
    // 声明切片，只是比数组少了长度
    var s1 [] int
    // :=
    s2 := []int{}
    // make()
    var s3 []int = make([]int, 0, 0)  // 与下一行等价
    var s3 []int = make([]int, 0)
    // 声明并初始化
    s4 := []int{1,2,3,4}
}
```

### 2. 内置函数

#### 1. append()

```go
func main(){
	var s1 []int
    s1 = append(s1, 1)
    fmt.Println(s1)
    s1 = append(s1, 2, 3...)
    fmt.Println(s1)
   	// 用make创建指定大小的切片
    s2 := make([]int, 5)
    s2 = append(s2, 6)
    fmt.Println(s2)
    // 创建并初始化
    s3 := []int{1,2,3}
    s3 = append(s3, 4,5)
    fmt.Println(s3)   
}
```

-   切片补充
-   可以在切片的基础上继续切片，只要不超过原数组就行

```go
// 切片是 view 操作
func main(){
    // 数组
    arr := [...]int{0, 1,2,3,4,5,6,7}
    s1 := arr[2:4]
	fmt.Println(s1, s1[0], s1[1], s1[:6])
    fmt.Println(len(s1))
    fmt.Println(cap(s1))			// 容量到原数组结尾
}
```

#### 2. copy()

-   两个切片之间复制数据

```go
func main(){
    arr := [...]int{0,1,2,3,4,5,6,7,8,9}
    s1 := arr[8:]  // 8，9
    s2 := arr[:5]	// 0，1，2，3，4
    // 将第二个切片元素copy到第一个
    // si从头开始覆盖，索引对应位置赋值，不同位置不变
    copy(s2, s1)
}
```

## 4. Map

#### 0. map内存结构

-   Go中的map在底层是用哈希表实现的，你可以在 $GOROOT/src/pkg/runtime/hashmap.goc 找到它的实现

### 1. 定义

-   **类似python的dict**
-   无序的健值对的集合，可以通过key找到value值

```go
func main(){
    // 定义map,直接声明
    var m1 map[int] string
    var m11 map[int] string = map[int]string{1:"abc", 2:"def"}
    // :=
    m2 := map[int]string{}
    m2 := map[int]string{1:"abc", 2:"def"}
    // make()
    m3 := make(map[int]string)
    // 指定容量的 make()
    m4 := make(map[int]string, 10)
    fmt.Println(m1, m2, m3)
	fmt.Println(m4, len(m4),)
}
```

### 2. 键值操作

```go
func main(){
    m1 := map[int]string{1:"北京", 2:"上海", 3:"广州"}
    // 增加
    m1[4] = "杭州"
    // 修改
    m1[1] = "beijing"
    fmt.Println(m1)
    // 删除
    delete(m1, 2)
    fmt.Println(m1)
}
```

### 3. 遍历

```go
func main(){
    m1 := map[int]string{1:"北京", 2:"上海", 3:"广州"}
    // 遍历
    for k, v := range m1{
        fmt.Println(k, v)
    }
    // 判断某个key对象的value是否存在
    val, ok := m1[1]
    fmt.Println(val, ok)
}
```

## 5. 结构体

### 1.  定义和初始化

-   go语言没有class，只有结构题 struct

```go
// 定义 type 结构体名 struct{}
// 学生的结构体
type Student struct{
    id int
    name string
    gender byte
    age int
    address string
}

func main(){
    // 顺序初始化
    var s1 Student = Student{1, "john", 'f', 20, "beijing"}
    // 指定初始化
    s2 := Student{id:2, age:20}
    // 结构体作为指针变量初始化
    var s3 *Student = &Student{3, "tom", 'm', 20, 'bj'}
    fmt.Println(s1, s2, s3)
    fmt.Println(s1.id, s1.name)
    // 获取指针对象的字段
    fmt.Println((*s3).name)
    fmt.Println(s3.name)      // 底层会转为 (*s3).name
    
}
```

### 2. 结构体参数

-   结构体可以作为函数参数传递

```go
// 定义 type 结构体名 struct{}
// 学生的结构体
type Student struct{
    id int
    name string
    gender string
    age int
    address string
}
// 定义传递学生对象的方法
func temStudent(tmp Student){
    tmp.id = 250
    fmt.Println("temStudent", tmp)
}

// 传递指针对象
func pStudent(p *Student){
    p.id = 300
    fmt.Println("pStudent", p)
}

func main(){
    // 顺序初始化
    var s1 Student = Student{1, "john", "female",  20, "beijing"}
    // 传递非指针
    temStudent(s1)
    fmt.Println("main", s1)
    // 传递非指针
    pStudent(&s1)
    fmt.Println("main", s1)
}
```

# 4. 面向"对象"

## 1. 实现类的特性

1.  没有封装、集成和多态的概念，同样通过结构体实现这些特性

```go
type Person struct{
    name string
    age int
    gender string
}

type Student struct{
    Person					// 匿名字段
    id int
    addr string
    name string
}

func main(){
    s1 := {Person{"echo", "21", "female"}, id :1, name:"henry"}
    var s Student
    s.name = "henry"
    fmt.Println(s)
    s.Person.name = "dean"
    fmt.Println(s)
}
```

-   同名字段情况：赋值问题
-   所有的内置类型和自定义类型都可以作为匿名字段去使用
-   **指针类型匿名字段**

```go
type Person struct{
    name string
    age int
    gender string
}

type Student struct{
	*Person					// 匿名字段
    id int
}

func main(){
    s1 := {&Person{"echo", "21", "female"}, id :1, name:"henry"}
    var s Student
    s.name = "henry"
    fmt.Println(s)
    s.Person.name = "dean"
    fmt.Println(s)
}
```

## 2. 方法

### 1. 定义

-   `func (接收参数 接收类型) 方法名 (参数列表) (返回值)`

```go
type Myint int
// 面向过程
func Add(a, b Myint) Myint{
    return a + b
}
// 面向对象
func (a Myint) Add(b, Myint) Myint{
    return a + b
}  

func main(){
    var a Myint = 1
    var b Myint = 1
    fmt.Println(Add(a, b))
    fmt.Println(a.Add(b))   
}
```

```go
type Person struct{
    name string
    gender string
    age int
}

func (p Person)PrintInfo(){
    fmt.Println(p.name, p.gender, p.age)
}

func main(){
    p := Person{"echo", "female", 18}
    p.PrintInfo()
}
```

### 2. 值语义和引用语义

1.  值语义：不会改变原有数据的值
2.  引用语义：可以修改原来数据

```go
type Person struct{
    name string
    gender string
    age int
}

// 引用语义
func (p *Person)SetInfoPointer(){
    (*p).name = "dean"
    p.gender = 22
}

// 值语义
func (p Person)SetInfoValue(){
    p.name = "eliane"
    p.gender = "female"
    p.age = 23
}

func main(){
    p := Person{"test", "male", 10}
    (&p).SetInfoPointer()
    fmt.Println(p)
    // 值语义，并不会改变原来的值
    p2 := Person{"exam", "female", 100}
    p2.SetInfoValue()
    fmt.Println(p)
}
```

### 3. 方法的继承

```go
type Person struct{
    name string
    gender string
    age int
}

func (p *Person)SetInfoPointer(){
    (*p).name = "dean"
    p.gender = 22
}

type Student struct{
    Person
}

func main(){
    s := Student{Person{"henry", "male", 18}}
    s.SetInfoPointer()
}
```

### 4. 方法的重写

```go
type Person struct{
    name string
    gender string
    age int
}

func (p Person)PrintInfo(){
    fmt.Println("Person", p.name, p.gender, p.age)
}

type Student struct{
    Person
    id int
    addr string
}
// 方法重写
func (s Student)PrintInfo(){
    fmt.Println(s.name, s.gender, s.age)
}

func main(){
    p := Person{"henry", "male", 18}
    p.PrintInfo()
    s := Student{Person{"echo", "female", 16}, 2, "bj"}
    s.PrintInfo()
    // 去调用 person 的方法
    s.Person.PrintInfo()
}
```

### 5. 方法值和方法表达式

```go
package main
import "fmt"
type Person struct {
	name string
	age  int
}

func (p *Person) SetInfo() {
	(*p).name = "henry"
	p.age = 19
}

func main() {
	p := Person{"echo", 18}
	v := p.SetInfo								// 方法值
	v()
	fmt.Println(p)
	f := (*Person).SetInfo						// 方法表达式
	f(&p)
	fmt.Println(p)
}
```

-   创建属性getter和setter 方法并调用

```go
type Dog struct{
    name string
    gender int
}

func (d *Dog)SetName(name string){
    d.name = name
}

func (d *Dog)GetName()string{
    return d.name
}

func (d *Dog)bite(){
    fmt.Println("dog", d.name)
}

func test01(){
    d := Dog{"erha", 1}
    d.bite()
}

func main(){
    test01()
}
```

## 3. 包和封装

### 1.  包的调用

-   包名.公共方法

```go

```

### 2. 接口

1.  接口(interface)是一个自定义类型，描述了**一系列方法集合**
2.  接口不能被实例化
3.  接口名习惯以 er 结尾**

#### 1. 接口使用

```go
// 接口
type Humaner interface{
    // 方法
    Speak()
}
// 结合多态
type Student struct{
    name string
    score int
}
type Teacher struct{
    name string
    group string
}
// 自定义类型
type MyStr string

// 实现接口的方法
func (*s Student)Speak(){
    fmt.Printf("student[%s, %d]\n", s.name, s.score)
}
func (*t Teacher)Speak(){
    fmt.Printf("teacher[%s, %d]\n", t.name, t.group)
}
func (str MyStr)Speak(){
    fmt.Printf("MyStr[%s]\n", "working")
}
// 多态的核心，参数为接口类型
func Who(i Humaner){
    i.Speak()
}

func main(){
    s := &Student{"henry", 90}
    t := &Teacher{"dean", "golang"}
    var tmp MyStr = "abc"
    // 基本调用
    s.Speak()
    t.Speak()
    tmp.Speak()
    // 多态
    Who(s)
    Who(t)
    Who(tmp)
}
```

#### 2. 接口继承

-   和类的继承一样

```go
type SuperHumaner interface {
    Humaner
}
func SuperWho(i SuperHumaner){
    i.Speak()
}

func main(){
    s := Student{"henry", 19}
    SuperWho(s)
}
```

#### 3. 空接口

-   代表是任意类型，如`fmt.Printf()`
-   `...interface{}`：表示任意格式的接口

```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

#### 4. 类型查询

-   `comma-ok`：断言
-   `value.(type)`：获取变量类型
-   `val, ok := value.(int)`：判断 val 是否是 int 类型数据

```go
// 空接口
type Element interface{}
type Person struct{
    name string
    age int
}

func main(){
    list := make([]Element, 3)
    list[0] = 1
    list[1] = "hello"
    list[2] = Person{"echo", 18}
    for index, element := list{
        // 类型断言：val, ok = element.(T)
        if , ok = element.(int); ok{
            fmt.Printf("list[%d]是int类型，值是%d\n", index, val)
        }else if val, ok = element.(string); ok{
            fmt.Printf("list[%d]是string类型，值是%s\n", index, val)
        }else{
            fmt.Println("list[%d]是其他类型")
        } 
    }   
}
```

-   **switch测试**

```go
for index, element := list{
    // 类型断言：val, ok = element.(T)
    switch value.(type){
        case int:
        fmt.Printf("%d是int型，值为%d\n", index, value)
        case string:
        fmt.Printf("%d是string型，值为%s\n", index, value)
        default:
        fmt.Printf("%d是其他类型", index)
    }   
```

# 5. 异常处理

## 1. 系统抛

```go
func test(){
    a := [5]int{1,2,3,4,5}
    index := 10
    a[index] = 123
}

func main(){
    test()
}
```

## 2. 自己抛

```go
package main
import "fmt"

func main(){
	fmt.Println(Improve(8))
    fmt.Println(Improve(-8))
}

func GetCircleArea(radius float64) (area float64){
	if radius < 0{
		panic("半径不能 < 0")
	}
	area = 3.14 * radius * radius
	return
}

func Improve(radius float64) (area float64){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	area = GetCircleArea(radius)
	return
}
```

## 3. 处理异常

-   `bytes, e := ioutil.ReadFile()`
-   定义函数时，一般需要处理异常

```go
func getCircleArea(radius float32)(area float32, err error){
    if radius < 0{
        // 创建异常并返回
        err = errors.New("半径不能为负")
        return
    }
    ret = 3.14 * radius * radius
    return
}

func main(){
    ret, err := getCircleArea(-5)
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Println(ret)
    }
}
```

# 6. 字符串和时间日期类型

## 1. 字符串

### 0. 字符串内存

#### 1. 内存模型

1.  字符串在Go语言内存模型中用一个**2字长的数据结构表示**。
    -   它包含一个**指向字符串存储数据的指针和一个长度数据**。
2.  因为string类型是不可变的，对于**多字符串共享同一个存储数据是安全的**。
3.  切分操作`str[i:j]`会得到一个新的2字长结构，一个可能不同的但仍指向同一个字节序列(即上文说的存储数据)的指针和长度数据。
4.  这意味着字符串切分可以在不涉及内存分配或复制操作
    -   这使得字符串切分的效率等同于传递下标。

#### 2. slice扩容

-   在对slice进行append等操作时，可能会造成slice的自动扩容。其扩容时的大小增长规则是
    -   如果新的大小是当前大小2倍以上，则大小增长为新大小
    -   否则循环以下操作：如果当前大小小于1024，按每次2倍增长，否则每次按当前大小1/4增长。直到增长的大小超过或等于新大小。

### 1. 字符串原理

1.  字符串底层就是`bytes`数组
2.  字符串是由 byte 字节组成，字符串船渡就是byte字节长度
3.  字符串中的字符串是不能直接修改的
4.  `rune` 类型用于表示`utf8`的类型

#### 1. 修改字符

```go
// 修改字符串中的内容
var str = "hello"
bs := []byte(str)
bs[0] = "a"
// 切片需要转回string
str = string(bs)
fmt.Println(str)
```

#### 2. rune 类型

-   一个rune字符由 1 个或多个 byte 类型组成

```go
str := "hello"
fmt.Println(len(str))
// 每个中文占 3个字节
str2 := "hello您好"
fmt.Println(len(str2))
// 将string 转为 rune
r := []rune(str)
fmt.Println(len(r))
```

#### 3. 字符串反转

```go
package main
import "fmt"

func main() {
	var str string = "hello您好"
	rs := strReverse(str)
	fmt.Println(rs)
	if str == rs{
		fmt.Println("是回文")
	}else{
		fmt.Println("不是回文")
	}
}

func strReverse(str string) string{
	// 字符串的反转
	r := []rune(str)
	for i := 0; i < len(r)/ 2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}
```

### 2. 字符串操作

-   可以通过Go标准库中的**strings**和**strconv**两个包中的函数进行相应的操作

#### 1. func len(v Type) int

-   获取字符串长度

```go
func main(){
    fmt.Println(len("hello"))
}
```

#### 2. func Contains(s, substr string) bool

-   n 字符串**s中是否包含substr**，返回bool值

```go
func main(){
	fmt.Println(strings.Contains("hello", "h"))
}
```

#### 3. func HasPrefix(s, prefix string) bool

-   判断字符串s是否以prefix为开头
-    HasSuffix是判断结尾

```go
func main(){
	fmt.Println(strings.HasPrefix("hello", "h"))
    fmt.Println(strings.HasSuffix("hello", "h"))
}
```

#### 4. func Join(a []string, sep string) string

-   字符串链接，把slice a通过sep链接起来

```go
func main(){
	sli := []string{"w", "o", "r", "l", "d"}
	fmt.Println(strings.Join(sli, "*"))
}
```

#### 5. func Index(s, sep string) int

-   在**字符串s中**查找sep所在的位置，返回位置值，找不到返回-1
-   LastIndex是从后往前查找

```go
func main(){
	fmt.Println(strings.Index("hello", "o"))
}
```

#### 6. func Repeat(s string, count int) string

-   重复s字符串count次，最后返回重复的字符串

```go
func main(){
	fmt.Println(strings.Repeat("hello", 3))
}
```

#### 7. func Replace(s, old, new string, n int) string

-   在s字符串中，把**old字符串替换为new字符串**，**n表示替换的次数**，小于0表示全部替换

```go
func main(){
	fmt.Println(strings.Replace("hello", "l", "world", 1))
}
```

#### 8. func Split(s, sep string) []string

-   把s字符串按照sep分割，返回slice
-   sep不会保留

```go
func main(){
	fmt.Println(strings.Split("hello", "l"))
}
```

#### 9. func Trim(s string, cutset string) string

-   在s字符串的头部和尾部去除**cutset**指定的字符串

```go
func main(){
	fmt.Println(strings.Trim("hello", "ho"))
}
```

#### 10. func Fields(s string) []string

-   去除s字符串两边的空格符，并且按照空格分割返回slice

```go
func main(){
	fmt.Println(strings.Fields("    hel lo "))
}
```

### 3. 字符串转换

-   **Append系列函数**：将整数等转换为字符串后，添加到现有的字节数组中

```go
func main(){
    str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
}
```

-   **Format系列函数**：把其他类型的转换为字符串

```go
func main(){
	a := strconv.FormatBool(false)
	b := strconv.FormatInt(-1234, 10)
	//Uint无符号
	c := strconv.FormatUint(1234, 10)
	//与FormatInt一样，简写
	d := strconv.Itoa(-2234)
	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(a))
}
```



## 2. 时间和日期

1.  time包下的 Time 类型用来表示时间
2.  time.Now()：获取当前时间
3.  time.Now().Unix()：获取时间戳

#### 1. 简单使用

```go
now := time.Now()
// 分别获取年月日
year := now.Year()
month := now.Month()
day := now.Day()
hour := now.Hour()
minute := now.Minute()
second := now.Second()
fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// 获取时间戳
timestamp := time.Now().Unix()
```

#### 2. 时间戳转时间

```go
// 获取时间戳对象
timeObj = time.Unix(timestamp, 0)
year := timeObj.Year
...
```

#### 3. time.Now().Format()

-   时间格式化

```go
now := time.Now()
fmt.Println(now)
// golang 中指定为语言诞生时间，时间格式化
timeStr := now.Format("2006/01/02 15:04:05")
fmt.Printf("time:%s\n", timeStr)
```

# 7. 处理json

-   使用go语言内置的 **encoding/json 标准库**

## 1.  json序列化

-   使用 json.Marshal()函数，将一组数据进行 json 格式编码

```go
func Marshal(v interface{})[]slice
```

### 1. 结构体转 json

```go
type Person struct{
    Name string
    Hobby string
}

func main(){
    p := Person{"henry", "play"}
    // 调用 api 转
    b, err = json.Marshal(p)
    // 处理err
    if err != nil{
        fmt.Println("json err", err)
    }else{
        fmt.Println(string(b))
    }
    
    // 格式化转换
    if b,err := json.MarshallIndent(p, "", "  ");err!=nil{
        fmt.Println("json err", err)
    }else{
        fmt.Println(string(b))
    }
}
```

### 2. struct tag

-   解决私有字段小写，无法转 **json** 的问题

```go
type Person struct{
    Name string `json:"name"`
    Hobby string
}
```

### 3. map 转 json

```go
// interface{}:表示任意类型
m := make(map[string]interface{})
m["name"] = "henry"
m["age"] = 18
m["niubility"] = true

mjson, err := json.Marshal(map)
if err != nil{
	fmt.Println("json err", err)
}else{
    fmt.Println(string(mjson))
}
```

## 2.  json反序列化

-   解析到结构体或者接口式，**目标位置必须是地址**

-   解析到结构体

```go
type Person struct{
    Name string `json:"name"`
    Hobby string `json:"hobby"`
}

func main(){
    b := []byte(`{"name":"henry", "hobby":"play"}`)
    // 声明结构体
    var p Person
    // 解析到结构体
    err := json.Unmarshal(b, &p)
    if err != nil{
        fmt.Println("json err", err)
    }else{
        fmt.Println(p)
    }
}
```

-   解析到 interface
    -   **默认解析到 map 上**

```go
func main(){
    b := []byte(`{"Name":"henry", "Hobby":"play"}`)
    // 声明结构体
    var i interface{}
    // 解析到接口
    err := json.Unmarshal(b, &i)
    if err != nil{
        fmt.Println("json err", err)
    }else{
        fmt.Println(i))
    }
    // 判断类型
    m := i.(map[string]interface{})
    for k, v := range m{
        switch val := v.(type){
            case string:
            fmt.Println(k, "是string类型", val)
            ...
            default:
             fmt.Println(k, "是其他类型", val)
        }
    }
}
```

# 8. 文件读写

```go
func main(){
    // 新建文件
    file, err := os.Create("./test.txt")
    if err != nil{
        fmt.Println(err)
        return
    }
    defer file.Close()
    // 写入内容
    for i:=0; i < 5; i++{
        file.WriteString("ab\n")
        file.Write([]byte("cd\n"))
    }
    // 读取文件
    
}
```

```go
// 读取文件
```

# 9. 网编并发

-   **并行**：在同一时刻，有多条指令在多个CPU处理器上同时执行
-   **并发**：在同一时刻，只能有一条指令执行，但多个进程指令被快速地轮换执行

## 1. go语言并发

-   go 从语言层面就支持了并发
-   简化了并发程序的编写

## 2. goroutine 

1.  是 **go 并发设计的核心**
2.  goroutine 就是协程，比线程更小，十几个 goroutine 在底层可能就是五六个线程
3.  go 语言内部实现了 goroutine 的内存共享，执行 goroutine 只需极少的栈内存大概 4-5 KB)

## 3. 创建 goroutine

1.  **只需要在语句前添加 go 关键字，就可以创建并发执行单元**
2.  开发人员无需了解任何执行细节，调度器会自动将其安排到合适的系统线程上执行
3.  如果**主协程退出**了，其他任务**也会退出**

```go
// 子协程
func tt(){
    i:=0
    for {
        i++
        fmt.Printf("tt：%d\n", i)
        time.Sleep(1 * time.Second)
    }
}
// 协程
func main(){
    go tt()
    // 协程执行语句
    go fmt.Println("123")
    // 自定义方法，开协程
    go func(){
        go for i := 0; i<10; i++{
            
        }
        ...
    }
    i:=0
    for {
        i++
        fmt.Printf("main：%d\n", i)
        time.Sleep(1 * time.Second)
    }
    
}
```

## 4. runtime包

### 1. runtime.Gosched()

-   出让 CPU 时间片

```go
func main(){
    go func(s string){
        for i:=0; i < 2; i++{
            fmt.Println(s)
        }
    }("world")
    for i:=0; i < 2; i++{
        // 出让时间片
        runtime.Gosched()
		fmt.Println("hello")
     }
}
```

### 2. runtime.Goexit()

-   立即终止当前协程

```go
func main(){
	go func(){
		defer fmt.Println("A.defer")
		func(){
			defer fmt.Println("B.defer")
			// 退出协程
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {}
}
// B.defer A.defer
```

### 3. runtime.GOMAXPROCS()

-   设置执行代码的CPU核心数
-   返回值：上次设置的值，默认是 logic CPUs

```go
func main(){
    n := runtime.GOMAXPROCS(3)
    fmt.Println(n)
    for{
        go fmt.Print(0)
        fmt.Print(1)
    }   
}
```

## 5. channel

1.  goroutine **运行在相同地址空间**，因此访问共享内存必须做好同步，处理好线程安全问题
2.  goroutine 奉行通过**通信来共享内存**，而不是共享内存来通信
3.  channel **是一个引用类型**，用于多个 goroutine 通讯，其**内部实现了同步，确保并发安全**

### 1. 基本使用

#### 1. 使用`make()`创建

-   定义一个channel 时，也需要定义发送到channel的值的类型

```go
make (chan 类型)
make (chan 类型, 容量)
```

-   当 capacity = 0 时

```go
// 发送数据到管道
channel <- value
// 取数据
<-channel
// 取并接收数据
x := <-channel
// 取、接收并检查数据。管道是否关闭
x, ok := <-channel
```

#### 2. 示例

-   **只要管道存在，取值操作就是阻塞事件**

```go
func main(){
    c := make(chan int)
    go func(){
        defer fmt.Println("子协程结束")
        fmt.Println("子协程运行")
        c <- 666
    }()
    num := <-c
    fmt.Println(num)
    fmt.Println("main")
}
```

### 2. 无(有)缓冲的 channel

-   **无缓冲的通道**是指在接收前没有能力保存任何值的通道
-   **有缓冲的通道**是指在接收前可以保存数据的通道

```go
func main(){
    // 无缓冲的 channel，容量 0
    c := make(chan int, 0)
    // 有缓冲的 channel，容量 3
    c := make(chan int, 3)
    fmt.Println(len(c), cap(c))
    // 子协程存数据
    go func(){
        defer fmt.Println("子协程结束")
        for i:=0; i<3; i++{
            c <- i
        	fmt.Prinln("子协程正在运行",i, len(c), cap(c))
        }
    }()
    time.Sleep(time.Second)
    // 主协程从管道取数据
    for i:=0; i<3; i++{
        num := <- c
		fmt.Println(num)
    }
   fmt.Println("主协程结束")
}
```

### 3. close()

-   可以通过内置的 **close()函数关闭channel**

```go
func main(){
    c := make(chan int)
    go func(){
        for i:=0; i<5; i++{
            c <- i
        }
        // 通知其他协程管道已经关闭
        close(c)
    }()
    
    for {
        if data,ok := <-c;ok{
            fmt.Println(data)
        }else{
            break
        }
    }
    fmt.Println("finished")
}
```

### 4. 单方向的 channel

#### 1. 单向管道

-   默认情况下，通道是双向的

```go
// 普通的通道
var ch1 chan in
// 只用于写 float64 数据的管道
var ch2 chan<- float64
// 只用于取 int 类型数据的管道
var ch3 <-chan int
```

-   可以将 channel 隐式转为单向 channel

```go
c := make(cha int, 3)
// 转为单向只写
var send chan<- int = c
// 转为单向只读
var recv <-chan int = c

send <- 1
<- recv
```

#### 2. 生产则和消费者

```go
// producer
func producer (out chan<-int){
    defer close(out)
    for i:= 0; i<5; i++{
        out <- i
    }
}
// consumer
func consumer(in <-chan int){
    for num := range in {
        fmt.Println(num)
    }
}

func main(){
    c := make(chan int)
    go produer(c)
    consumer(c)
    fmt.Println("done")
}
```

### 5. 定时器

-   **基于管道实现**

#### 1. Timer

-   时间到了，只响应一次

```go
func main(){
	// 基本使用
    timer1 = time.NewTimer(2 * time.Second)
    // 当前时间
    t1 := time.Now()
    fmt.Printf("t1:%v\n", t1)
    t2 := <-timer1.C
    fmt.Printf("t1:%v\n", t2)
    
    // timer 响应事件
    timer2 := time.NewTimer(time.Second)
    for {
       	<-timer2.C
        fmt.Println("时间到")
        break
    }
    
    // 通过 timer 实现延时功能
    // 1
    time.Sleep(2*time.Second)
    fmt.Println("2s")
    // 2
    time3 := time.NewTimer(2*time.Second)
    <-timer3.C
    fmt.Println("2s...")
    // 3
    <-time.After(2 * time.Second)
    fmt.Println("2s.......")
    // 4 停止定时器
    timer4 := time.NewTimer(time.Second)
    go func(){
        <-timer4.C
        fmt.Println("定时时间到")
    }()
    stop := timer4.Stop()			// 返回 bool 值，停止成功则为 true
    if stop{
        fmt.Println("timer4已经关闭")
    }
    // 5 重制定时器
    timer5 := time.NewTimer(3 * time.Second)
    timer5.Reset(time.Second)
    fmt.Println(time.Now())
    fmt.Println(<-time5.C)
}
```

#### 2. Ticker

-   时间到了，多次响应，循环响应

```go
func main(){
    // 创建定时器
    ticker := time.NewTicker(time.Second)
    i := 0
    go func(){
        for {
            // 每取一次，执行一回
            fmt.Println(<-ticker.C)
            i++
            fmt.Println(i)
            if i == 5{
                ticker.Stop()
            }
        }
    }()
    for {}
}
```

### 6. select

1.  go 语言提供了 select 关键字，可以监听 channel 上单 数据流滚动
2.  语法与 switch 类似，区别是 select 要求**每个 case 语句里必须是一个 IO 操作**
3.  如果有符合条件的，则任意选择一条使用

#### 1. select 使用

```go
select {
    // 从管道中正确取出数据，执行
    case <-chan1:
    ...
    // 从管道中正确写入数据，执行
    case chan2<-:
    ...
}
```

#### 2. 示例

```go
func main(){
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)
    go func(){
        int_chan<- 1
    }()
    go func(){
        string_chan<- "hello"
    }()
    
    select {
        case value := <-int_chan:
        fmt.Println("int_chan", value)
        case value := <-string_chan:
        fmt.Println("string_chan:", value)
    }
	fmt.Println("done:")
}
```

### 7. 协程同步锁

1.  go 中 channel 实现了同步，确保并发安全，同时也提供了锁的操作方式
2.  go 中 **sync** 包提供了锁相关的支持
3.  Mutex：以加锁方式解决并发安全问题**(map 并发需要加锁)**

```go
type Account struct{
    money int
    // 声明锁
    flag sync.Mutex
}
// 模仿银行检测
func (a *Account) Check(){
    time.Sleep(time.Second)
}
// 设置余额
func (a *Account) SetAccount(n int){
    a.money += n
}
// 查询余额
func (a *Account) GetAccount() int{
   return a.money
}
// 购买
func (a *Account) Buy(n int){
    a.flag.Lock()
    if a.money > n{
        a.Check()
        a.money -= n
    }
    a.flag.UnLock()
}

func main(){
    var account Account
    account.SetAccount(10)
    go account.Buy(6)
    go account.Buy(5)
    time.Sleep(time.Second)
    fmt.Println(account.GetAccount())
    
}
```

#### 2. sync.WaitGroup

-   用来等待一组子协程结束

```go
func main(){
    // 创建一个数据通道
    ch := make(chan int)
    // 记录子协程个数
    count := 2
    go func(){
        fmt.Println("子协程1")
        ch<-1
    }()
    go func(){
        fmt.Println("子协程2")
        ch<-1
    }()
    for range ch{
        count--
        if count == 0{
            close(ch)
        }
    }
}
```

-   **sync.WaitGroup**
    1.  Add()：添加计数
    2.  Done()：计数 -1
    3.  Wait()：等待所有操作完成

```go
func main(){
	// 声明等待组
	var wg sync.WaitGroup()
    wg.Add(2)
    go func(){
        fmt.Println("子协程1")
        wg.Done()
    }()
    go func(){
        fmt.Println("子协程2")
		 wg.Done()
    }()
    wg.Wait()  
}
```

## 6. socket 编程

```go
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
    // conn.RemoteAddr()是：127.0.0.1:59168 类型：*net.TCPAddr，需要转为字符串
    ipAddr ：= conn.RemoteAddr().String()
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
        msg := buf[:n]
        fmt.Printf("接收到%s, %s\n", ipAddr, string(msg))
        // 若对方发送exit，退出连接
        if "exit" == string(msg){
            fmt.Printf("%s退出连接\n", ipAddr)
            retuo
        }
        // 把接收的数据转成大写，返回客户端
        res := strings.ToUpper(string(msg))
        conn.Write([]byte(res))
    }
}
```

```go
// client_server.go
func main(){
    // 客户端主动连接server
    conn, err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil{
        fmt.Println(err)
        return
    }
    defer conn.Close()
   	// 缓冲读取
    buff := make([]byte, 1024)
    for{
        fmt.Println("请输入发送的内容：")
        fmt.Scan(&buff)
        fmt.Println("发送的内容:%s\n", string(buff))
        // 发送数据
        conn.Write(buff)
        // 接收
        n, err := conn.Read(buff)
        if err != nil{
            fmt.Println(err)
            return
    	}
        msg := buff[:n]
        fmt.Printf(string(msg))
    }
}
```

# 10. 

## 1. web工作流程

- Web服务器的工作原理可以简单地归纳为：

  1.  客户机通过TCP/IP协议建立到服务器的TCP连接
2.  客户端向服务器发送HTTP协议请求包，请求服务器里的资源文档
  3.  服务器向客户机发送HTTP协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理“动态内容”，并将处理得到的数据返回给客户端。
4.  客户机与服务器断开。由客户端解释HTML文档，在客户端屏幕上渲染图形结果。

## 2. HTTP协议

1.  超文本传输协议(HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络协议，它详细规定了浏览器和万维网服务器之间互相通信的规则，通过因特网传送万维网文档的数据传送协议
2.  HTTP协议通常承载于TCP协议之上

### 1. HTTP服务端

```go
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
```

### 2. HTTP客户端

```go
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
```

## 3. 案例实践

1.  读取大量数据
2.  对数据进行清洗(清除不想要的数据)
3.  对数据分类(省份，年龄，性别)

### 1. 读入数据

- 需要先下载解决编码问题：C:\Users\Xu jk>go get github.com/axgle/mahonia

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

func main() {
	read2()
}

//读取数据，适合小文件读取
func read() {
	contentBytes, err := ioutil.ReadFile("D:/课上视频/复习go语言day16/数据/kaifang.txt")
	if err != nil {
		fmt.Println(err)
	}
	contentStr := string(contentBytes)
	//逐行打印
	lineStrs := strings.Split(contentStr, "\n\r")

	for _, linestr := range lineStrs {
		fmt.Println(linestr)
		//ConvertEncoding方法处理中文乱码
		nStr := ConvertEncoding(linestr, "GBK")
		fmt.Println(nStr)

	}
}

//大文件读取,用缓冲区读取
func read2() {
	//找到文件
	file, err := os.Open("D:/课上视频/复习go语言day16/数据/kaifang.txt")
	if err != nil {
		fmt.Println(err)
	}
	//关闭资源
	defer file.Close()
	//缓冲读取
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		utfStr := ConverEncoding(string(line), "GBK")
		fmt.Println(utfStr)
	}
}

//处理中文乱码
func ConvertEncoding(src string, encoding string) (dst string) {
	//创建编码处理器
	enc := mahonia.NewDecoder(encoding)
	utfStr := enc.ConvertString(src)
	dst = utfStr
	return
}
```

### 2. 数据清洗

- 把身份不对的筛选掉，需要按照逗号拆分，取index=1的元素。
- 清理好的数据存一个文件，不要数据存另一个文件(数据分离)。

#### 思路

-   根据身份证，按照34个省划分数据
-   根据34个省，创建34个数据管道
-   读取优质文件，写入对应的省管道
-   将34个管道的数据写出到34个省份文件

```go
package main

import (
	"os"
	"fmt"
	"sync"
	"bufio"
	"io"
	"strings"
 )

//声明省份对象
type Province struct{
	//省份id, 身份证号前2位
	Id string
	//省份名
	Name string
	//对应一个省份文件如：北京.txt
	File *os.File
	//本省文件的数据管道
	chanData chan string
}
//声明等待组
var wg sync.WaitGroup

func main(){
	//创建map,用于存放省份
	//key 为省份编号前2位，值为省份对象
	pMap := make(map[string]*Province)
	//声明34个省市
	ps := []string{"北京市11", "天津市12", "河北省13",
	"山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
	"黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
	"福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
	"湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
	"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
	"陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
	"香港特别行政区81", "澳门特别行政区82", "台湾省83",}
	//遍历所有省市，组装数据
	for _,p := range ps{
		//取到name和id
		name := p[:len(p)-2]
		id :=p[len(p)-2:]
		province := Province{Id:id,Name:name}
		//添加进map
		//key为id,值为省市对象
		pMap[id] = &province
		//为每个省份打开一个文件
		file,_:=os.OpenFile(
			"D:/课上视频/复习go语言day16/数据/34省数据"+province.Name+".txt",
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		//为省份对象的File赋值
		province.File = file
		defer file.Close()
		//创建每个省的数据管道
		province.chanData = make(chan string)
		fmt.Println(name, "管道已创建")
	}
	//遍历34个省市的map,开34个协程向对应文件写数据
	for _,Province := range pMap{
		wg.Add(1)
		//想文件写入数据
		go writeFile(province)
	}
	//读取good文件数据，写入对应省份管道
	file,_:=os.Open("D:/课上视频/复习go语言day16/数据/好坏数据/kaifang_good.txt")
	defer file.Close()
	//新建一个缓冲区，把内容先放在缓冲区
	reader := bufio.NewReader(file)
	//逐行判断身份证前两位
	for{
		//按行读取
		lineBytes,_,err:=reader.ReadLine()
		//读取完毕时，关闭数据管道
		if err == io.EOF{
			for _,Province:=range pMap{
				close(province.chanData)
				fmt.Println(province.Name,"管道已关闭")
			}
			break
		}
		//拿到每行数据，转为字符串
		lineStr:=string(lineBytes)
		//按逗号切分
		fieldsSlice := strings.Split(lineStr,",")
		//id为身份证前2位
		id:=fieldsSlice[1][0:2]
		//对号入座，写入相应的管道
		if province,ok :=pMap[id];ok{
			province.chanData <- (lineStr+"\n")
		}else{
			fmt.Println("未知的省")
		}

	}
	//阻塞等待协程结束
	wg.Wait()

}

//向文件写入数据
func writeFile(province *Province){
	//死循环读取管道，管道关闭时循环结束
	for lineStr:=range province.chanData{
		province.File.WriteString(lineStr)
		fmt.Println(province.Name,"写入",lineStr)
	}
	//标记协程结束
	wg.Done()
}




//处理中文乱码
func ConvertEncoding(src string, encoding string) (dst string) {
	//创建编码处理器
	enc := mahonia.NewDecoder(encoding)
	utfStr := enc.ConvertString(src)
	dst = utfStr
	return
}

```

​	

## 4.  微服务

### 1. 什么是微服务 

- 使用一套小服务来开发单个应用的方式，每个服务运行在独立的进程里，一般采用轻量级通讯机制互联，并且它们可以通过自动化方式部署。

### 2. 微服务的特点

- 单一职责，独立的业务单独放在一个项目里，比如订单服务作为一个项目
- 轻量级的通信：http，rpc  非轻量级(**如java的RMI**)
- 隔离性：每个服务项目隔离，不干扰
- 有自己的数据
- 技术多样性

### 3. 微服务诞生背景

- 互联网行业的快速发展，需求变化快，用户数量变化
- 敏捷开发深入人心，用最小代价，做最快的迭代，频繁修改，测试，上线
- **容器技术的成熟**，是微服务的技术基础。

### 4. 互联网架构演进之路

#### 1. 单体架构

1.  所有功能集成在一个项目中
2.  项目整个打包，可以部署到服务器运行
3.  应用与数据库可以分开部署，提高性能
4.  优点：
    -   小项目的首选，开发成本低，架构简单
5.  缺点：
    -   项目复杂之后，很难扩展和维护
    -   扩展成本，有瓶颈
    -   技术栈受限。

#### 2. SOA架构(面向服务的架构)

1.  SOA defines a way to make software components reusable via service interfaces.
2.  将重复性能功能进行抽取，抽取成对应的服务
3.  通过ESB服务总线去访问，esb是SOA架构必不可少的组件
4.  优点：
    -   提高系统可重复性
    -   ESB减少系统接口耦合问题
5.  缺点：
    -   系统与服务界限模糊，不利于开发
    -   ESB服务接口协议不固定，不利于系统维护
    -   抽取力度较大，有一些耦合性。

#### 3. 微服务架构

1.  将服务层一个一个抽取为微服务
2.  遵循单一原则
3.  微服务之间采用一些轻量协议传输数据
4.  优点：
    -   服务拆分粒度非常细，利于开发
    -   提高系统可维护性
    -   比ESB更轻量
    -   适用于互联网更新换代快的情况
5.  缺点
    -   服务过多，服务治理成本高
    -   开发技术要求高

#### 5. 微服务发现2种方式

- 客户端发现

  - 为护肤启动后，将自己IP和端口进行注册，客户端查询注册，的带提供服务的IP和端口，通过负载均衡访问微服务。

  ![1569765382255](C:\Users\Xu jk\AppData\Roaming\Typora\typora-user-images\1569765382255.png)

- 服务端发现

  - 客户端访问时，不去注册中心了，通过服务端发现代理去直接访问

  ![1569765505909](C:\Users\Xu jk\AppData\Roaming\Typora\typora-user-images\1569765505909.png)

#### 6. 微服务如何部署，更新和扩容

- 微服务部署到docker容器
- 设计服务编排K8S,swarm

## 5. RPC

### 1. RPC简介

- 远程过程调用 是一个计算机通信协议

- 该协议允许运行于一台计算机的程序调用另一台计算机的子程序，而程序员无需额外地为这个交互作用编程

- 如果设计的软件采用面向对象编程，那么远程过程调用亦可程为远程调用或远程方法调用

  

### 2. Golang中如何实现RPC

- Golang中实现RPC非常简单，官方提供封装好的库，还有一些第三方的库。
- Golang官方的net/rpc库使用encoding/gob进行编码、解码，支持tcp和http数据传输方式，由于其他语言不支持gob编码方式，所以Golang的RPC只支持Golang开发的服务器与客户端之间的交互
- 官方还提供了net/rpc/jsonrpc库实现RPC方法，jsonrpc采用JSON进行数据编码，解码，因而支持跨语言调度，目前jsonrpc库是基于tcp协议实现的，暂不支持http传输方式
- Golang的RPC必须符合4个条件才可以
  - 结构体字段首字母要大写，要跨域访问，所以大写
  - 函数名必须首字母大写（可以序列号导出）
  - 函数第一个参数是接收参数，第二个参数返回给客户端参数，必须是指针类型
  - 函数必须有一个返回值error
- 示例：Golang实现RPC程序，实现求矩形面积和周长

#### 1. RPC server端

```go
package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

//声明结构体
type Rect struct {
}

//声明参数结构体
type Params struct {
	Width, Height int
}

//定义公共的rpc方法，供外界访问  求面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

//求周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	//1.注册服务
	rect := new(Rect)
	//注册成rpc服务
	rpc.Register(rect)
	//2.把服务绑定到http上
	rpc.HandleHTTP()
	//3.监听服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}

```

#### 2. RPC client

```go
package main

import (
	"fmt"
	"net/rpc"
)

//声明参数结构体
type Params struct {
	Width, Height int
}

//调用服务
func main() {
	//1连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	//2调用远程方法
	//ret是接收结果
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("面积:", ret)
	//周长
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("周长:", ret)
}
/*
面积: 5000
周长: 300
*/
```

**MTBSystem框架**
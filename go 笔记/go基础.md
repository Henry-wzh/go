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

-   去除s字符串的空格符，并且按照空格分割返回slice

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



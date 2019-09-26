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
	// 方法值
	v := p.SetInfo
	v()
	fmt.Println(p)
	// 方法表达式
	f := (*Person).SetInfo
	f(&p)
	fmt.Println(p)

}

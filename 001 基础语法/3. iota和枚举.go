package main

import (
	"fmt"
	"reflect"
)



func main(){
	//enums()
	//fmt.Println(name)
	//enums2()
	unit()
}

func enums() {
	const(
		a = iota
		b
		c = 9
		d = iota
		e
	)
	fmt.Println(a, b, c, d, e)
	fmt.Println(reflect.TypeOf(e))
}

var name string = "echo"




// 枚举
func enums2(){
	const(
		python = 1
		java
		golang = 2
		c
	)
	fmt.Println(python, java, golang, c)
}

// iota 参与计算

func unit(){
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(b, kb, mb, gb)
}
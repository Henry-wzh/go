package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 整型，浮点型
	var a, b = 1, 2
	var c float32 = 1
	fmt.Println(a, b, reflect.TypeOf(a))
	fmt.Println(unsafe.Sizeof(c))

	var d float32 = 1
	fmt.Println(unsafe.Sizeof(d))
	// 强制类型转换
	fmt.Println(a + int(d))

	// bool
	age := (1==2)
	fmt.Println(age)

	if !age {
		v := "age为假"
		fmt.Println(v)
	}

}

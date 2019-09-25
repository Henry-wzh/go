package main

import (
	"fmt"
	"reflect"
)

func main(){
	var m1 map[int]string
	fmt.Println(m1, reflect.TypeOf(m1))
	m2 := map[int]string{1:"beijing", 2:"shanghai"}
	fmt.Println(m2)
	m3 := make(map[int]string, 2)
	fmt.Println(m3,)
	m3[1] = "北京"
	m3[2] = "上海"
	fmt.Println(m3,)
	for i, v := range m3{
		fmt.Println(i, v)
	}

	delete(m3, 1)
	fmt.Println(m3,)


}

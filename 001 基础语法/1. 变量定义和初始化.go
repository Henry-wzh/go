package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a = 10
	fmt.Println(a)
	var b int
	b = a
	fmt.Println(b, reflect.TypeOf(b))

}
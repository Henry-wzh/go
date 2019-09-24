package main

import (
	"fmt"
	"reflect"
)

func main(){
	type myint int
	var a myint = 5
	fmt.Println(a, reflect.TypeOf(a))
}

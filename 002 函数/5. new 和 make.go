package main

import (
	"fmt"
	"reflect"
)

func main(){
	p := new([] int)
	fmt.Println(reflect.TypeOf(p))

	q := make([]int, 3, 5)
	fmt.Println(q, len(q), cap(q))
	q[2] = 1
	fmt.Println(q, reflect.TypeOf(q))

}
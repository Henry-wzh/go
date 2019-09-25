package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s1 [] int
	fmt.Println(s1, reflect.TypeOf(s1))
	s2 := []int{}
	fmt.Println(s2, reflect.TypeOf(s2))
	s3 := []int{1, 2, 3, 4, 5}
	fmt.Println(s3, reflect.TypeOf(s3))


}

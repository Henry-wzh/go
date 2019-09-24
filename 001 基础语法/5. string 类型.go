package main

import (
	"fmt"
	"reflect"
)

func main(){
	// byte 和 string
	name := "henry"
	fmt.Printf("%c\n", name[0])
	fmt.Println( reflect.TypeOf(name[0]))

	hobby := `play play play`
	fmt.Println(hobby)
}


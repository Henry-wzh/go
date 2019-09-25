package main

import (
	"fmt"
)

func main(){
	var name *string
	//fmt.Println(name)
	//fmt.Println(reflect.TypeOf(name))
	if name == nil{
		fmt.Println("name指针为空")
	}

}

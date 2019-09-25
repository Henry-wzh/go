package main

import "fmt"

func main(){
	stu1 := new(Student)
	stu1.id = 1
	stu1.name = "henry"
	fmt.Println(*stu1)

	stu2 := Student{id:2,name:"echo"}
	fmt.Println(stu2)

	var stu3 *Student = &Student{id:3, name:"elaine"}
	fmt.Println(stu3.name)
	fmt.Println(stu3.id)
}

type Student struct{
	id int
	name string
	age int
}
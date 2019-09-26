package main

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct{
	Person
	id int
	score int
	name string
}


func main(){
	p := Person{"henry", 20}
	s := Student{Person:Person{"echo", 20}, name:"henry"}
	fmt.Println(p)
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.Person.name)

}

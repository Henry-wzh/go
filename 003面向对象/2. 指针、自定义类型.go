package main

import "fmt"
type Person struct {
	name string
	age int
}

type Student struct {
	*Person
	id int
}

func main(){
	s := Student{&Person{"henry", 20}, 2}
	fmt.Println(s.name, s.age, s.id)
}
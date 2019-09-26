package main

import "fmt"

type Humaner interface {
	Speak()
}

type Student struct {
	name string
	age  int
}

type Teacher struct {
	name    string
	subject string
}

type mystr string

func (stu Student) Speak() {
	fmt.Println("I'm a student")
}

func (tea Teacher) Speak() {
	fmt.Println("I'm a teacher")
}

func (my mystr) Speak() {
	fmt.Println("I'm a mystr")
}

func Who(i Humaner) {
	i.Speak()
}

type SuperHumaner interface {
	Humaner
}

func SuperWho(i SuperHumaner) {
	i.Speak()
}

func main() {
	stu := Student{"henry", 19}
	tea := Teacher{"dean", "golang"}
	var my mystr
	//mys := mystr("abc")
	//fmt.Println(mys)
	Who(stu)
	Who(tea)
	Who(my)

	fmt.Println("继承")
	s := Student{"henry", 19}
	SuperWho(s)
}

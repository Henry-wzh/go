package main

import "fmt"

type Person struct {
	name string
	age int
}


func (p *Person)SetInfo(){
	(*p).name = "henry"
	p.age = 19
	//fmt.Println(p)
}


type Student struct {
	Person
}

func (s *Student)SetInfo(){
	s.name = "dean"
}


func main(){
	p := Person{"echo", 18}
	(&p).SetInfo()
	fmt.Println(p)


	var stu Student = Student{}
	//(&stu).SetInfo()
	(&stu).SetInfo()
	fmt.Println(stu)




}
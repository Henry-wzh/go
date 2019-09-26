package main

import "fmt"

type Dog struct {
	nickname string
	age int
}

func (d Dog) info(){
	fmt.Println(d.nickname)
}

func main(){
	d := Dog{nickname:"erha"}
	d.info()
}
package main

import (
	"fmt"
	"time"
)

// producer
func producer (out chan<-int){
	defer close(out)
	for i:= 0; i<5; i++{
		time.Sleep(2 * time.Second)
		out <- i
	}
}
// consumer
func consumer(in <-chan int){
	for num := range in {
		fmt.Println(num)
	}
}

func main(){
	c := make(chan int)
	go producer(c)
	consumer(c)
	fmt.Println("done")
}

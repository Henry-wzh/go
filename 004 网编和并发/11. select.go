package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 1
	}()
	go func() {
		chan2 <- "hello"
	}()
	select {
	// 从管道中正确取出数据，执行
	case v := <-chan1:
		fmt.Println("chan1", v)
	// 从管道中正确写入数据，执行
	case s := <- chan2:
		fmt.Println("chan2", s)
	}

}

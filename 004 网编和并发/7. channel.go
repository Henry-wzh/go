package main

import (
	"fmt"
)

func main(){
	//ch := make(chan int, 5)
	ch := make(chan int)
	go func() {
		defer fmt.Println("子协程执行结束")
		defer close(ch)
		for i:=1; i<=5;i++{
			ch<-i
			fmt.Println("子协程正在运行",i, len(ch), cap(ch))
		}
	}()
	//time.Sleep(time.Second)
	for i:=0; i<100;i++{
		fmt.Println("主协程：", <-ch)
	}



}

package main

import (
	"fmt"
)

func main(){
	ch := make(chan int, 2)
	//ch := make(chan int)
	go func() {
		defer fmt.Println("子协程执行结束")
		//defer close(ch)
		for i:=1; i<=6;i++{
			ch<-i
			fmt.Println("子协程正在运行",i, len(ch), cap(ch))
		}
	}()
	//time.Sleep(time.Second)
	for i:=0; i<6;i++{
		fmt.Println("主协程：", <-ch)
	}



}

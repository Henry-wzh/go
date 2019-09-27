package main

import (
	"fmt"
	"runtime"
)

//
//func main(){
//	//n := runtime.GOMAXPROCS(2)
//	//n1 := runtime.GOMAXPROCS(4)
//	//num := runtime.NumCPU()
//	//fmt.Println(n,n1, num)
//
//	for{
//		go fmt.Print(0)
//		fmt.Print(1)
//	}
//}

//func main(){
//	go func(s string){
//		for i:=0; i < 2; i++{
//			fmt.Println(s)
//		}
//	}("world")
//	for i:=0; i < 2; i++{
//		// 出让时间片
//		runtime.Gosched()
//		fmt.Println("hello")
//	}
//}


func main(){
	go func(){
		defer fmt.Println("A.defer")
		func(){
			defer fmt.Println("B.defer")
			// 退出协程
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {}
}

// B.defer A.defer


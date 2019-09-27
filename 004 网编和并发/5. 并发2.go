package main

import (
	"fmt"
	"runtime"
)

//func main(){
//	go func() {
//		for i:=0; i<5;i++{
//			runtime.Gosched()
//			fmt.Println(i)
//		}
//	}()
//
//	go func() {
//		for i:=5; i<10;i++{
//			time.Sleep(time.Microsecond)
//			fmt.Println(i)
//		}
//	}()
//
//	time.Sleep(time.Second)
//}


// runtime.Goexit()
func main(){
	go func() {
		defer fmt.Println("A.defer")
		func(){
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for{}
}
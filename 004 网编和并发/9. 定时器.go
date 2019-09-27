package main

import (
	"fmt"
	"time"
)

func main() {
	//// 基本使用
	//timer1 := time.NewTimer(2 * time.Second)
	//// 当前时间
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	//t2 := <-timer1.C
	//fmt.Printf("t1:%v\n", t2)
	//
	//// timer 响应事件
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//	break
	//}
	//
	//// 通过 timer 实现延时功能
	//// 1
	//time.Sleep(2 * time.Second)
	//fmt.Println("2s")
	//// 2
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2s...")
	//// 3
	//<-time.After(2 * time.Second)
	//fmt.Println("2s.......")
	//// 4 停止定时器
	//timer4 := time.NewTimer(time.Second)
	//go func() {
	//	<-timer4.C
	//	fmt.Println("定时时间到")
	//}()
	//stop := timer4.Stop()
	//fmt.Println(stop)
	//if stop {
	//	fmt.Println("timer4 已经关闭")
	//}

	// 5 重制定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
}

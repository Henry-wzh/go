package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	//fmt.Println(now)
	//fmt.Println(now.Year(), now .Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//timestamp := now.Unix()
	//fmt.Println(timestamp)
	timeObj := time.Unix(now.Unix(), 0)
	fmt.Println(timeObj)

	timeStr := timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)
}

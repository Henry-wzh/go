package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建定时器
	ticker := time.NewTicker(time.Second)
	i := 0
	go func() {
		for {
			// 每取一次，执行一回
			fmt.Println(<-ticker.C)
			i++
			fmt.Println(i)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for {}
}

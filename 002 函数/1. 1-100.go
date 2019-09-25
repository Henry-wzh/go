package main

import "fmt"

func main() {
	res := rsum(100)
	fmt.Println(res)

}

// 循环版
func sum(num int) (res int) {
	res = 0
	for i := 1; i <= num; i++ {
		res += i
	}
	return
}

// 递归版
func rsum(num int) (res int) {
	if num == 1 {
		return 1
	}
	return num + rsum(num-1)
}

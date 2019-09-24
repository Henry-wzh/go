package main

import "fmt"

func test() (int, string) {
	return 666, "xixixi"
}

func consts() (float64){
	const (
		PI = 3.14
	)
	return PI
}

const PI = 3

func main() {
	//i, _ := test()
	//fmt.Println(i, )

	// 常量
	//const (
	//	PI = 3.1415926
	//)

	PI := consts()
	fmt.Println(PI)
}

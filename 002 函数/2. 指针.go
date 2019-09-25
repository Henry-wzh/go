package main

import (
	"fmt"
)

func main() {
	//a := 10
	//var p *int = &a
	//fmt.Printf("%x\n", p)
	//fmt.Println(*p)
	//
	//var b string = "abc"
	//var ip *string = &b
	//fmt.Printf("%X\n", ip)
	//fmt.Println(*ip)

	//test()


	var p, q *int
	a := 10
	p = &a
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", q)

}

func test() {
	x, y := 1, 2
	var arr = [...]int{5: 2}
	//数组指针
	var pf *[6]int = &arr
	fmt.Println(*pf)
	fmt.Println(pf)

	//指针数组
	pfArr := [...]*int{&x, &y}
	fmt.Println(pfArr)
	fmt.Println(*pfArr[0])
	fmt.Println(*pfArr[1])

}

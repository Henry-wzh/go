package main

import "fmt"

func main() {
	//sum(100)
	//sum2(100)
	irange("abc")
}

func sum(num int) {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum2(num int) {
	i := 0
	sum := 0
	for i <= num {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func irange(s string) {
	for i, c := range s {
		fmt.Printf("%d %c \n", i, c)
	}
}

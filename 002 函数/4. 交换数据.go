package main

import "fmt"

// go 语言版本的交换数据
func swap(a, b *int){
	fmt.Printf("a地址：%0x --> a指针：%0x --> 值：%d\n",&a, a, *a)
	fmt.Printf("b地址：%0x --> b指针：%0x --> 值：%d\n",&b, b, *b)
	a, b = b, a
	//*a, *b = *b, *a
	fmt.Printf("a地址：%0x --> a指针：%0x --> 值：%d\n",&a, a, *a)
	fmt.Printf("b地址：%0x --> b指针：%0x --> 值：%d\n",&b, b, *b)
}

func main(){
	a, b := 3, 4
	fmt.Printf("a指针：%0x --> 值：%d\n",&a, a,)
	fmt.Printf("b指针：%0x --> 值：%d\n",&b, b,)
	swap(&a, &b)
	//fmt.Println(a, b)
}

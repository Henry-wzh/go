//package main
//
//import "fmt"
//
//func main() {
//	var str string = "hello您好"
//	//list := []byte(str)
//	//list[0] = 'a'
//	//str = string(list)
//	//fmt.Println(str)
//
//	// 字符串的反转
//	r := []rune(str)
//	for i := 0; i < len(r)/ 2; i++ {
//		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
//	}
//	str = string(r)
//	fmt.Println(str)
//}

package main

import "fmt"

func main() {
	var str string = "hello您好"
	rs := strReverse(str)
	fmt.Println(rs)
	if str == rs{
		fmt.Println("是回文")
	}else{
		fmt.Println("不是回文")
	}
}

func strReverse(str string) string{
	// 字符串的反转
	r := []rune(str)
	for i := 0; i < len(r)/ 2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}


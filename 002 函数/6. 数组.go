package main

import (
	"fmt"
)

func main() {
	//arr := [5] int{1, 2, 3, 4, 5}
	//fmt.Println(arr)
	//arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	//fmt.Println(arr2, len(arr2), cap(arr2), reflect.TypeOf(arr2))
	//// 数组的遍历
	//for i, v := range arr2 {
	//	fmt.Println(i, v)
	//}
	//// 方式二
	//for i := 0; i < len(arr2); i++ {
	//	fmt.Println(arr2[i])
	//}

	//var arr3 [10] int
	//fmt.Println(arr3, reflect.TypeOf(arr3))
	//
	//// 二维数组
	//var grid [2][3] int
	//fmt.Println(grid)

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//a := arr[2:6]
	//fmt.Println(reflect.TypeOf(arr))
	//fmt.Println(a, reflect.TypeOf(a))
	//a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//fmt.Println(arr, a)
	//fmt.Println(&arr, &a)
	s1 := arr[2:4]
	//fmt.Println(s1, s1[0], s1[1], s1[:6])
	//s1 = append(s1,8888,88888)
	//fmt.Println(s1)
	//fmt.Println(arr)
	s2 := arr[6:]
	fmt.Println(s1, s2)
	copy(s1, s2)
	fmt.Println(s1, s2)



}

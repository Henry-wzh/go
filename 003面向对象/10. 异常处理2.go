package main

import (
	"errors"
	"fmt"
)

func getCircleArea(radius float64) (area float64, err error) {
	if radius < 0 {
		err = errors.New("半径不能 < 0")
		return
	}
	area = 3.14 * radius * radius
	return
}

func  main()  {
	if area, err := getCircleArea(-9); err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println(area)
	}
}

//package main
//
//import (
//	"errors"
//	"fmt"
//)
//
//
//func getCircleArea(radius float32) (area float32, err error) {
//	if radius < 0 {
//		// 创建异常并返回
//		err = errors.New("半径不能为负")
//		return
//	}
//	area = 3.14 * radius * radius
//	return
//}
//
//func main() {
//	area, err := getCircleArea(-5)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(area)
//	}
//}

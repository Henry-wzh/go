//package main
//
//import "fmt"
//
//func test(){
//	arr := [5]int{}
//	index := 10
//	arr[index] = 10
//}
//
//
//func main(){
//	//test()
//	//area := GetCircleArea(-1)
//	//fmt.Println(area)
//	fmt.Println(Improve(8))
//
//}
//
//
//func GetCircleArea(radius float64) (area float64){
//	if radius < 0{
//		panic("半径不能 < 0")
//	}
//	area = 3.14 * radius * radius
//	return
//}
//
//func Improve(radius float64) (area float64){
//	defer func(){
//		if err := recover(); err != nil{
//			fmt.Println(err)
//		}
//	}()
//	area = GetCircleArea(radius)
//	return
//
//}
package main
import "fmt"

func main(){
	fmt.Println(Improve(8))
	fmt.Println(Improve(-8))
}

func GetCircleArea(radius float64) (area float64){
	if radius < 0{
		panic("半径不能 < 0")
	}
	area = 3.14 * radius * radius
	return
}

func Improve(radius float64) (area float64){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	area = GetCircleArea(radius)
	return
}
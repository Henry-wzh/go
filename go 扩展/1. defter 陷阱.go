package main

import "fmt"


func main(){
	fmt.Println(f())
}

//func f() (result int) {
//	defer func() {
//		result++
//	}()
//	return 0
//}


func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}



//func f() (r int) {
//	defer func(r int) {
//		r = r + 5
//	}(r)
//	return 1
//}
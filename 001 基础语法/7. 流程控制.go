package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	//max(2, 3)
	read("abc.txt")
}


func max(a int, b int){
	var max int
	if a > b {
		max = a
	}else{
		max = b
	}
	fmt.Println(max)
}


func read(filename string){
	//content, e := ioutil.ReadFile(filename)
	//if e != nil {
	//	fmt.Println(e)
	//}else{
	//	fmt.Printf("%s\n" ,content)
	//}

	if content, e := ioutil.ReadFile(filename); e != nil{
		fmt.Println(e)
	}else{
		fmt.Printf("%s\n", content)
	}

}
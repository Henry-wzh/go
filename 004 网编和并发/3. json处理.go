package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int
}



func main(){
	//p := Person{"henry", 19}
	//b, err := json.Marshal(p)
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(string(b))
	//}
	//
	//var rp Person
	//error := json.Unmarshal(b, &rp)
	//if error != nil{
	//	fmt.Println(error)
	//}else{
	//	fmt.Println(p)
	//}

	mp := make(map[string]interface{}, 3)
	mp["name"] = "henry"
	mp["age"] = 19
	mp["niubility"] = true

	jmp, e := json.Marshal(mp)
	if e != nil{
		fmt.Println(e)
	}else{
		fmt.Println(string(jmp))
	}

	var i interface{}
	e = json.Unmarshal(jmp, &i)
	if e!=nil{
		fmt.Println(e)
	}else{
		fmt.Println(i)
	}

}
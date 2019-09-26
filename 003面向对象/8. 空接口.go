package main

import "fmt"

type Element interface{}

func main() {
	list := make([]Element, 3)
	list[0] = "henry"
	list[1] = 1
	list[2] = 'b'
	for index, value := range list {
		//if val, ok := value.(int); ok {
		//	fmt.Printf("%d是int型，值为%d\n", index, val)
		//} else if val, ok := value.(string); ok {
		//	fmt.Printf("%d是string型，值为%s\n", index, val)
		//} else {
		//	fmt.Printf("%d是其他类型", index)
		//}

		switch value.(type){
			case int:
				fmt.Printf("%d是int型，值为%d\n", index, value)
			case string:
				fmt.Printf("%d是string型，值为%s\n", index, value)
			default:
				fmt.Printf("%d是其他类型", index)
		}

	}
}

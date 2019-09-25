package main
import "fmt"

func printArr(arr *[5]int) {
	fmt.Println(*arr)
	fmt.Println(arr)

	(*arr)[0] = 18
	fmt.Println(*arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	arr := [...] int{1, 2, 3, 4, 5}
	printArr(&arr)
	for _, v := range arr {
		fmt.Println(v)
	}

}

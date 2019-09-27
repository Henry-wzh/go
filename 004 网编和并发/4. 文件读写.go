package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, e := os.Create("text.txt")
	if e != nil {
		fmt.Println(e)
		return
	} else {
		for i := 0; i < 10; i++ {
			file.WriteString("ab\n")
			file.Write([]byte("cd\n"))
		}
	}
	defer file.Close()
	bf := bufio.NewReaderSize(file, 1)
	fmt.Println(*bf)




}

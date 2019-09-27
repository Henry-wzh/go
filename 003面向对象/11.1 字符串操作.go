package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(len("hello"))
	//fmt.Println(strings.Contains("hello", "h"))
	//fmt.Println(strings.HasPrefix("hello", "h"))
	//fmt.Println(strings.HasSuffix("hello", "h"))
	//
	//sli := []string{"w", "o", "r", "l", "d"}
	//fmt.Println(strings.Join(sli, "*"))
	//
	//fmt.Println(strings.Index("hello", "o"))
	//fmt.Println(strings.Repeat("hello", 3))
	//fmt.Println(strings.Replace("hello", "l", "world", 2))
	//
	//fmt.Println(strings.Split("hello", "l"))
	//fmt.Println(strings.SplitAfter("hello", "l"))
	//fmt.Println(strings.SplitAfterN("hello", "l", 7))
	//
	//fmt.Println(strings.Trim("hello", "ho"))
	//
	fmt.Println(strings.Fields("    hel lo "))

	//str := make([]byte, 0, 100)
	//str = strconv.AppendInt(str, 4567, 10)
	//str = strconv.AppendBool(str, false)
	//str = strconv.AppendQuote(str, "abcdefg")
	//str = strconv.AppendQuoteRune(str, '单')
	//fmt.Println(string(str))

	//a := strconv.FormatBool(false)
	//b := strconv.FormatInt(-1234, 10)
	////Uint无符号
	//c := strconv.FormatUint(1234, 10)
	////与FormatInt一样，简写
	//d := strconv.Itoa(-2234)
	//fmt.Println(a, b, c, d)
	//fmt.Println(reflect.TypeOf(a))
}

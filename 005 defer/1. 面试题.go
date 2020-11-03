// 1. 在A和B处填入代码，使输出为foo
package main

type S struct {
	m string
}

func f()*S  {
	return  &S{"foo"} //A
}

func main()  {
	p:= f()   	//B
	print(p.m)
}

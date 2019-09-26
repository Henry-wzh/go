package main
import "fmt"

type Dog struct {
	name   string
	gender int
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) GetName() string{
	return d.name
}

func (d *Dog) bite() {
	fmt.Println("dog", d.name)
}

func test01() {
	d := Dog{"erha", 1}
	d.bite()
	d.SetName("xixihaha")
	res := d.GetName()
	fmt.Println(res)
}

func main() {
	test01()

}

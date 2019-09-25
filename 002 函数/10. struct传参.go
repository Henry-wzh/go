package main
import (
	"fmt"
)

func main(){
	pu := Pupil{1, "henry", 18}
	fmt.Println(pu, "main")
	tmpPupil(pu)
	fmt.Println(pu, "main")
	ptrPupil(&pu)
	fmt.Println(pu, "main")
}

type Pupil  struct {
	id int
	name string
	age int
}

func tmpPupil(pupil Pupil){
	pupil.name = "echo"
	fmt.Println(pupil, "in tmpPupil")
}
func ptrPupil(pupil *Pupil){
	pupil.name = "dean"
	fmt.Println(pupil, "in ptrPupil")
}



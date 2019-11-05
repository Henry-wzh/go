package main
import "fmt"

func main(){
	res := grade(60)
	fmt.Println(res)
}

func grade(score int)(string){
	// 定义一个 string 用于返回
	res := ""
	switch {
	case score < 0 || score > 100:
		res = "输入有误"
		// 若想继续执行，添加 fallthrough
		//fallthrough
	case score < 60:
		res = "不及格"
	case score < 80:
		res = "良好"
	default:
		res = "优秀"
	}
	return res
}
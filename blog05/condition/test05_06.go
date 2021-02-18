package main
import "fmt"

func main() {
	//输入变量
	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	//多分支判断
	if score >= 90 {
		fmt.Println("该同学成绩为A")
	} else if score >=80 && score < 90 {
		fmt.Println("该同学成绩为B")
	} else if score >=70 && score < 80 {
		fmt.Println("该同学成绩为C")
	} else if score >=60 && score < 70 {
		fmt.Println("该同学成绩为D")
	} else {
		fmt.Println("该同学成绩为E")
	}
	fmt.Printf("成绩为:%d", score)
}
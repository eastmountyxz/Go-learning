package main
import "fmt"

func main() {
	//switch后不接表达式 类似于if-else分支来使用
	var age int = 10

	switch {
		case age == 10:
			fmt.Println("年龄为10")
		case age == 20:
			fmt.Println("年龄为20")
		default:
			fmt.Println("没有匹配到")
	}

	//case中也可以对范围进行判断
	var score int = 90
	switch {
		case score > 90:
			fmt.Println("成绩优秀...")
		case score >= 70 && score <= 90:
			fmt.Println("成绩良好...")
		case score >=60 && score <70:
			fmt.Println("成绩及格...")
		default:
			fmt.Println("不及格")
	}
}
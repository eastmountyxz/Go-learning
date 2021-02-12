package main
import "fmt"

func main() {
	//从控制台接收用户信息，包括：姓名、年龄、薪水、是否通过GO课程考试
	var name string
	var age byte
	var money float32
	var isPass bool
	
	//输入
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水:")
	fmt.Scanln(&money)
	fmt.Println("请输入是否通过GO语言课程:")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 考试 %v \n", name, age, money, isPass)
}
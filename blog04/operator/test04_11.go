package main
import "fmt"

func main() {
	//从控制台接收用户信息，包括：姓名、年龄、薪水、是否通过GO课程考试
	var name string
	var age byte
	var money float32
	var isPass bool
	
	//输入空格间隔
	fmt.Println("请输入姓名, 年龄, 薪水, 是否通过考试")
	fmt.Scanf("%s %d %f %t", &name, &age, &money, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 考试 %v \n", name, age, money, isPass)
}
package main
import "fmt"

func main() {
	//输入年龄
	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	//单分支判断
	if age > 18 {
		fmt.Println("你的年龄大于18岁")
	} else {
		fmt.Println("你的年龄小于等于18岁")
	}
}
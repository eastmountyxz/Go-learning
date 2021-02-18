package main
import "fmt"

func main() {
	var sex string
	var second float64

	fmt.Println("请输入秒数:")
	fmt.Scanln(&second)

	//嵌套分支
	if second <=8 {
		fmt.Println("请输入性别:")
		fmt.Scanln(&sex)
		if sex == "男" {
			fmt.Println("进入决赛的男子组")
		} else if sex == "女" {
			fmt.Println("进入决赛的女子组")
		} else {
			fmt.Println("性别输入错误")
		}
	} else {
		fmt.Println("淘汰未进入决赛")
	}
}
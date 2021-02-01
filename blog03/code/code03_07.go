package main
import "fmt"

func main() {
	var (
		name string
		age int
		weight float32
	)

	//从终端获取用户的输入内容
	fmt.Scan(&name, &age, &weight)
	fmt.Println("我的名字是:", name)
	fmt.Println("我的年龄是:", age)
	fmt.Println("我的体重是:", weight)
}
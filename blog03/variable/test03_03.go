package main
import "fmt"

func main() {
	//方法一：指定变量类型 int默认值为0
	var i int
	fmt.Println("i =", i)

	//方法二：根据值自行判定变量类型
	var num = 3.14
	fmt.Println("num =", num)

	//方法三：省略var  :=声明变量并赋值
	name := "eastmount"
	fmt.Println("name =", name)
}
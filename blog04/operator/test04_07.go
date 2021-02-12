package main
import "fmt"

func main() {
	//赋值运算符
	a := 9
	b := 2

	//有两个变量a和b 要求将其进行交换并输出结果
	fmt.Printf("交换前的结果 a = %v , b = %v \n", a, b)

	//定义一个临时变量
	t := a
	a = b
	b = t

	//输出交换后的结果
	fmt.Printf("交换后的结果 a = %v , b = %v \n", a, b)
}
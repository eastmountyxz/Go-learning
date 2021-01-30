package main
import "fmt"

func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func main() {
	sum, sub := getVal(30, 20)
	fmt.Println("sum =", sum, "sub =", sub)
	//只获取第一个返回值
	sum2, _ := getVal(30, 10)
	fmt.Println("sum =", sum2)
}

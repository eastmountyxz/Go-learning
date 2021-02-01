package main
import "fmt"

func main() {
	//转换过程中 按溢出结果处理
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)
}
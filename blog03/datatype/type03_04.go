package main
import "fmt"

func main() {
	//浮点数定义
	var price float32 = 89.12
	fmt.Println("price =", price)

	var num1 float32 = -0.00081
	var num2 float64 = -78942.00912
	fmt.Println("num1 =", num1, "num2 =", num2)
	
	//精度损失
	var num3 float32 = -123.0000203
	var num4 float64 = -123.0000203
	fmt.Println("\n精度损失")
	fmt.Println("num3 =", num3, "num4 =", num4)
}
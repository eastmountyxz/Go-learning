package main
import "fmt"

func main() {
	var a int
	var b int
	
	//输入
	fmt.Println("请输入a:")
	fmt.Scanln(&a)
	fmt.Println("请输入b:")
	fmt.Scanln(&b)

	//计算
	c := a & b
	d := a | b
	e := a ^ b
	fmt.Printf("与运算 a=%v b=%v c=%v \n", a, b, c)
	fmt.Printf("或运算 a=%v b=%v d=%v \n", a, b, d)
	fmt.Printf("异或运算 a=%v b=%v e=%v \n", a, b, e)
}
package main
import "fmt"
import "unsafe"

func main() {
	//浮点数默认声明为 float64类型
	var num = 89.12
	fmt.Printf("num的数据类型是 %T; 占用字节数是 %d \n", num, unsafe.Sizeof(num))

	//十进制数形式
	num1 := 3.14
	num2 := .123
	fmt.Println("num1 =", num1, "num2 =", num2)

	//科学计数法形式
	num3 := 1.234e2      //1.234 * 10的2次方
	num4 := 1.234E2      //1.234 * 10的2次方
	num5 := 1.234E-2     //1.234 / 10的2次方
	fmt.Println("num3 =", num3, "num4 =", num4, "num5 =", num5)
}
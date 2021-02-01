package main
import "fmt"
import "unsafe"

func main() {
	var num1 int = -12
	fmt.Println("变量名称\t数据类型\t占用空间\t对应的值")
	fmt.Printf("num1 \t %T \t %d \t %v \n", num1, unsafe.Sizeof(num1), num1)

	var num2 byte = 127
	fmt.Printf("num2 \t %T \t %d \t %v \n", num2, unsafe.Sizeof(num2), num2)

	var num3 int64 = 12345
	fmt.Printf("num3 \t %T \t %d \t %v \n", num3, unsafe.Sizeof(num3), num3)

	var num4 float32 = 3.14
	fmt.Printf("num4 \t %T \t %d \t %v \n", num4, unsafe.Sizeof(num4), num4)

	var num5 float64 = 3.1415e2
	fmt.Printf("num5 \t %T \t %d \t %v \n", num5, unsafe.Sizeof(num5), num5)

	var num6 int = '杨'
	fmt.Printf("num6 \t %T \t %d \t %v \n", num6, unsafe.Sizeof(num6), num6)

	var num7 bool = false
	fmt.Printf("num7 \t %T \t %d \t %v \n", num7, unsafe.Sizeof(num7), num7)

	var num8 string = "Eastmount"
	fmt.Printf("num8 \t %T \t %d \t %v \n", num8, unsafe.Sizeof(num8), num8)
}

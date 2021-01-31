package main
import "fmt"
import "unsafe"

func main() {
	//查看变量的数据类型
	//fmt.Printf() 用于格式化输出
	var num int = 1024
	fmt.Println("num =", num)
	fmt.Printf("num 的数据类型是 %T \n", num)

	//查看变量占用的空间大小
	var n int64 = 10
	var name = "eastmount"
	fmt.Printf("n 的数据类型是 %T; 占用字节数是 %d \n", n, unsafe.Sizeof(n))
	fmt.Printf("name 的数据类型是 %T; 占用字节数是 %d \n", name, unsafe.Sizeof(name))

	//Golang程序中整型变量使用尽量遵守保小原则
	//在保证程序正确运行下，尽量使用空间小的数据类型
	var age byte = 28
	fmt.Printf("n 的数据类型是 %T; 占用字节数是 %d \n", age, unsafe.Sizeof(age))
}
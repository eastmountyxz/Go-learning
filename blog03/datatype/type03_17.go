package main
import "fmt"

func main() {
	//基本数据类型
	var num int = 10
	fmt.Printf("num的地址=%v 原始值=%v\n", &num, num)

	//指针
	var ptr *int
	ptr = &num
	fmt.Printf("ptr的地址=%v 指向的值为=%v 自身=%v\n", &ptr, *ptr, ptr)

	//修改值
	*ptr = 666
	fmt.Printf("num修改后的值=%v\n", num)
}
package main
import "fmt"
import "unsafe"

func main() {
	//定义数据类型
	var num = false
	fmt.Println("num =", num)

	//注意bool类型占用1个字节 只能取true或false
	fmt.Println("占用空间 =", unsafe.Sizeof(num))
}
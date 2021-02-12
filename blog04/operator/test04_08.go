package main
import "fmt"

func main() {
	//地址
	a := 100
	fmt.Println("a的地址=", &a)
	fmt.Println("a的值=", a)

	//指针变量
	var ptr *int = &a
	fmt.Println("ptr指向的值是=", *ptr)
	fmt.Println("ptr的地址是=", ptr)
}
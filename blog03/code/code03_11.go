package main
import "fmt"

func main() {
	//获取一个int变量num的地址并显示终端
	var num int
	fmt.Scanf("%c", &num)
	fmt.Println("num的地址=", &num)
	fmt.Println("num的值=", num)

	//定义ptr指针变量
	var ptr *int = &num
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)

	//通过ptr去修改num的值
	*ptr = 512
	fmt.Printf("修改值后: num=%d %v\n", num, &num)
	fmt.Printf("修改值后: ptr=%d %v\n", *ptr, &ptr)
}
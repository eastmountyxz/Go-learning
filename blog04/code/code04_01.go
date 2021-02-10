package main
import "fmt"

func main() {
	//不引用第三个变量交换a和b
	var a int = 10
	var b int = 20
	fmt.Printf("原始值：a=%v b=%v \n", a, b)

	a = a + b
	b = a - b   //b = a + b - b ==> b=a
	a = a - b   //a = a + b - a ==> a=b

	fmt.Printf("交换值：a=%v b=%v \n", a, b)
}

package main
import "fmt"

//测试函数
func test() int {
	var num int = 10
	return num
}

func main() {
	var a int = 3
	var b int = 4
	var c int
	var d int

	//赋值运算执行顺序是从右到左
	c = a + 3
	fmt.Println("a=", a, "c=", c)

	//赋值运算左边只能是变量 右边可以使变量、表达式、常量
	//表达式：任何有值都可以看做表达式
	d = a
	d = 8 + 2 * 5
	d = test() + 90
	fmt.Println(d)

	//运算后赋值
	fmt.Println("b =", b)
	b += 2
	fmt.Println(b)
	b -= 3
	fmt.Println(b)
	b *= 2
	fmt.Println(b)
	b /= 3
	fmt.Println(b)
	b %= 2
	fmt.Println(b)
}
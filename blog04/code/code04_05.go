package main
import "fmt"

func main() {
	var a int
	var b int
	var c int
	var t int
	
	//输入
	fmt.Println("请输入a:")
	fmt.Scanln(&a)
	fmt.Println("请输入b:")
	fmt.Scanln(&b)
	fmt.Println("请输入c:")
	fmt.Scanln(&c)

	//排序
	if a > b {
		t = a
		a = b
		b = t   //互换后 a是小数
	}
	if a > c {
		t = a
		a = c
		c = t   //互换后 a是三者小数
	}
	if b > c {
		t = b
		b = c
		c = t   //互换后 b是次小数
	}
	fmt.Printf("从小到到排序：a=%v b=%v c=%v \n", a, b, c)
}
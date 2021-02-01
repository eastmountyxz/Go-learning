package main
import "fmt"

func main() {
	//变量定义
	var num1 int = 99
	var num2 float64 = 3.14
	var b bool = true
	var char byte = 'h'
	var str string
	
	//fmt.Sprintf转换
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)
	
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type %T str=%q\n", str, str)
}
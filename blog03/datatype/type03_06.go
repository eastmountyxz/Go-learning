package main
import "fmt"

func main() {
	//定义字符类型
	var c1 byte = 'a'
	var c2 byte = '6'

	//当我们直接输出byte值 即输出对应字符的ASCII码值
	fmt.Println("c1 =", c1)
	fmt.Println("c2 =", c2)

	//如果需要输出对应字符 需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//var c3 byte = '杨'
	//overflow溢出
	var c3 int = '杨'
	fmt.Printf("c3=%c c3对应码值=%d\n", c3, c3)
}
package main
import "fmt"

func main() {
	//直接赋值数字 格式化输出字符
	var c1 int = 22269   // 22269->'国'  120->'x'
	fmt.Printf("c1=%c\n", c1)

	//字符类型可以运算 相当于一个整数 运算时按码值运行
	var c2  = 10 + 'a'
	fmt.Printf("c3=%c c3对应码值=%d\n", c2, c2)
}
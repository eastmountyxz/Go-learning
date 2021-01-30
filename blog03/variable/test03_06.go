package main
import "fmt"

func main() {
	//加法运算
	var i, j = 2, 3
	var res = i + j
	fmt.Println("res =", res)

	//字符串拼接
	var str1 = "hello"
	var str2 = "world"
	var str3 = str1 + str2
	fmt.Println("str3 =", str3)
}
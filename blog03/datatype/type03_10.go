package main
import "fmt"

func main() {
	//双引号字符串
	str1 := "Eastmount\nCSDN"
	fmt.Println(str1)

	//反引号字符串 ``
	str2 := `
	package main
	import "fmt"

	func main() {
		//定义字符串类型
		var name string = "我的名字叫Eastmount!"
		fmt.Println(name)
	}
	`
	fmt.Println(str2)

	//字符串拼接
	var str3 = "hello" + "world"
	str3 += "yangxiuzhang"
	fmt.Println(str3)
}
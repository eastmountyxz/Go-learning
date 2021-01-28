package main

import "fmt"    //引入包fmt 提供格式化、输出、输入函数
func main() {
	//转义字符
	fmt.Println("Eastmount\tCSDN\tyangxiuzhang")

	fmt.Println("Hello\nWorld!!!")

	fmt.Println("C:\\Users\\xiuzhang\\Desktop\\data")

	fmt.Println("秀璋说\"我爱\b娜娜\"")

	//注意:回车会替换前面的内容
	fmt.Println("秀璋爱娜娜\r珞珞")

	fmt.Println("\a")
}

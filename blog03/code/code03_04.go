package main

import (
    "fmt"
    "unicode/utf8"
)
 
func main() {
    var str string
	str = "Eastmount"

	//计算字符串长度
	fmt.Printf("The length of \"%s\" is %d. \n", str, len(str))
	
	//循环计算字符串长度
	var num int = 0
	for _, s := range str {
		fmt.Printf("%c ", s)
		num += 1
	}
	fmt.Printf("\nThe length of \"%s\" is %d. \n", str, num)

	//多字节字符
	test := "Hello, 世界"
	fmt.Println("bytes =", len(test))                      //bytes = 13
    fmt.Println("runes =", utf8.RuneCountInString(test))   //runes = 9
}
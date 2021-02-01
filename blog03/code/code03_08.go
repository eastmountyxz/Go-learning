package main
import "fmt"

func main() {
	var ch byte
	var re byte

	//从终端获取用户的输入内容
	fmt.Println("请输入任意字母:")
	fmt.Scanf("%c", &ch)
	fmt.Println("对应的ASCII码值:", ch)

	//大小写转换
	if ch >= 'A' && ch <= 'Z' {
		re = ch + 32
	} else if ch >= 'a' && ch <= 'z' {
		re = ch - 32
	}
	fmt.Printf("%c => %c", ch, re)
}
package main
import "fmt"

func main() {
	var ch byte
	var res byte
	fmt.Println("请输入字母:")
	fmt.Scanf("%c", &ch)

	//大小写转换
	if ch >= 'A' && ch <= 'Z' {
		res = ch + 32
	} else if ch >= 'a' && ch <= 'z' {
		res = ch - 32
	} else {
		res = ch
	}
	fmt.Printf("输入字母:%c %v \n", ch, ch)
	fmt.Printf("输出结果:%c %v \n", res, res)
}
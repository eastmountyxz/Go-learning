package main
import "fmt"
 
func main() {
	str := "East 秀璋"

	//方法1：Unicode遍历字符串
	fmt.Printf("Unicode遍历字符串\n")
	for _, s := range str {
		fmt.Printf("Unicode: %c  %d\n", s, s)
	}
	

	//方法2：utf-8遍历字符串
	fmt.Printf("utf-8遍历字符串\n")
	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Printf("Unicode: %c  %d\n", ch, ch)
	}
}
package main
import "fmt"
 
func main() {
	str := "East=秀璋"

	//方法1
	fmt.Printf("utf-8遍历字符串\n")
	for _, s := range str {
		fmt.Printf("utf-8遍历: %c %v \n", s, s)
	}
	fmt.Println("")

	//方法2
	str2 := []rune(str)
	for i := 0; i < len(str2); i++  {  
		fmt.Printf("utf-8遍历: index=%d, value=%c \n", i, str2[i])
	}
	fmt.Println("")

	//方法3 乱码
	fmt.Printf("unicode遍历字符串\n")
	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Printf("unicode遍历: %c  %d\n", ch, ch)
	}
}
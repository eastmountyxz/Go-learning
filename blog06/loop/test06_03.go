package main
import "fmt"

func main() {
	//1.字符串遍历-传统方法
	str := "CSDN!秀璋"
	str2 := []rune(str)           //将str转换成[]rune
	for i := 0; i < len(str2); i++  {  
		fmt.Printf("index=%d, value=%c \n", i, str2[i])
	}
	fmt.Println("")

	//2.字符串遍历-for range
	for index, val := range str {
		fmt.Printf("index=%d, value=%c \n", index, val)
	}
	fmt.Println("")

	//3.如果不转rune直接遍历(乱码)
	for i := 0; i < len(str); i++  {  
		fmt.Printf("index=%d, value=%c \n", i, str[i])
	}
}
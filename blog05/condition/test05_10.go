package main
import "fmt"

func main() {
	var key byte
	fmt.Println("请输入a到g的一个字符")
	fmt.Scanf("%c", &key)

	//switch语句
	switch key {
		case 'a':
			fmt.Println("今天星期一")
		case 'b':
			fmt.Println("今天星期二")
		case 'c':
			fmt.Println("今天星期三")
		case 'd':
			fmt.Println("今天星期四")
		case 'e':
			fmt.Println("今天星期五")
		case 'f':
			fmt.Println("今天星期六")
		case 'g':
			fmt.Println("今天星期天")
		default:
			fmt.Println("输入有误....")
	}
}
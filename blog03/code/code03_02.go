package main
import "fmt"

func main() {
	var num int = 9
	fmt.Println("num =", num)

	//判断奇偶性
	if num % 2 == 0 {
		fmt.Println("这是一个偶数")
	} else {
		fmt.Println("这是一个奇数")
	}
}
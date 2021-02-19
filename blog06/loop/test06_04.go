package main
import "fmt"

func main() {
	var result = 0
	var num = 0
	for i := 1; i <= 100; i++ {
		result += i
		num++
	}
	fmt.Println("1+2+...+100 =", result)
	fmt.Println("个数 =", num)
}
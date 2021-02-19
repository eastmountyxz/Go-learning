package main
import "fmt"

func main() {
	//return语句
	for i := 1; i<=10; i++ {
		if i==3 {
			return
		}
		fmt.Println("输出结果", i)
	}
	fmt.Println("over")
}
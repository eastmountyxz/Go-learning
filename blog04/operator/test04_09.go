package main
import "fmt"

func main() {
	//求两个数的最大值
	var a int = 10
	var b int = 20
	var max int
	
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Println("max =", max)
}
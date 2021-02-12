package main
import "fmt"

func main() {
	//求三个数中的最大值并输出结果
	var a int = 10
	var b int = 60
	var c int = 30
	var max int
	fmt.Printf("a=%v b=%v c=%v \n", a, b, c)

	if a > b {
		max = a
	} else {
		max = b
	}

	if c > max {
		max = c
	}

	fmt.Println("三个数中最大值是=", max)
}
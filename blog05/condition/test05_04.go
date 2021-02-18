package main
import "fmt"

func main() {
	var x int = 4
	var y int = 1
	
	//双分支嵌套判断
	if (x > 2) {
		if (y > 2) {
			fmt.Println(x + y)
			fmt.Println("执行语句:----a----")
		}
		fmt.Println("执行语句:----b----")
	} else {
		fmt.Println("x is", x, "y is", y)
		fmt.Println("执行语句:----c----")
	}
}
package main
import "fmt"

func main() {
	//多分支判断
	var b bool = true
	if b = false {
		fmt.Println("a")
	} else if b {
		fmt.Println("b")
	} else if !b {
		fmt.Println("c")
	} else {
		fmt.Println("d")
	}
}
package main
import "fmt"

func main() {
	//goto语句
	fmt.Println("aaaaa")
	var n int = 30
	if n > 20 {
		goto label
	}
	fmt.Println("bbbbb")
	fmt.Println("ccccc")
	label:
	fmt.Println("ddddd")
	fmt.Println("eeeee")
	fmt.Println("fffff")
}
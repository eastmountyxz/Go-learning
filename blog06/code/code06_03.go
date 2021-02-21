package main
import "fmt"
 
func main() {
	//9*9乘法表
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
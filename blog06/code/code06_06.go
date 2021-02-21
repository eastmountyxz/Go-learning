package main
import "fmt"
 
func main() {
	var n int = 4
	var m int = 5
	var res int = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%v\t", i*j)
			res++
		}
		fmt.Println("")
	} 
}
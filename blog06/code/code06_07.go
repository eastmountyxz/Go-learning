package main
import "fmt"
 
func main() {
	//Fibonacci数列
	var f1 int = 1
	var f2 int = 1
	var f3 int

	fmt.Printf("%v\n%v\n", f1, f2)
	for i := 1; i <= 10; i++ {
		f3 = f1 + f2
		fmt.Printf("%v\n", f3)
		f1 = f2
		f2 = f3
	} 
}
package main

import "fmt" 
func main() {
	var result = 0
	for i := 0; i <= 100; i++ {
		result += i 
	}
	fmt.Println("1+2+...+100=", result)
}
package main

import "fmt" 
func main() {
	var len = 5
	for i := 0; i < len; i++ {
		for j := len - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
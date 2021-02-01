package main
import "fmt"
 
func main() {
	var str string
	var tt string
	var ch byte = 'a'

	for i := 0; i < 26; i++ {
		tt = fmt.Sprintf("%c", ch)
		str += tt
		ch += 1
	}
	fmt.Println(str)
}
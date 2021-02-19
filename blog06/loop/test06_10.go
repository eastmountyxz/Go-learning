package main
import "fmt"

func main() {
	//continue跳出当前循环
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j ==2 {
				continue
			}
			fmt.Println("j =", j, "i =", i)
		}
	} 
}
package main
import "fmt"

func main() {
	//1.for-range遍历字符串
	str := "Eastmount CSDN"
	for index, val := range str {  
		fmt.Printf("index=%d, value=%c \n", index, val)
	}

	//2.for-range遍历数组
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Printf("index=%d, value=%d \n", i, num)
	}
}
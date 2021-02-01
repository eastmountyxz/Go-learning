package main
import "fmt"

func main() {
	var f float64
	var c float64

	//温度转换
	f = 69.0
	c = (5.0 / 9) * (f - 32)
	fmt.Println("华氏度 f=", f)
	fmt.Println("摄氏度 c=", c)
	fmt.Println("摄氏度整数 c=", int64(c))
}
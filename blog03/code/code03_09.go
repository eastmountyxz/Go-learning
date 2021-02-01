package main
import "fmt"
import "strconv"

func main() {
	//整型->浮点型
	var a int32 = 100
	var b float32 = float32(a)
	fmt.Printf("a=%v %T, b=%v %T\n", a, a, b, b)

	//浮点型->整型
	var c float32 = 3.14
	var d int32 = int32(c)
	fmt.Printf("c=%v %T, d=%v %T\n", c, c, d, d)

	//其他类型->string
	var e float64 = 3.14
	var f string
	f = fmt.Sprintf("%f", e)
	fmt.Printf("e=%v %T, f=%v %T\n", e, e, f, f)

	//string->其他类型
	var g string = "123.456"
	var h float64
	h, _ = strconv.ParseFloat(g, 64)
	fmt.Printf("e=%v %T, f=%v %T\n", g, g, h, h)
}
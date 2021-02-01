package main
import "fmt"

func main() {
	//数据类型转换
	var a int32 = 100

	//整型->浮点型
	var b float32 = float32(a)
	var c int8 = int8(a)

	//低精度->高精度
	var d int64 = int64(a)
	fmt.Printf("a=%v b=%v c=%v d=%v\n", a, b, c, d)
	fmt.Printf("a=%T b=%T c=%T d=%T\n", a, b, c, d)

	//浮点型->整型
	var e float32 = 3.14
	var f int32 = int32(e)
	fmt.Printf("e=%v f=%v\n", e, f)
	fmt.Printf("e=%T f=%T\n", e, f)
}
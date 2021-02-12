package main
import(
	"fmt"
	"math"
)

func main() {
	//给出三角形的三边长，求三角形面积
	var a float64 = 3.67
	var b float64 = 5.43
	var c float64 = 6.21
	var s float64
	var area float64

	s = (a + b + c) / 2
	area = math.Sqrt(s*(s-a)*(s-b)*(s-c))
	fmt.Printf("输出三条边: a=%v b=%v c=%v \n", a, b, c)
	fmt.Printf("area=%v \n", area)
}
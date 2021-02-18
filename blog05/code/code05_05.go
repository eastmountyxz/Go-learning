package main
import "fmt"
import "math"

func main() {
	/* 
	 分析思路
	   1) a,b,c是给出的变量
	   2) 使用给出的数学公式和多分支语句计算
	   3) 导入math包计算平方根
	*/
	var a float64
	var b float64
	var c float64
	fmt.Println("请输入abc三个变量")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	fmt.Printf("a=%v b=%v c=%v \n", a, b, c)

	m := b * b - 4 * a * c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v x2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Printf("无解")
	}
}
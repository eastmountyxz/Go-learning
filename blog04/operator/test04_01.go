package main
import "fmt"

func main() {
	//加法 减法 乘法
	var a int = 10
	var b int = 4
	var n1, n2, n3 int
	n1 = a + b
	n2 = a - b
	n3 = a * b
	fmt.Println("加法:", n1)
	fmt.Println("减法:", n2)
	fmt.Println("乘法:", n3)

	//除法：如果运算数都是整数，那么除后去掉小数部分，保留整数部分
	var n4 float32 = 10 / 4
	fmt.Println("除法:", n4)
	
	//除法：如果需要保留小数部分，则需要有浮点数参与运算
	var n5 float32 = 10.0 / 4
	fmt.Println("除法:", n5)

	//余数(%)：a % b = a - a / b * b
	fmt.Println("10 % 3 =", 10 % 3)      //=10-10/3*3=1
	fmt.Println("-10 % 3 =", -10 % 3)    //=-10-(-10)/3*3=-1
	fmt.Println("10 % -3 =", 10 % -3)    //=10-10/(-3)*3=1
	fmt.Println("-10 % -3 =", -10 % -3)  //=-10-(-10)/(-3)*(-3)=-1

	//++和--使用
	var i int = 10
	i++               //i = i + 1
	fmt.Println("自增:", i)
	i--               //i = i - 1
	fmt.Println("自减:", i)
}

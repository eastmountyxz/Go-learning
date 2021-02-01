package main
import "fmt"

func main() {
	//数据类型默认值
	var a int           // 0
	var b float32       // 0
	var c float64       // 0
	var d bool          // false
	var e string        // ""

	//%v表示按照变量的值输出
	//a=%v,b=%v,c=%v,d=%v,e=%v 0 0 0 false 
	fmt.Println("a=%v,b=%v,c=%v,d=%v,e=%v", a,b,c,d,e)
}
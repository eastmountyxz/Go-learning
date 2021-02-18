package main
import "fmt"

func main() {
	//type-switch
	var x interface{}     //定义空接口
	var y = 10.0
	x = y

	switch i := x.(type) {
		case nil:
			fmt.Printf("x的类型是:%T", i)
		case int:
			fmt.Printf("x是int型")
		case float64:
			fmt.Printf("x是float64型")   //输出结果
		case func(int) float64:
			fmt.Printf("x是func(int)型")
		case bool, string:
			fmt.Printf("x是bool或string型")
		default:
			fmt.Println("未知类型")
	}
}
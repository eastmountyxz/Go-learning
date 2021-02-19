package main
import "fmt"

func main() {
	//类似while循环实现输出10句"hello world"
	var i int = 1
	for {
		//循环条件
		if i > 10 {
			break           //结束循环
		}
		fmt.Println("hello world", i)
		i++
	}

	//使用do-while实现输出10句"hello world"
	var j int = 1
	for {
		//先执行后判断
		fmt.Println("hello world", j)
		j++
		if j > 10 {
			break           //结束循环
		}
	}
}
package main
import "fmt"

func main() {
	//变量定义
	var a int
	var b int
	
	//输入两个整数
	fmt.Println("请输入两个整数:")
	fmt.Scanf("%d %d", &a, &b)

	//p1指向a p2指向b
	var p1 *int = &a
	var p2 *int = &b
	var p *int

	//如果a<b则交换p1与p2的值 a存储大值
	if a < b {
		p = p1
		p1 = p2
		p2 = p
	}

	//从大到小输出a、b结果
	fmt.Printf("a=%v, b=%v\n", a, b)

	//输出p1和p2所指向变量的值
	fmt.Printf("max=%v, min=%v\n", *p1, *p2)
}
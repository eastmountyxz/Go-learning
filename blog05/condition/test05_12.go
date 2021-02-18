package main
import "fmt"

func main() {
	//switch穿透fallthrough
	var age int = 10

	switch age {
		case 10:
			fmt.Println("年龄为10")
			fallthrough   //默认只能穿透一层
		case 20:
			fmt.Println("年龄为20")
			fallthrough
		case 30:
			fmt.Println("年龄为30")
		default:
			fmt.Println("没有匹配到")
	}

}
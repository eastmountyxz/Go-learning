package main
import "fmt"

func main() {
	//假如还有97天放假 问多少个星期多少天
	var num int = 97
	var week int
	var day int
	week = num / 7
	day = num % 7
	fmt.Printf("%d个星期零%d天 \n", week, day)

	//定义一个变量保存华氏温度 转换成摄氏温度
	//摄氏温度=5/9*(摄氏温度-32)
	var huashi float32 = 134.2
	var sheshi float32
	sheshi = 5.0 / 9 * (huashi - 32)
	fmt.Printf("%v 对应的摄氏温度=%v \n", huashi, sheshi)
}

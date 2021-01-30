package main
import "fmt"

//定义全局变量
var n = 100
var m = 200
var k = "eastmount"

//等价于一次性声明
var (
	o = 300
	p = 400
	q = "yxz"
)

func main() {
	//输出全局变量
	fmt.Println("n =", n, "m =", m, "k =", k)
}
package main
import "fmt"

func main() {
	//二进制输出
	var i int = 10
	fmt.Printf("%b \n", i)

	//八进制：0-7, 满8进1，以数字0开头
	var j int = 011      //对应十进制8+1=9
	fmt.Println("j =", j)

	//十六进制：0-9及A-F，满16进1，以0x或0X开头
	var k int = 0x11    //对应十进制16+1=17
	fmt.Println("k =", k)
}

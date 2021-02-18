package main
import "fmt"

func main() {
	//判断两个数之和大于等于50
	var n1 int32 = 10
	var n2 int32 = 50
	if n1 + n2 >= 50 {
		fmt.Println("相加结果大于等于50")
	}

	//判断第一个数大于10.0，并且第二个数小于20.0
	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("输出两数之和 =", (n3+n4))
	}
	
	//判断两者之和是否能被3又能被5整除
	var n5 int32 = 10
	var n6 int32 = 5
	if (n5 + n6) % 3 == 0 && (n5 + n6) % 5 ==0 {
		fmt.Println("能被3和5整除")
	}
}
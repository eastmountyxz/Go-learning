package main
import "fmt"

//测试函数
func test() bool {
	fmt.Println("CSDN Eastmount...")
	return true
}

func main() {
	var age int = 28

	//短路与
	//由于 age<30 为false 因此后面test()函数不执行
	if age < 20 && test() {
		fmt.Println("短路与判断")
	}

	//短路或
	//由于 age>20 为true 因此后面test()函数不执行
	if age > 20 && test() {
		fmt.Println("短路或判断")
	}
}
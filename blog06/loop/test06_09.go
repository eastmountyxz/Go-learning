package main
import "fmt"

func main() {
	//求出当和第一次大于20的当前数字
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i   //求和
		if sum > 20 {
			fmt.Println("sum大于20时当前数字是", i)
			break  //跳出循环
		} 
	}

	//通过标签跳出for循环指定层
	label1:
	for i := 0; i < 4; i++ {
		//label2:
		for j := 0; j < 10; j++ {
			if j ==2 {
				//默认跳出最近的循环 这里尝试跳出外层循环
				break label1
			}
			fmt.Println("j =", j, "i =", i)
		}
	} 
}
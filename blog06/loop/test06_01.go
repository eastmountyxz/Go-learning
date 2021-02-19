package main
import "fmt"

func main() {
	//for循环
	for n := 1; n <= 10; n++ {
		fmt.Println("Eastmount", n)
	}

	//for循环第二种写法
	n := 1              //1.循环变量初始化
	for n <= 9 {        //2.循环主体
		fmt.Println("Eastmount CSDN", n)
		n++             //3.循环变量迭代
	}

	//for循环第三种写法 配合break语句
	k := 1
	for {              //等价于 for ; ; {}
		if k < 10 {
			fmt.Println("Eastmount for第三种", k)
		} else {
			break     //跳出循环
		}
		k++           //循环变量迭代
	}
}

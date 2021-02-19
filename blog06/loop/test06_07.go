package main
import "fmt"

func main() {
	var totalLevel int = 5

	//1.打印矩形
	for i := 1; i <= totalLevel; i++ {     //i表示层数
		for j := 1; j <= totalLevel; j++ { //每层输出
			fmt.Print("*")
		}
		fmt.Println()  //换行
	}
	fmt.Println()

	//2.打印正三角半个金字塔
	for i := 1; i <= totalLevel; i++ {   //i表示层数
		for j := 1; j <= i; j++ {        //每层输出
			fmt.Print("*")
		}
		fmt.Println()  //换行
	}
	fmt.Println()

	//3.打印倒三角半个金字塔
	for i := 1; i <= totalLevel; i++ {   //i表示层数
		for j := (totalLevel - i); j >= 0; j-- {       //每层输出
			fmt.Print("*")
		}
		fmt.Println()  //换行
	}
}
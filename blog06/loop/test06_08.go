package main
import "fmt"

func main() {
	/* 打印金字塔
			*           1层 1个* 规律: 2*层数-1   空格 4个 规律:总层数-当前层数
		   ***          2层 3个* 规律: 2*层数-1   空格 3个 规律:总层数-当前层数
		  *****
		 *******
		*********
	*/

	var totalLevel int = 5

	//方法一 打印实心金字塔
	for i := 1; i <= totalLevel; i++ {               //i表示层数
		//先打印空格
		for j := (totalLevel - i); j >= 0; j-- {
			fmt.Print(" ")
		}
		//打印星号
		for k := (2 * i - 1); k > 0; k-- {     
			fmt.Print("*")
		}
		fmt.Println()  //换行
	}
	fmt.Println()

	//方法二 打印空心金字塔
	for i := 1; i <= totalLevel; i++ {               //i表示层数
		//先打印空格
		for j := 1; j <= (totalLevel - i); j++ {
			fmt.Print(" ")
		}
		//打印星号
		for k := 1; k <= (2 * i - 1); k++ {
			//每层第一个和最后一个打印* 最底层全部打印*
			if k == 1 || k == (2 * i - 1) || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()  //换行
	}
}
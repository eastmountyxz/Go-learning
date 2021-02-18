package main
import "fmt"

func main() {
	//出票系统
	var month byte
	var age byte
	var price float64 = 60.0

	fmt.Println("请输入游玩月份:")
	fmt.Scanln(&month)
	fmt.Println("请输入游客年龄:")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		//旺季
		if age > 60 {
			fmt.Printf("%v月 票价 %v 年龄 %v \n", month, price / 3, age)
		} else if age >= 18 {
			fmt.Printf("%v月 票价 %v 年龄 %v \n", month, price, age)
		} else {
			fmt.Printf("%v月 票价 %v 年龄 %v \n", month, price / 2, age)
		}
	} else {
		//淡季
		if age >= 18 && age < 60 {
			fmt.Println("淡季成人票价 40")
		} else {
			fmt.Println("淡季儿童和老人票价 20")
		}
	}
}
package main
import "fmt"

func main() {
	//判断闰年
	var year int
	fmt.Println("请输入年份:")
	fmt.Scanln(&year)

	if (year % 4 == 0 && year % 100 !=0) || year % 400 == 0 {
		fmt.Println(year, "是闰年~")
	} else {
		fmt.Println(year, "不是闰年~")
	}
}

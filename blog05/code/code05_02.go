package main
import "fmt"

func main() {
	var score int
	fmt.Println("请输入分数:")
	fmt.Scanln(&score)

	//if判断
	if score >= 90 {
		fmt.Println("分数等级为A")
	} else if score >= 80 && score < 90 {
		fmt.Println("分数等级为B")
	} else if score >= 70 && score < 80 {
		fmt.Println("分数等级为C")
	} else if score >= 60 && score < 70 {
		fmt.Println("分数等级为D")
	} else {
		fmt.Println("不及格等级为E")
	}

	//switch判断
	switch {
		case score > 90:
			fmt.Println("分数等级为A")
		case score >= 80 && score < 90:
			fmt.Println("分数等级为B")
		case score >=70 && score < 80:
			fmt.Println("分数等级为C")
		case score >=60 && score < 70:
			fmt.Println("分数等级为D")
		default:
			fmt.Println("不及格等级为E")
	}
}
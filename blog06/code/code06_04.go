package main
import "fmt"
 
func main() {
	//模拟网站登录
	var name string
	var pwd string
	var loginChance = 3

	for i := 1; i <= 3; i++ {
		//输入信息
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		
		if name == "Eastmount" && pwd == "666666" {
			fmt.Println("恭喜你登录成功")
			break
		} else {
			loginChance--
			fmt.Printf("输入错误，你还有%v次机会\n", loginChance)
		}
	}
	if loginChance <= 0 {
		fmt.Println("输入错误次数过多，无法登录")
	}
}
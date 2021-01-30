package main
import "fmt"

func main() {
	//方法一：指定变量类型 int默认值为0
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)

	//方法二：根据值自行判定变量类型
	var m1, name, m3 = 100, "csdn", 3.14
	fmt.Println("m1 =", m1, "name =", name, "m3 =", m3)

	//方法三：省略var  :=声明变量并赋值
	k1, k2, k3 := 100, "yxz", 3.14
	fmt.Println("k1 =", k1, "k2 =", k2, "k3 =", k3)
}
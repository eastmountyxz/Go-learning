package main
import (
	"fmt"
	//导入utils.go文件的变量或函数 引入该model包
	"go_code/chapter03/datatype/model"
)

func main() {
	//该区域的数据可以在同一类型范围内变化
	var n int = 10
	n = 30
	n = 50
	fmt.Println("n=", n)

	//使用utils.go的HeroName变量 包名.标志符
	fmt.Println(model.HeroName)
}
package main
import "fmt"

func main() {
	//求出各个班的平均分和所有班级的平均分(3个班 每个班有4名同学)
	//解题:外层循环班级数 内存循环计算每个班级平均成绩
	var classNum int = 3
	var stuNum int = 4
	var totalsum float64 = 0.0

	for j := 1; j <= classNum; j++ { 
		var sum float64 = 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩 \n", j, i)
			fmt.Scanln(&score)
			//累计总分
			sum += score
		}
		//注意golang数据类型需要转换一致才能运算 int->float64
		fmt.Printf("第%d个班级的平均分是%v \n", j, sum / float64(stuNum))
		//计算各个班的总结成绩
		totalsum += sum
	}
	result := totalsum / float64(stuNum * classNum)
	fmt.Printf("各个班级的总成绩是%v 平均分是%v \n", totalsum, result)
}
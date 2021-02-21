package main
import (
	"fmt"
	"math/rand"
	"time"
)
 
func main() {
	//随机生成1到100的一个数，如果生成99这个数就停止
	var count int = 0
	for {
		//种子生成 使用系统时间的不确定性来进行初始化
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		fmt.Println("num =", num)
		count++
		if (num == 99) {
			break
		}
	}
	fmt.Println("生成随机数99共使用次数 =", count)
}
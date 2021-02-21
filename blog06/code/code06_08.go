package main
import "fmt"
 
func main() {
	//求2到200间的全部素数
	var n int
	var k int
	var i int
	var m int = 0

	for n = 2; n <= 200; n++ {
		k = n / 2
		//如果n被i整除终止内循环
		for i = 2; i <= k; i++ {
			if n % i == 0 {
				break
			}
		}
		//如果i>k+1表示n未被整除
		if i >= k + 1 {
			fmt.Printf("%d\t", n)
			//m控制换行 输出10个素数换行
			m = m + 1
			if m % 10 == 0 {
				fmt.Println()
			}
		}
	}
}
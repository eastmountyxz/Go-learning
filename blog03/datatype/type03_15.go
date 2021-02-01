package main
import "fmt"
import "strconv"

func main() {
	//变量定义
	var num1 int = 99
	var num2 float64 = 3.14
	var b bool = true
	var num3 int64 = 4567
	var str string
	
	//strconv.FormatInt
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv.FormatFloat(num2, 'f', 10, 64)
	//说明：'f'格式 10表示小数位保留10位 64表示小数float64
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.Itoa(int(num3))
	fmt.Printf("str type %T str=%q\n", str, str)
}
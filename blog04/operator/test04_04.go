package main
import "fmt"

func main() {
	var age int = 28

	//逻辑与运算符 &&
	if age > 20 && age < 30 {
		fmt.Println("age>20 && age<30")
	}

	if age > 30 && age < 40 {
		fmt.Println("age>30 && age<40")
	}

	//逻辑或运算符 ||
	if age > 20 || age < 30 {
		fmt.Println("age>20 || age<30")
	}

	if age > 30 || age < 40 {
		fmt.Println("age>30 || age<40")
	}

	//逻辑非运算符 !
	if age > 30 {
		fmt.Println("age>30")
	}

	if !(age > 30) {
		fmt.Println("!(age>30)")
	}
}
package main
import "fmt"

func main() {
	var a int = 1 >> 2      //001  => 0
	var b int = -1 >> 2     
	var c int = 1 << 2      //100  => 4
	var d int = -1 << 2     //-100 => 4
	var e int = 7 >> 2      //111  => 1
	var f int = 7 << 2      //111  => 11100=28

	fmt.Println("a=", a, "b=", b)
	fmt.Println("c=", c, "d=", d)
	fmt.Println("e=", e, "f=", f)
}
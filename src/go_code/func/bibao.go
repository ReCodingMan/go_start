package main
import (
	"fmt"
)

func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(36) //1=>ascii
		fmt.Println("str=", str)//1. str="hello1" 2. str="hello12" 3. str="hello123"
		return n
	}
}


func main() {
	//使用前面的代码
	f := AddUpper()

	fmt.Println(f(1))//11
	fmt.Println(f(2))//13
	fmt.Println(f(3))//16
}
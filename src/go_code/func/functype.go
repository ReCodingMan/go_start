package main
import (
	"fmt"
)

//函数类型
type myFunctype func(int, int) int

func main() {

	//类型不是一个
	type myInt int

	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)

	fmt.Println("num1=", num1, "num2=", num2)

	//函数是一种类型，在go中，可以作为参数，并且调用
	func myFun2(funvar myFunctype, num1 int, num2 int) int {
		return funvar(num1, num2)
	}

	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

}
package main
import (
	"fmt"
)

//函数类型
type myFunctype func(int, int) int

//函数是一种类型，在go中，可以作为参数，并且调用
func myFun2(funvar myFunctype, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//求和
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//支持返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	//类型不是一个
	type myInt int

	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)

	fmt.Println("num1=", num1, "num2=", num2)

	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	//求和和差
	a, b := getSumAndSub(1, 2)
	fmt.Println("a=", a, "b=", b)
}
package main
import "fmt"

func main() {
	//golang 的变量使用方式1
	//第一种：指定变量类型，声明后不赋值，使用默认值
	var i int
	fmt.Println("i=", i)

	//第二种：根据值自行判断变量类型（类型推导）
	var num = 10.11;
	fmt.Println("num=", num)

	//第三种：省略var，注意:=左侧的不应该是声明过的，否则编译失败
	name := "cheng"
	fmt.Println("cheng=", name)
}
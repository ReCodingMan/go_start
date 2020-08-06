package main
import (
	"fmt"
	//引入包
	"go_code/func/utils"
)

var age int = test()

func test() int {
	fmt.Println("test()")
	return 90
}

//init函数，通常完成初始化工作
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...")
	fmt.Println("Age=", utils.Age)
}
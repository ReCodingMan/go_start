package main
import "fmt"

// 变量使用注意事项
func main() {

	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	//i = 1.2 类型错误
	//var i int = 59 重复定义
}
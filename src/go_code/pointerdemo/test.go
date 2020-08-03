package main
import (
	"fmt"
)

// 指针类型
func main() {

	var num int = 10
	fmt.Println("num的地址= ", &num)

	var ptr *int = &num
	*ptr = 9
	fmt.Println("num= ", num)

}
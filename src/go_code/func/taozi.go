package main
import (
	"fmt"
)

func main() {
	fmt.Println("第一天桃子数量：", taozi(1))
}

func taozi(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入天数不对")
	}

	if n == 10 {
		return 1
	} else {
		return (taozi(n + 1) + 1) * 2
	}
}
package main
import (
	"fmt"
)

func cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符号错误...")
	}

	return res
}

func main() {
	
	var level int = 9

	for i:=1; i<=level; i++ {
		//先打印空格
		for k:=1; k<=level-i; k++ {
			fmt.Print(" ")
		}

		//打印*
		for j:=1; j<=2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}


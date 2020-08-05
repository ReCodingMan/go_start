package main
import (
	"fmt"
)

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
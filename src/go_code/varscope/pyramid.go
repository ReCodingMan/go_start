package main
import (
	"fmt"
)

//打印金字塔封装到方法中
func pyramid(totalLevel int) {

	for i:=1; i<=totalLevel; i++ {
		//先打印空格
		for k:=1; k<=totalLevel-i; k++ {
			fmt.Print(" ")
		}

		//打印*还是空格
		for j:=1; j<=2*i-1; j++ {
			if j==1 || j==2*i-1 || i== totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func main() {
	pyramid(9)
}
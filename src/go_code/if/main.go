package main
import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	//单分支
	// if age > 18 {
	// 	fmt.Println("已成年")
	// }

	// if age2 := 20; age2 > 18 {
	// 	fmt.Println("已成年222")
	// }

	//双分支
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	if age2 := 20; age2 > 18 {
		fmt.Println("已成年222")
	} else {
		fmt.Println("未成年222")
	}
}
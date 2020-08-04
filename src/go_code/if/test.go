package main
import (
	"fmt"
)

func main() {
	var num1 int = 10
	var num2 int = 50
	if num1 + num2 >= 50 {
		fmt.Println("大于50")
	}

	var num3 float64 = 11.0
	var num4 float64 = 17.0
	if num3 > 10.0 && num4 < 20.0 {
		fmt.Println("和是=", (num3 + num4))
	}

	var num5 int32 = 10
	var num6 int32 = 5
	if (num5 + num6) % 3 == 0 && (num5 + num6) % 5 == 0 {
		fmt.Println("能被3和5整除")
	}

	var year int = 2020
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println("是闰年")
	}
}
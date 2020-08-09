package main
import (
	"fmt"
)

func main() {

	var mychars [26]byte
	for i:=0; i<26; i++ {
		mychars[i] = 'A' + byte(i) //i需要转换成byte
	}

	for i:=0; i<26; i++ {
		fmt.Printf("%c", mychars[i])
	}

	fmt.Println()
	var intArr [5]int = [...]int {1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0

	for i:=1; i<len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}

	fmt.Printf("maxVal=%v maxValueIndex=%v", maxVal, maxValIndex)
}
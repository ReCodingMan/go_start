package main
import (
	"fmt"
)

//冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", arr)
	tmp := 0

	for i:=0; i<len(*arr)-1; i++ {
		for j:=0; j<len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}
	
	fmt.Println("排序后arr=", arr)
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}

	BubbleSort(&arr)
}
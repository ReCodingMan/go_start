package main

import (
	"fmt"
)

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 如果 leftIndex > rightIndex
	if leftIndex > rightIndex {
		fmt.Printf("找不到\n")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		// 应该在 leftIndex -- middle-1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 应该在 middle+1 -- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标为%v\n", middle)
	}
}

func main() {
	var arr [6]int = [6]int{4,5,6,7,8,9}
	BinaryFind(&arr, 0, len(arr)-1, 7)
}

package main

import "fmt"

func bubbleSort(arr *[5]int) {
	fmt.Println("排序前", *arr)

	for j:=0; j<len(*arr)-1; j++ {
		for i:=0; i<len(*arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				tmp := (*arr)[i+1]
				(*arr)[i+1] = (*arr)[i]
				(*arr)[i] = tmp
			}
		}
	}
}

func main() {
	var arr = [5]int{10,31,20,50,9}

	bubbleSort(&arr)
	fmt.Println("排序后", arr)
}

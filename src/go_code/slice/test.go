package main

import (
	"fmt"
)

func main() {
	var intArr [5]int =[...]int{1,2,3,4,5}
	slice := intArr[1:3] // [1,3)

	// 修改切片
	slice[0] = 88
	fmt.Println("slice=", slice)
	fmt.Println("intArr=", intArr)

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i,v := range slice {
		fmt.Printf("i=%v, v=%v \n", i, v)
	}
}

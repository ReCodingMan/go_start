package main

import "fmt"

func main() {
	var intArr [5]int =[...]int{1,2,3,4,5}
	slice := intArr[1:3] // [1,3)

	// 修改切片
	slice[0] = 88
	fmt.Println("slice=", slice)
	fmt.Println("intArr=", intArr)
}

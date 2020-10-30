package main

import "fmt"

func main()  {
	// 基本使用
	var intArr [5]int =[...]int{1,2,3,4,5}
	slice := intArr[1:3] // [1,3)
	fmt.Println("slice=", slice)
	fmt.Println("容量=", cap(slice))
	fmt.Println("长度=", len(slice))
}

package main

import "fmt"

func main()  {
	// 基本使用
	var intArr [5]int =[...]int{1,2,3,4,5}
	slice := intArr[1:3] // [1,3)
	fmt.Println("slice=", slice)
	fmt.Println("容量=", cap(slice))
	fmt.Println("长度=", len(slice))

	// 第二种方式，make
	var slice2 []int = make([]int, 4, 10)
	fmt.Println("slice2=", slice2)

	// 第三种方式，直接定义
	var slice3 []int = []int{1,2,3}
	fmt.Println("slice3=", slice3)
}

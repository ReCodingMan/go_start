package main

import "fmt"

func test02(arr *[4]int) {
	fmt.Printf("地址是%p\n", arr)
	(*arr)[0] = 88
}

func main() {
	arr := [4]int{11,22,33,44}
	fmt.Printf("首地址是%p\n", &arr)
	test02(&arr)
	fmt.Println(arr)
}

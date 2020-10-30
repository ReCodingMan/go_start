package main

import "fmt"

func main() {
	str := "abcdefg"
	slice := str[1:3]
	fmt.Println(slice)

	//slice[0] = 'h'
	//不许这样写，报错

	// 英文替换字符串
	var arr1 = []byte(str)
	arr1[0] = 'z'
	str2 := string(arr1)
	fmt.Println(str2)

	// 中文替换字符串
	var arr2 = []rune(str)
	arr2[0] = '北'
	str3 := string(arr2)
	fmt.Println(str3)
}

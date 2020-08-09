package main
import (
	"fmt"
)

func main() {
	//byte数组
	str := "hello@atguigu"
	//转换成切片
	slice := str[6:]
	fmt.Println("slice=", slice)

	//str[0] = 'z'
	//string是不可改变的

	//直接转切片，但是不能处理中文，byte原因
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//转成[]rune即可，按字符处理的，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=", str)



}
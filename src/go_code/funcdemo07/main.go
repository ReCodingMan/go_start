package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func (int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	//使用
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	//测试 makeSuffix
	str1 := makeSuffix(".jpg")
	fmt.Println(str1("haha.jpg"))
	fmt.Println(str1("cc"))
}


func makeSuffix(suffix string) func (string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
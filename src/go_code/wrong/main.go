package main

import "fmt"

func test() {
	// 使用 defer + recover 捕获和处理异常
	defer func() {
		err := recover()
		if err != nil { // 说明捕获到异常
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	test()

	fmt.Println("main下面的代码...")
}

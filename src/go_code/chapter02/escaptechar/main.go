package main

import "fmt"

func main() {
	fmt.Println("nihao\thaha")

	fmt.Println("nihao\nhaha")

	fmt.Println("nihao\\haha")

	fmt.Println("nihao\"haha\"")
	// \r 回车，从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("nihao\rhaha")

	fmt.Println("hehfhehfaflalflalfla",
	"hehfhehfaflalflalfla",
	"hehfhehfaflalflalfla",
	"hehfhehfaflalflalfla")
}
package main

import "fmt"

func TypeJudge(items... interface{}) {
	for index, x := range items {
		switch x.(type) {
			case bool:
				fmt.Printf("第%v个参数是 bool 类型，值是%v \n", index, x)
			case float32:
				fmt.Printf("第%v个参数是 float32 类型，值是%v \n", index, x)
			default:
				fmt.Printf("第%v个参数是 类型不确定，值是%v \n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	TypeJudge(n1,n2,n3,name,address,n4)
}

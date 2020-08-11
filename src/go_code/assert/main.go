package main
import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1,2}
	a = point
	//如何将 a 赋给 Point 变量
	var b Point
	b = a.(Point) // 类型断言
	fmt.Println(b)

	//类型断言其他案例
	//var x interface{}
	//var c float32 = 1.1
	//x = c//空接口，可以接受任意类型
	//y := x.(float32)
	//fmt.Println(y)

	//带检测的类型断言
	var x interface{}
	var c float32 = 1.1
	x = c//空接口，可以接受任意类型

	if y,ok := x.(float32); ok {
		fmt.Println("断言成功")
		fmt.Println(y)
	} else {
		fmt.Println("断言失败")
	}
}

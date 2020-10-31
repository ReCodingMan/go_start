package main

import "fmt"

// map引用类型
func modify(map1 map[int]int) {
	map1[10] = 900
}

// 定义一个结构体
type Stu struct {
	Name string
	Age int
	Address string
}

func main() {
	map1 := make(map[int]int)
	map1[1] = 10
	map1[2] = 20
	map1[3] = 30
	map1[4] = 40

	modify(map1)
	fmt.Println(map1)
	fmt.Println()

	// 结构体 value
	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	// 遍历各个学生的信息
	for k,v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Address)
		fmt.Println()
	}
}

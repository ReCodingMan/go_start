package main

import "fmt"

func main() {
	var cities map[string]string
	cities = make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "南京"
	cities["no3"] = "东京"

	// 修改
	cities["no3"] = "西京"
	fmt.Println(cities)

	// 删除
	delete(cities, "no2")
	delete(cities, "no6")
	fmt.Println(cities)

	// 删除全部，或者循环删除
	//cities = make(map[string]string)
	fmt.Println(cities)

	// 查询
	val,ok := cities["no1"]
	if ok {
		fmt.Println("存在", val)
	} else {
		fmt.Println("不存在", val)
	}

}

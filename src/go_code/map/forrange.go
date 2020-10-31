package main

import "fmt"

func main() {

	// 遍历map
	var cities map[string]string
	cities = make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "南京"
	cities["no3"] = "东京"

	for key,value := range cities {
		fmt.Printf("k=%v, v=%v \n", key, value)
	}
}

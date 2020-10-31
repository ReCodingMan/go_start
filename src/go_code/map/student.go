package main

import "fmt"

func main() {

	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Li"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "南京"

	fmt.Println(studentMap)
}

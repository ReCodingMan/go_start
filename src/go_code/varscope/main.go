package main
import (
	"fmt"
)

var age int = 50
var Name string = "jack~"

//函数
func test() {
	//作用域在函数内
	age := 10
	Name := "tom~"
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
}

func main() {
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
}
package main
import "fmt"

var n1 = 100
var n2 = 200
var name = "jack"

//一次性声明
var (
	n4 = 900
	name2 = "mary"
)

func main() {
	//一次声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1, "n2=",n2, "n3=",n3)

	//一次声明多个不同类型变量
	// var n1, name, n3 = 100, "tom", 888
	// fmt.Println("n1=",n1, "name=",name, "n3=",n3)

	//类型推导
	// n1, name, n3 := 100, "tom", 888
	// fmt.Println("n1=",n1, "name=",name, "n3=",n3)

	//输出全局变量
	fmt.Println("n1=",n1, "name=",name, "n2=",n2)
	fmt.Println("n4=",n4, "name2=",name2)
}
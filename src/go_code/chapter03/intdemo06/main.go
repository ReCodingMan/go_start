package main
// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 1
	fmt.Println("i=", i)

	//测试int8范围，-128～127
	var j int8 = -128
	fmt.Println("j=", j)
	
	//类型
	var n1 = 100
	fmt.Printf("n1 的类型 %T ", n1)

	//占用字节
	var n2 int64 = 10
	fmt.Printf("n2的数据类型 %T	n2占用字节数是 %d ", n2, unsafe.Sizeof(n2))

	//浮点型
	var price float32 = 89.12
	fmt.Println("price = ", price)

	//精度丢失问题
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=",num3, "num4=",num4)

	var c1 byte = 'a'
	var c2 byte = '0'

	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	//如果需要输出字符，需要格式化
	fmt.Printf("c1=%c", c1)
	fmt.Printf("c2=%c", c2)

	//boolean类型
	var b = false
	fmt.Println("b=", b)
	fmt.Println("b占用空间= ", unsafe.Sizeof(b))

	//string类型
	var address string = "等放假等放假 100 dfkfalfl";
	fmt.Println(address)

}
package main
import (
	"fmt"
	_ "unsafe"
	"strconv"
)

// string转基本类型使用
func main() {
	//字符串转int
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v", b, b)

	var str2 string = "1234567890"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)

	fmt.Printf("n1 type %T n1=%v", n1, n1)
	fmt.Printf("n2 type %T n2=%v", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v", f1, f1)

}
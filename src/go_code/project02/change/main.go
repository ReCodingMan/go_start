package main
import (
	"fmt"
	_ "unsafe"
	"strconv"
	"src/go_code/project02/transtype"
)

// 基本类型转string使用
func main() {
	fmt.Println(exec.HeroName)

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string // 空string

	//使用第一种方式 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str= %v", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str= %v", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str= %q", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str= %q", str, str)


	//使用第二种方法 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str= %q", str, str)

	//strconv.FormatFloat(num4, 'f', 10, 64)
	//说明：'f'格式，10：表示小数位保留10位，64：表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str= %q", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str= %q", str, str)

	//int转成string
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str= %q", str, str)
}
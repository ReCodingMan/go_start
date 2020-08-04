package main
import (
	"fmt"
)

func main() {
	
	var str string = "hello,world"
	for i:=0; i<len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	//第二种方式
	for index, val:=range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}

	//有中文
	var str2 string = "hello,world！南京"
	var str3 = []rune(str2)//str 转换成 切片[]rune
	for i:=0; i<len(str3); i++ {
		fmt.Printf("%c", str3[i])
	}
}
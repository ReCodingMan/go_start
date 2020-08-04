package main
import (
	"fmt"
)

func main() {
	
	var day byte
	fmt.Println("请输入星期几：abcdefg")
	fmt.Scanf("%c", &day)

	switch test(day)+1 {
		case 'a':
			fmt.Println("周一")
		case 'b':
			fmt.Println("周二")	
		case 'c':
			fmt.Println("周三")
		case 'd':
			fmt.Println("周四")
		case 'e':
			fmt.Println("周五")
		case 'f':
			fmt.Println("周六")
		case 'g':
			fmt.Println("周日")
		default:
			fmt.Println("输入有误")
	}

	var num int = 10
	switch num {
		case 10:
			fmt.Println("ok1")
			fallthrough //默认穿透一层
		case 20:
			fmt.Println("ok2")
		case 30:
			fmt.Println("ok3")
		default:
			fmt.Println("没找到匹配...")
	}
}

func test(b byte) byte {
	return b + 1
}
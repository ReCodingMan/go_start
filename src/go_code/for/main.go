package main
import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("打印出来了", i)
	}

	//注意事项
	//第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("你好，cc", j)
		j++
	}

	//第三种写法
	for {
		//死循环
		fmt.Println("ok~")
	}
}
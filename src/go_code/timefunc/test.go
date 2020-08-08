package main
import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	//执行前时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒", end-start)
}

func test03() {
	str := ""
	for i:=0; i<100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
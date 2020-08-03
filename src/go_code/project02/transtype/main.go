package main
import "fmt"

func main() {

	//基本数据类型转换
	var i int32 = 100
	// int 转 float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("n1=%v i=%v n2=%v n3=%v", n1, i, n2, n3)
}
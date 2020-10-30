package main

import "fmt"

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	//第一个数和第二个数都为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i:=2; i<n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	var result []uint64 = fbn(10)
	fmt.Println(result)
}

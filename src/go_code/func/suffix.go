package main
import (
	"fmt"
	"strings"
)

func main() {
	//测试
	f := makeSuffix(".jpg")
	fmt
	str1 := makeSuffix()
}

func makeSuffix(suffix string) func (string) string {

	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}
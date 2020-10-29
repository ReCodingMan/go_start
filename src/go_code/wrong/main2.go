package main

import (
	"errors"
	"fmt"
)

/**
	读取一个文件 init.conf
	如果文件不存在，报错
 */
func readConf(name string) error {
	if name ==  "init.conf" {
		return nil
	} else {
		// 返回自定义错误
		return errors.New("读取文件错误...")
	}
}

func main() {
	err := readConf("init.conf")
	if err != nil {
		// 发生错误
		panic(err)
	}

	fmt.Println("后面代码继续执行")
}

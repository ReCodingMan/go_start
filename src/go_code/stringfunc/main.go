package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	//长度
	str := "hello北"
	fmt.Println("str len=", len(str))

	//切片
	str2 := "hello北京欢迎您"
	r := []rune(str2)//转换成切片了
	for i:=0; i<len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成功", n)
	}

	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v str=%T", str, str)

	//转byte切片
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//转进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("str对应的二进制=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("str对应的十六进制=%v\n", str)

	//子串是否包含
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)

	//返回有几个子串
	num := strings.Count("cheese", "e")
	fmt.Printf("num=%v\n", num)

	//不区分大小写比较
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)

	//返回第一次出现的地方
	index := strings.Index("NLT_Abc", "Abc")
	fmt.Printf("index=%v\n", index)

	//最后一次出现位置
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index)

	//字符串替换
	str = strings.Replace("go go hello", "go", "go语言", 1)
	fmt.Printf("str=%v\n", str)

	//字符串进行拆分
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	//大写，小写
	str = "golang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	//去除两边空格
	str = " golang Hello       "
	str = strings.TrimSpace(str)
	fmt.Printf("str=%q\n", str)

	//去除两边指定的字符
	str = "!golello!"
	str = strings.Trim(str, "!")
	fmt.Printf("str=%q\n", str)

	//以xx开头
	str = "ftp://!golello!"
	b = strings.HasPrefix(str, "ftp")
	fmt.Printf("b=%v\n", b)

	//以xx结尾
	str = "ftp://!golello!"
	b = strings.HasSuffix(str, "llo!")
	fmt.Printf("b=%v\n", b)
}
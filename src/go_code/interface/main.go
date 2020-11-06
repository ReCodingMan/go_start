package main

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing()  {
	fmt.Println(monkey.Name, "生来会爬树")
}

// 声明接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}

type Test struct {

}

func (test Test) Swimming()  {
	fmt.Println("这是测试的游泳")
}
func (test Test) Flying()  {
	fmt.Println("这是测试的飞翔")
}

func (LittleMonkey *LittleMonkey) Flying() {
	fmt.Println(LittleMonkey.Name, "通过学习会飞翔")
}
func (LittleMonkey *LittleMonkey) Swimming() {
	fmt.Println(LittleMonkey.Name, "通过学习会游泳")
}
func (LittleMonkey *LittleMonkey) Studying(able BirdAble) {
	able.Flying()
}

func main()  {

	monkey := LittleMonkey{
		Monkey{
			Name: "猴子",
		},
	}

	test := Test{}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
	monkey.Studying(test)

}

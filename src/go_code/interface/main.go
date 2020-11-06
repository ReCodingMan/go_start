package main

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing()  {
	fmt.Println(monkey.Name, "生来会爬树")
}

type LittleMonkey struct {
	Monkey
}

// 声明接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (LittleMonkey *LittleMonkey) Flying() {
	fmt.Println(LittleMonkey.Name, "通过学习会飞翔")
}
func (LittleMonkey *LittleMonkey) Swimming() {
	fmt.Println(LittleMonkey.Name, "通过学习会游泳")
}

func main()  {

	monkey := LittleMonkey{
		Monkey{
			Name: "猴子",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

}

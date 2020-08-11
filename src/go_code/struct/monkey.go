package main
import (
	"fmt"
)

// Monkey结构体
type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树")
}

// Monkey结构体
type LittleMonkey struct {
	Monkey
}

// 声明接口
type BirdAble interface {
	Flying()
}

// 声明接口
type FishAble interface {
	Swimming()
}

// 让 LittleMonkey 实现 BirdAble 接口
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔")
}

// 让 LittleMonkey 实现 FishAble 接口
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳")
}

func main() {

	// 创建一个 LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct {

}

func (p Phone) start()  {
	fmt.Println("手机开始了")
}

func (p Phone) stop()  {
	fmt.Println("手机停止了")
}

type Camera struct {

}

func (c Camera) start()  {
	fmt.Println("相机开始了")
}

func (c Camera) stop()  {
	fmt.Println("相机停止了")
}

type Computer struct {
	
}

//这个方法接受usb接口类型
func (c Computer) Working(usb Usb)  {
	usb.start()
	usb.stop()
}

func main() {
	c := Computer{}
	p := Phone{}
	c.Working(p)
}